package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
)

func main() {
	db := dal.New()
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// Add a coffee
	coffeeDAO := coffee.NewDAO(db)

	roast, _ := time.Parse("2006-01-02", "2023-12-01")
	co := coffee.Coffee{
		Roaster:        "Anthracite Coffee",
		Name:           "Ethiopia Djimmah Decaffeine",
		Origin:         "Djimmah, Ethiopia",
		AltitudeLowerM: 1700,
		AltitudeUpperM: 2200,
		Level:          "medium",
		Flavors:        "Pumpkin Yeot, Green Tangerine, Maplesyrup",
		Info:           "Long Aftertaste, Mountain Water Process Washed",
		RoastingDate:   roast,
	}

	co2, err := coffeeDAO.Create(context.Background(), co)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(co2)

	// Have a cup
	cupDAO := cup.NewDAO(db)

	c := cup.Cup{
		CoffeeID: co2.ID,
		Drink:    "espresso",
		CoffeeG:  14,
		BrewMl:   25,
		WaterMl:  25,
		MilkMl:   0,
		SugarG:   0,
		Vegan:    true,
		Rating:   5,
	}
	c2, err := cupDAO.Create(context.Background(), c)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(c2)
}
