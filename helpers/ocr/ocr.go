package ocr

import (
	"encoding/json"
	"errors"
	"time"

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
