package ocr

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/spf13/viper"
	"xn--gckvb8fzb.com/kopi/bag"
	"xn--gckvb8fzb.com/kopi/coffee"
	"xn--gckvb8fzb.com/kopi/cup"
	"xn--gckvb8fzb.com/kopi/equipment"
	"xn--gckvb8fzb.com/kopi/helpers/out"
)

type OCRData struct {
	Coffee    string `json:"coffee"`
	Roaster   string `json:"roaster"`
	Origin    string `json:"origin"`
	Altitude  string `json:"altitude"`
	Roast     string `json:"roast"`
	Flavors   string `json:"flavors"`
	Info      string `json:"info"`
	Decaf     string `json:"decaf"`
	Drink     string `json:"drink"`
	Equipment string `json:"equipment"`
	Price     string `json:"price"`
	Vegan     string `json:"vegan"`
	Sugar     string `json:"sugar"`
	Hot       string `json:"hot"`
	Cold      string `json:"cold"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Rating    string `json:"rating"`
}

const DEFAULT_MODEL string = "llama3.2-vision"

func GetDataFromPhoto(photoFile string) ([]OCRData, error) {
	var od []OCRData

	if !viper.GetBool("LLM.Enabled") {
		return []OCRData{}, errors.New(
			"The LLM backend is not enabled, but it is required for OCR." +
				" Please configure and enable it in your config first.")
	}

	host := viper.GetString("LLM.Host")
	if host == "" {
		return []OCRData{}, errors.New(
			"The LLM backend is not configured properly. Please set the" +
				" `Host` in your config first.")
	}

	model := viper.GetString("LLM.Model")
	if model == "" {
		model = DEFAULT_MODEL
	}

	apiKey := viper.GetString("LLM.APIKey")
	if apiKey == "" {
		apiKey = "ollama"
	}

	imageURL, err := encodeImage(photoFile)
	if err != nil {
		return []OCRData{}, err
	}

	client := openai.NewClient(
		option.WithBaseURL(host),
		option.WithAPIKey(apiKey),
	)

	prompt := "Look at this photo and extract all the text content, focusing on structural elements. Extract lists and maintain their structure. Preserve any hierarchical relationships. Do not comments on what the photo is and only output the extracted text content in JSON format, similar to this: [{ \"coffee\": \"La Gran Manzana\", \"roaster\": \"Nozy Coffee\", \"drink\": \"Espresso\", \"rating\": \"4/5\", \"date\": \"2025-01-30\", \"time\": \"13:20\" }, { \"coffee\": \"La Loma\", \"roaster\": \"Glitch Coffee\", \"drink\": \"Cappuccino\", \"rating\": \"5/5\", \"date\": \"2025-02-10\", \"time\": \"15:44\" }] Possible additional attributes for the JSON include: origin, altitude, roast, flavors, info, decaf, equipment, drink, price, vegan, sugar, hot, cold; Possible formats for the \"date\" attribute can be: 2025-01-30 (Year-Month-Day), 2025/01/30 (Year/Month/Day). All output must be in valid JSON. Don't add explanation beyond the JSON."

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()

	completion, err := client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model: openai.ChatModel(model),
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage([]openai.ChatCompletionContentPartUnionParam{
					openai.TextContentPart(prompt),
					openai.ImageContentPart(
						openai.ChatCompletionContentPartImageImageURLParam{
							URL: imageURL,
						},
					),
				}),
			},
			Temperature: openai.Float(0.8),
		},
	)
	if err != nil {
		return []OCRData{}, err
	}

	if len(completion.Choices) == 0 {
		return []OCRData{}, errors.New("Generated output is empty")
	}

	generatedOutput := strings.TrimSpace(completion.Choices[0].Message.Content)
	if len(generatedOutput) == 0 {
		return []OCRData{}, errors.New("Generated output is empty")
	}

	out.Debug("LLM OCR:\n%s\n", generatedOutput)

	if err := json.Unmarshal([]byte(generatedOutput), &od); err != nil {
		return []OCRData{}, fmt.Errorf("%s\n\nOutput:\n%s", err, generatedOutput)
	}
	return od, nil
}

func encodeImage(photoFile string) (string, error) {
	data, err := os.ReadFile(photoFile)
	if err != nil {
		return "", err
	}
	mime := mimetype.Detect(data).String()
	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mime, encoded), nil
}

func (od *OCRData) ToEquipment(equipmentEntity *equipment.Equipment) error {
	if od.Equipment != "" {
		equipmentEntity.Name = od.Equipment
	}

	return nil
}

func (od *OCRData) ToCoffee(cfe *coffee.Coffee) error {
	if od.Roaster != "" {
		cfe.Roaster = od.Roaster
	}

	if od.Coffee != "" {
		cfe.Name = od.Coffee
	}

	if od.Origin != "" {
		cfe.Origin = od.Origin
	}

	if od.Altitude != "" {
		alts := ExtractAltitudes(od.Altitude)

		if len(alts) > 0 {
			cfe.AltitudeLowerM = alts[0]
			if len(alts) > 1 {
				cfe.AltitudeUpperM = alts[1]
			} else {
				cfe.AltitudeUpperM = cfe.AltitudeLowerM
			}
		}
	}

	if od.Roast != "" {
		cfe.Level = strings.ToLower(od.Roast)
	}

	if od.Flavors != "" {
		cfe.Flavors = od.Flavors
	}

	if od.Info != "" {
		cfe.Info = od.Info
	}

	if od.Decaf != "" {
		cfe.Decaf = true
	}

	if od.Date != "" {
		tst, err := ParseDate(od.Date)
		if err == nil {
			cfe.Timestamp = tst
			if od.Time != "" {
				tst, err = UpdateTime(cfe.Timestamp, od.Time)
				if err == nil {
					cfe.Timestamp = tst
				}
			}
		}
	}

	return nil
}

func (od *OCRData) ToBag(bg *bag.Bag) error {
	return nil
}

func (od *OCRData) ToCup(cp *cup.Cup) error {
	if od.Drink != "" {
		cp.Drink = strings.ToLower(od.Drink)
	}

	if od.Sugar != "" {
		n, u := ExtractNumberAndUnit(od.Sugar)
		lu := strings.ToLower(u)
		switch lu {
		case "g", "gram", "grams":
			cp.SugarG = n
		case "tsp", "teaspoon", "teaspoons":
			cp.SugarG = n * 4
		}
	}

	if od.Vegan != "" {
		cp.Vegan = true
	}

	if od.Rating != "" {
		if strings.Index(od.Rating, "/") > -1 {
			splitRating := strings.Split(od.Rating, "/")
			od.Rating = splitRating[0]
		}
		r, err := strconv.ParseInt(od.Rating, 10, 8)
		if err == nil {
			cp.Rating = int8(r)
		}
	}

	if od.Date != "" {
		tst, err := ParseDate(od.Date)
		if err == nil {
			cp.Timestamp = tst
			if od.Time != "" {
				tst, err = UpdateTime(cp.Timestamp, od.Time)
				if err == nil {
					cp.Timestamp = tst
				}
			}
		}
	}

	return nil
}

func ExtractAltitudes(input string) []uint16 {
	re := regexp.MustCompile(
		`\b\d+(?:[.,]\d+)?(?:m|masl)?(?:\s?-\s?\d+(?:[.,]\d+)?(?:m|masl)?)?\b`,
	)
	matches := re.FindAllString(input, -1)

	var numbers []uint16
	for _, match := range matches {
		var cleaned strings.Builder
		for _, r := range match {
			if (r >= '0' && r <= '9') || r == '-' {
				cleaned.WriteRune(r)
			}
		}

		rangeParts := strings.Split(cleaned.String(), "-")
		for _, part := range rangeParts {
			part = strings.TrimSpace(part)
			if num, err := strconv.Atoi(part); err == nil {
				numbers = append(numbers, uint16(num))
			}
		}
	}
	return numbers
}

func ExtractNumberAndUnit(input string) (uint8, string) {
	re := regexp.MustCompile(
		`(?i)\b(\d+)\s*(g|gram|grams|tsp|teaspoon|teaspoons)?\b`,
	)
	match := re.FindStringSubmatch(input)

	if len(match) < 2 {
		return 0, ""
	}

	num, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, ""
	}

	unit := ""
	if len(match) > 2 {
		unit = strings.ToLower(match[2])
	}

	return uint8(num), unit
}

func ParseDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func UpdateTime(existingTime time.Time, timeStr string) (time.Time, error) {
	const layout = "15:04"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	updatedTime := time.Date(existingTime.Year(), existingTime.Month(), existingTime.Day(),
		parsedTime.Hour(), parsedTime.Minute(), existingTime.Second(), existingTime.Nanosecond(), existingTime.Location())

	return updatedTime, nil
}
