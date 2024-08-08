package cupDrinkCmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/drink"
	"github.com/mrusme/kopi/helpers"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/mrusme/kopi/method"
)

var theme *huh.Theme = huh.ThemeBase()

func formCup(cupDAO *cup.DAO, accessible bool) {
	var form *huh.Form
	var err error

	var coffeeSuggestions []string

	if cp.BagID != -1 {
		if _, err = cupDAO.GetByID(context.Background(), cp.BagID); err != nil {
			out.Die("Bag could not be found in database: %s", err)
		}
	} else if cp.BagID == -1 {
		bagLabelDAO := bagLabel.NewDAO(cupDAO.DB())
		labels, err := bagLabelDAO.GetLabelsOfNonEmptyBags(context.Background())
		if err != nil {
			out.Die("Bag labels could not be loaded: %s", err)
		}

		for _, label := range labels {
			coffeeSuggestions = append(coffeeSuggestions, fmt.Sprintf(
				"%s %s",
				label.Roaster,
				label.Name,
			))
		}

		var pick string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&pick).
					Title("Bag").
					Description("Which bag was used for the cup?").
					Placeholder("").
					Suggestions(coffeeSuggestions).
					Validate(func(s string) error {
						for _, label := range labels {
							if strings.ToLower(s) == strings.ToLower(fmt.Sprintf(
								"%s %s", label.Roaster, label.Name)) {
								cp.BagID = label.BagID
								return nil
							}
						}
						return errors.New("Bag was not found")

					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	//
	//
	// BrewMl  uint16 `validate:"gt=0,lte=1000,ltefield=WaterMl"`
	// WaterMl uint16 `validate:"gt=0,lte=1000,gtefield=BrewMl"`
	// MilkMl  uint16 `validate:"gte=0,lte=1000"`
	// SugarG  uint16 `validate:"gte=0,lte=100"`
	// Vegan   bool   `validate:""`

	// Drink  string `validate:"required"`
	var theDrink drink.Drink
	drinkDAO := drink.NewDAO(cupDAO.DB())
	drinks, err := drinkDAO.List(context.Background())
	if err != nil {
		out.Die("%s", err)
	}
	if cp.Drink == "" {

		var opts []huh.Option[string]
		for i, dks := range drinks {
			opt := huh.NewOption[string](dks.Name, dks.ID)
			if i == 0 {
				opt = opt.Selected(true)
			}
			opts = append(opts, opt)
		}

		form = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&cp.Drink).
					Title("Drink").
					Description("What was the drink?").
					Options(
						opts...,
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
	for _, drk := range drinks {
		if drk.ID == cp.Drink {
			theDrink = drk
			break
		}
	}
	// Method string `validate:"required"`
	if cp.Method == "" {
		methodDAO := method.NewDAO(cupDAO.DB())
		methods, err := methodDAO.List(context.Background())
		if err != nil {
			out.Die("%s", err)
		}

		var opts []huh.Option[string]
		for i, mth := range methods {
			if !strings.Contains(theDrink.CompatibleMethods, mth.ID) {
				continue
			}
			opt := huh.NewOption[string](mth.Name, mth.ID)
			if i == 0 {
				opt = opt.Selected(true)
			}
			opts = append(opts, opt)
		}

		form = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&cp.Method).
					Title("Method").
					Description("What method was used for the cup?").
					Options(
						opts...,
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
	// EquipmentIDs string `validate:"is_idslist"`
	// TODO
	// CoffeeG uint16 `validate:"gt=0,lte=200"`
	if cp.CoffeeG == 0 {
		var val string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&val).
					Title("Grams of coffee").
					Description("How much coffee was used to prepare the drink, in grams?").
					Placeholder(strconv.FormatUint(uint64(theDrink.CoffeeG), 10)).
					Validate(func(s string) error {
						if val == "" {
							cp.CoffeeG = theDrink.CoffeeG
							return nil
						}
						x, err := strconv.ParseUint(val, 10, 8)
						if err != nil {
							return err
						}
						cp.CoffeeG = uint8(x)
						return cupDAO.ValidateField(cp, "CoffeeG")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
}
