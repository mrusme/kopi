package developer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/equipment"
)

func InjectDummyEquipment(
	db *dal.DAL,
) ([]equipment.Equipment, error) {
	var eqpt []equipment.Equipment
	var err error

	equipmentDAO := equipment.NewDAO(db)

	eqpt = append(eqpt,
		equipment.Equipment{
			Name:         "9Barista",
			Description:  "An Espresso Jet Engine",
			Type:         "espresso_maker",
			PurchaseDate: time.Now(),
			PriceUSDct:   0,
			PriceSats:    0,
		},
		equipment.Equipment{
			Name:         "Comandante C40 Mk4",
			Description:  "Grinding away",
			Type:         "grinder",
			PurchaseDate: time.Now(),
			PriceUSDct:   0,
			PriceSats:    0,
		},
	)

	for i := 0; i < len(eqpt); i++ {
		eqpt[i], err = equipmentDAO.Create(context.Background(), eqpt[i])
		if err != nil {
			return nil, err
		}
	}

	return eqpt, nil
}

func InjectDummyCoffee(
	db *dal.DAL,
) ([]coffee.Coffee, error) {
	var coffees []coffee.Coffee
	var err error

	coffeeDAO := coffee.NewDAO(db)

	coffees = append(coffees,
		coffee.Coffee{
			Roaster:        "Anthracite Coffee",
			Name:           "Ethiopia Djimmah Decaffeine",
			Origin:         "Djimmah, Ethiopia",
			AltitudeLowerM: 1700,
			AltitudeUpperM: 2200,
			Level:          "medium",
			Flavors:        "Pumpkin Yeot, Green Tangerine, Maplesyrup",
			Info:           "Long Aftertaste, Mountain Water Process Washed",
			Decaf:          true,
		},
		coffee.Coffee{
			Roaster:        "das ist PROBAT",
			Name:           "#1",
			Origin:         "40% Brazil Pico do Mirante Pulped Natural, 30% Guatemala El Morito Washed, 30% India Badra",
			AltitudeLowerM: 0,
			AltitudeUpperM: 0,
			Level:          "medium",
			Flavors:        "Malt, chocolate",
			Info:           "Well balanced",
			Decaf:          false,
		},
		coffee.Coffee{
			Roaster:        "Kona Coffee Purveyors",
			Name:           "Kona Peaberry",
			Origin:         "Hawaii",
			AltitudeLowerM: 0,
			AltitudeUpperM: 0,
			Level:          "medium",
			Flavors:        "Brown Sugar, Fruity, Hazelnut",
			Info:           "Batch Nr. 3451",
			Decaf:          false,
		},
	)

	for i := 0; i < len(coffees); i++ {
		coffees[i], err = coffeeDAO.Create(context.Background(), coffees[i])
		if err != nil {
			return nil, err
		}
	}

	return coffees, nil
}

func InjectDummyBags(
	db *dal.DAL,
	coffees []coffee.Coffee,
) ([]bag.Bag, error) {
	var bags []bag.Bag
	var err error

	bagDAO := bag.NewDAO(db)

	roast, _ := time.Parse("2006-01-02", "2023-12-01")

	for _, cfe := range coffees {
		bags = append(bags,
			bag.Bag{
				CoffeeID: cfe.ID,

				WeightG: 450,
				Grind:   "beans",

				RoastDate:    roast,
				OpenDate:     roast,
				PurchaseDate: roast,

				PriceUSDct: 14000,
				PriceSats:  0,
			},
		)
	}

	for i := 0; i < len(bags); i++ {
		bags[i], err = bagDAO.Create(context.Background(), bags[i])
		if err != nil {
			return nil, err
		}
	}

	return bags, nil
}

func InjectDummyCups(
	db *dal.DAL,
	eqpt []equipment.Equipment,
	coffees []coffee.Coffee,
	bags []bag.Bag,
) ([]cup.Cup, error) {
	var cups []cup.Cup
	var err error

	cupDAO := cup.NewDAO(db)

	for i := 0; i < 20; i++ {
		cofIdx := rand.Intn(len(coffees))
		cups = append(cups,
			cup.Cup{
				BagID:        bags[cofIdx].ID,
				Method:       "espresso_maker",
				Drink:        "espresso",
				EquipmentIDs: fmt.Sprintf("%d %d", eqpt[0].ID, eqpt[1].ID),
				CoffeeG:      14,
				BrewMl:       25,
				WaterMl:      25,
				MilkMl:       0,
				MilkType:     "none",
				SugarG:       0,
				Vegan:        true,
				Rating:       int8(rand.Intn(6)),
			},
		)
	}

	for i := 0; i < len(cups); i++ {
		cups[i], err = cupDAO.Create(context.Background(), cups[i])
		if err != nil {
			return nil, err
		}
	}

	return cups, nil
}
