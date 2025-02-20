package fuma

import (
	"testing"
)

type fumaTestStruct struct {
	Title string
}

// TestFindMatch tests the fuzzy matching.
func TestFindMatch(t *testing.T) {
	var ftsList []fumaTestStruct = []fumaTestStruct{
		{
			Title: "Blue Mountain",
		},
		{
			Title: "Kona",
		},
		{
			Title: "Hacienda La Esmeralda",
		},
		{
			Title: "Finca El Injerto",
		},
		{
			Title: "Kopi Luwak",
		},
		{
			Title: "Bat Coffee",
		},
	}

	var testStrings []string = []string{
		"blu Mountain",
		"blumountain",
		"bluemountain",
		"Kopi Lewark",
		"Copy luvak",
		"Brat Covfefe",
		"Finca del Inhertwo",
	}

	for _, testString := range testStrings {
		fts, err := FindMatch(&ftsList, "Title", testString)
		if err != nil {
			t.Fatal(err)
		}

		if fts == nil {
			t.Fatal("FindMatch didn't work")
		}

		t.Logf("Found: %s\n", fts.Title)
	}
}
