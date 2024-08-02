package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/coffee/ranking"
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
		},
	)

	for i := 0; i < len(coffees); i++ {
		coffees[i], err = coffeeDAO.Create(context.Background(), coffees[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	// Add bags
	bagDAO := bag.NewDAO(db)

	var bags []bag.Bag
	bags = append(bags,
		bag.Bag{
			CoffeeID: coffees[0].ID,

			WeightG: 450,
			Grind:   "beans",

			RoastDate:    roast,
			OpenDate:     roast,
			PurchaseDate: roast,

			PriceUSDct: 14000,
			PriceSats:  0,
		},
		bag.Bag{
			CoffeeID: coffees[1].ID,

			WeightG: 450,
			Grind:   "beans",

			RoastDate:    roast,
			OpenDate:     roast,
			PurchaseDate: roast,

			PriceUSDct: 14000,
			PriceSats:  0,
		},
		bag.Bag{
			CoffeeID: coffees[2].ID,

			WeightG: 450,
			Grind:   "beans",

			RoastDate:    roast,
			OpenDate:     roast,
			PurchaseDate: roast,

			PriceUSDct: 14000,
			PriceSats:  0,
		},
	)

	for i := 0; i < len(bags); i++ {
		bags[i], err = bagDAO.Create(context.Background(), bags[i])
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
				BagID:   bags[cofIdx].ID,
				Method:  "espresso_maker",
				Drink:   "espresso",
				CoffeeG: 14,
				BrewMl:  25,
				WaterMl: 25,
				MilkMl:  0,
				SugarG:  0,
				Vegan:   true,
				Rating:  int8(rand.Intn(6)),
			},
		)
	}

	for i := 0; i < len(cups); i++ {
		cups[i], err = cupDAO.Create(context.Background(), cups[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	avgRating, err := cupDAO.GetAvgRatingForBagID(context.Background(), bags[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Bag with ID %d has an average rating of %f\n", bags[0].ID, avgRating)

	avgRating, err = cupDAO.GetAvgRatingForCoffeeID(context.Background(), coffees[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Coffee with ID %d has an average rating of %f\n", coffees[0].ID, avgRating)

	fmt.Println()

	cupsNumber, err := cupDAO.GetCupsForPeriod(context.Background(),
		roast, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Cups for period: %d\n", cupsNumber)

	caffeine, err := cupDAO.GetCaffeineForPeriod(context.Background(),
		roast, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Caffeine for period: %fmg\n", caffeine)

	water, err := cupDAO.GetWaterForPeriod(context.Background(),
		roast, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Water for period: %dml\n", water)

	milk, err := cupDAO.GetMilkForPeriod(context.Background(),
		roast, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Milk for period: %dml\n", milk)

	realMilk, err := cupDAO.GetRealMilkForPeriod(context.Background(),
		roast, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Real milk for period: %dml\n", realMilk)

	plantMilk, err := cupDAO.GetPlantMilkForPeriod(context.Background(),
		roast, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Plant milk for period: %dml\n", plantMilk)

	fmt.Println()

	cupsPerBag, err := cupDAO.GetCupsForPeriodByBagID(context.Background(),
		roast, time.Now(), bags[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Cups per bag: %d\n", cupsPerBag)

	coffeeLeftG, err := cupDAO.GetCoffeeLeftByBagID(context.Background(),
		bags[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Coffee left in bag: %dg\n", coffeeLeftG)

	// See some rankings
	rankingDAO := ranking.NewDAO(db)

	rankedCups, err := rankingDAO.GetRanking(context.Background())
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
