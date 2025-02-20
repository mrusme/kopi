package ocr

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/equipment"
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

func TestToEntity(t *testing.T) {
	var od []OCRData
	var data string = `[{
    "coffee": "La Gran Manzana",
    "roaster": "Nozy Coffee",
    "rating": "4/5",
    "date": "2025-02-12",
    "drink": "Americano"
}, {
    "coffee": "Flower Child",
    "roaster": "Bear Pond Espresso",
    "rating": "5/5",
    "date": "2025-02-12",
    "drink": "Espresso"
}]`

	if err := json.Unmarshal([]byte(data), &od); err != nil {
		t.Fatal(err)
	}

	for _, ode := range od {
		equipmentEntity := equipment.Equipment{ID: -1}
		ode.ToEquipment(&equipmentEntity)

		coffeeEntity := coffee.Coffee{ID: -1}
		ode.ToCoffee(&coffeeEntity)

		bagEntity := bag.Bag{ID: -1, CoffeeID: -1}
		ode.ToBag(&bagEntity)

		cupEntity := cup.Cup{ID: -1, BagID: -1}
		ode.ToCup(&cupEntity)

		if strings.ToLower(ode.Drink) != cupEntity.Drink {
			t.Logf("%s should be %s\n", cupEntity.Drink, ode.Drink)
			t.Fatal("Unexpected drink")
		}
	}
}
