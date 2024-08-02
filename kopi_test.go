package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
)

// TestBasic tests the basic features.
func TestBasic(t *testing.T) {
	dbString := "file:test.db?cache=shared&mode=memory&_foreign_keys=1"
	db := dal.New(dbString)
	err := db.Init()
	if err != nil {
		t.Fatal(err)
	}

	// Add coffees
	coffeeDAO := coffee.NewDAO(db)
	roast, _ := time.Parse("2006-01-02", "2023-12-01")

	var coffees []coffee.Coffee
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
			RoastingDate:   roast,
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
			RoastingDate:   roast,
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
			RoastingDate:   roast,
		},
	)

	for i := 0; i < len(coffees); i++ {
		coffees[i], err = coffeeDAO.Create(context.Background(), coffees[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Have some cups
	cupDAO := cup.NewDAO(db)

	var cups []cup.Cup
	for i := 0; i < 20; i++ {
		cofIdx := rand.Intn(len(coffees))
		cups = append(cups,
			cup.Cup{
				CoffeeID: coffees[cofIdx].ID,
				Drink:    "espresso",
				CoffeeG:  14,
				BrewMl:   25,
				WaterMl:  25,
				MilkMl:   0,
				SugarG:   0,
				Vegan:    true,
				Rating:   int8(rand.Intn(6)),
			},
		)
	}

	for i := 0; i < len(cups); i++ {
		cups[i], err = cupDAO.Create(context.Background(), cups[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	avgRating, err := cupDAO.GetAvgRatingForCoffeeID(context.Background(), coffees[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Coffee with ID %d has an average rating of %f\n\n", coffees[0].ID, avgRating)

	rankedCups, err := cupDAO.GetRanking(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for _, rankedCup := range rankedCups {
		rankedCoffee, err := coffeeDAO.GetByID(context.Background(), rankedCup.CoffeeID)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("Rank #%d with an average rating of %f: %s\n",
			rankedCup.Ranking,
			rankedCup.AvgRating,
			rankedCoffee.Name,
		)
	}
	fmt.Println()

}
