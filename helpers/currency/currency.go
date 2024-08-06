package currency

import (
	"encoding/xml"
	"errors"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Cube struct {
	Cube []CubeEntry `xml:"Cube>Cube>Cube"`
}

type CubeEntry struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

func FetchRates() (map[string]float64, error) {
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cube Cube
	err = xml.Unmarshal(body, &cube)
	if err != nil {
		return nil, err
	}

	rates := make(map[string]float64)
	for _, entry := range cube.Cube {
		rates[entry.Currency], err = strconv.ParseFloat(entry.Rate, 64)
		if err != nil {
			return nil, err
		}
	}

	return rates, nil
}

func ConvertCurrencyToUSDctsWithRates(rates map[string]float64, cents int64, curr string) (int64, error) {
	var rate float64
	var ok bool
	if rate, ok = rates[curr]; !ok {
		return 0, errors.New("Currency not available")
	}

	usdCents := int64(math.RoundToEven(float64(cents) * (rates["USD"] / rate)))

	return usdCents, nil
}

func ConvertCurrencyToUSDcts(cents int64, curr string) (int64, error) {
	rates, err := FetchRates()
	if err != nil {
		log.Fatalln(err)
	}

	return ConvertCurrencyToUSDctsWithRates(rates, cents, curr)
}
