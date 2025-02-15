package ocr

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mrusme/kopi/coffee"
	"github.com/xyproto/ollamaclient/v2"
)

type OCRData struct {
	Coffee   string `json:"coffee"`
	Roaster  string `json:"roaster"`
	Origin   string `json:"origin"`
	Altitude string `json:"altitude"`
	Roast    string `json:"roast"`
	Flavors  string `json:"flavors"`
	Info     string `json:"info"`
	Decaf    string `json:"decaf"`
	Drink    string `json:"drink"`
	Price    string `json:"price"`
	Vegan    string `json:"vegan"`
	Hot      string `json:"hot"`
	Cold     string `json:"cold"`
	Date     string `json:"date"`
	Rating   string `json:"rating"`
}

func GetDataFromPhoto(photoFile string) ([]OCRData, error) {
	var od []OCRData
	oc := ollamaclient.NewConfig(
		"http://10.0.0.10:11434", // TODO
		"llama3.2-vision",
		256,         // TODO: SeedOrNegative
		0.8,         // TODO: TempIfNegativeSeed
		1*time.Hour, // TODO: PullTimeout
		1*time.Hour, // TODO: HTTPTimeout
		true,        // TODO: TrimSpace
		false,       // TODO: Verbose
	)

	_, err := oc.Version()
	if err != nil {
		return []OCRData{}, err
	}

	err = oc.PullIfNeeded(true)
	if err != nil {
		return []OCRData{}, err
	}

	if found, err := oc.Has("llama3.2-vision"); err != nil || !found {
		return []OCRData{}, err
	}

	base64image, err := ollamaclient.Base64EncodeFile(photoFile)
	if err != nil {
		return []OCRData{}, err
	}

	prompt := "Look at this photo and extract all the text content, focusing on structural elements. Extract lists and maintain their structure. Preserve any hierarchical relationships. Do not comments on what the photo is and only output the extracted text content in JSON format, similar to this: { \"coffee\": \"La Gran Manzana\", \"roaster\": \"Nozy Coffee\", \"rating\": \"4/5\", \"date\": \"2025-01-30\" } Possible additional attributes for the JSON include: origin, altitude, roast, flavors, info, decaf, drink, price, vegan, sugar, hot, cold; Possible formats for the \"date\" attribute can be: 2025-01-30 (Year-Month-Day), 2025/01/30 (Year/Month/Day). Output only the JSON, nothing else."

	generatedOutput, err := oc.GetOutputChatVision(prompt, base64image)
	if err != nil {
		return []OCRData{}, err
	}
	if len(generatedOutput) == 0 {
		return []OCRData{}, errors.New("Generated output is empty")
	}

	generatedOutput = "[" + generatedOutput + "]"
	if err := json.Unmarshal([]byte(generatedOutput), &od); err != nil {
		return []OCRData{}, err
	}
	return od, nil
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

func ParseDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
