package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/coffee/ranking"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	equipmentLog "github.com/mrusme/kopi/equipment/log"
)

// TestBasic tests the basic features.
func TestBasic(t *testing.T) {
	db, err := dal.Open("test", true)
	if err != nil {
		t.Fatal(err)
	}

	// Add equipment
	eqpt, err := developer.InjectDummyEquipment(db)
	if err != nil {
		t.Fatal(err)
	}

	// Log equipment update
	equipmentLogDAO := equipmentLog.NewDAO(db)
	equipmentLogDAO.Create(
		context.Background(),
		equipmentLog.Log{
			EquipmentID: eqpt[1].ID,

			Key:   "grind_level",
			Value: "+14",

			Timestamp: time.Now(),
		},
	)

	// Add coffees
	coffees, err := developer.InjectDummyCoffee(db)
	if err != nil {
		t.Fatal(err)
	}

	// Add bags
	bags, err := developer.InjectDummyBags(db, coffees)
	if err != nil {
		t.Fatal(err)
	}

	// Have some cups
	_, err = developer.InjectDummyCups(db, eqpt, coffees, bags)
	if err != nil {
		t.Fatal(err)
	}

	// Declare required DAOs
	coffeeDAO := coffee.NewDAO(db)
	cupDAO := cup.NewDAO(db)

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

	dtFrom, _ := dateparser.Parse(nil, "1970-01-01")

	cupsNumber, err := cupDAO.GetCupsForPeriod(context.Background(),
		dtFrom.Time, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Cups for period: %d\n", cupsNumber)

	caffeine, err := cupDAO.GetCaffeineForPeriod(context.Background(),
		dtFrom.Time, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Caffeine for period: %fmg\n", caffeine)

	water, err := cupDAO.GetWaterForPeriod(context.Background(),
		dtFrom.Time, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Water for period: %dml\n", water)

	milk, err := cupDAO.GetMilkForPeriod(context.Background(),
		dtFrom.Time, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Milk for period: %dml\n", milk)

	realMilk, err := cupDAO.GetRealMilkForPeriod(context.Background(),
		dtFrom.Time, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Real milk for period: %dml\n", realMilk)

	plantMilk, err := cupDAO.GetPlantMilkForPeriod(context.Background(),
		dtFrom.Time, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Plant milk for period: %dml\n", plantMilk)

	fmt.Println()

	cupsPerBag, err := cupDAO.GetCupsForPeriodByBagID(context.Background(),
		dtFrom.Time, time.Now(), bags[0].ID)
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
