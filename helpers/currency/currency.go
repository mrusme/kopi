package currency

import (
	"encoding/xml"
	"io"
	"net/http"
)

type Cube struct {
	Cube []CubeEntry `xml:"Cube>Cube>Cube"`
}

type CubeEntry struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

func FetchRates() (map[string]string, error) {
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

	rates := make(map[string]string)
	for _, entry := range cube.Cube {
		rates[entry.Currency] = entry.Rate
	}

	return rates, nil
}
