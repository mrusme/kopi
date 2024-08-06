package currency

import (
	"testing"
)

// TestConvertCurrencyToUSDctsWithRates tests the currency conversion.
func TestConvertCurrencyToUSDctsWithRates(t *testing.T) {
	var rates map[string]float64 = make(map[string]float64)
	rates["USD"] = 1.0835
	rates["JPY"] = 161.37

	usdCt, err := ConvertCurrencyToUSDctsWithRates(rates, 1000000, "JPY")
	if err != nil {
		t.Fatal(err)
	}

	if usdCt != 6714 {
		t.Fatalf("Expected 6714, got %d", usdCt)
	}
}
