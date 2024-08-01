package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
)

func main() {
	db := dal.New()
	err := db.Init()
	fmt.Print(err)

	cupDAO := cup.NewDAO(db)

	c := cup.Cup{
		CoffeeID:  1,
		Drink:     "espresso",
		Timestamp: time.Now(),
	}
	c2, err := cupDAO.Create(context.Background(), c)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	fmt.Print(c2)
}
