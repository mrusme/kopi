package ocr

import (
	"testing"

	"github.com/mrusme/kopi/coffee"
)

// TestGetDataFromPhoto tests the OCR.
func TestGetDataFromPhoto(t *testing.T) {
	od, err := GetDataFromPhoto("sample.jpg")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", od)

	for _, ode := range od {
		cfe := coffee.Coffee{}
		ode.ToCoffee(&cfe)
		t.Logf("%v", cfe)
	}
}
