package ocr

import (
	"testing"
)

// TestGetDataFromPhoto tests the OCR.
func TestGetDataFromPhoto(t *testing.T) {
	od, err := GetDataFromPhoto("sample.jpg")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", od)
}

