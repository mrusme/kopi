package bagOpenCmd

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/helpers"
	"github.com/mrusme/kopi/helpers/currency"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/mrusme/kopi/helpers/tui"
)

var theme *huh.Theme = huh.ThemeBase()

func FormCoffee(
	coffeeDAO *coffee.DAO,
	coffeeEntity *coffee.Coffee,
	bagEntity *bag.Bag,
	title string,
	description string,
	accessible bool,
) {
	var form *huh.Form
	var err error

	var coffeeEntitys []coffee.Coffee
	var roasterSuggestions []string
	var coffeeSuggestions []string

	if bagEntity.CoffeeID != -1 {
		if (*coffeeEntity), err = coffeeDAO.GetByID(context.Background(), bagEntity.CoffeeID); err != nil {
			out.Die("Coffee could not be found in database: %s", err)
		}
	} else {
		coffeeEntitys, err = coffeeDAO.List(context.Background())
		if err != nil {
			out.Die("%s", err)
		}

		for _, coffeeEntity := range coffeeEntitys {
			roasterSuggestions = append(roasterSuggestions, coffeeEntity.Roaster)
			coffeeSuggestions = append(coffeeSuggestions, coffeeEntity.Name)
		}
	}

	if bagEntity.CoffeeID == -1 {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title(title).
					Description(description).
					Next(false).
					NextLabel("Let's go!"),
			),
			huh.NewGroup(
				huh.NewInput().
					Value(&coffeeEntity.Roaster).
					Title("Coffee roaster").
					Description("What is the name of the coffee roaster?").
					Placeholder("").
					Suggestions(roasterSuggestions).
					Validate(func(s string) error {
						for i, roaster := range roasterSuggestions {
							if strings.ToLower(s) == strings.ToLower(roaster) {
								(*coffeeEntity) = coffeeEntitys[i]
								break
							}
						}

						return coffeeDAO.ValidateField((*coffeeEntity), "Roaster")
					}),

				huh.NewInput().
					Value(&coffeeEntity.Name).
					Title("Coffee name").
					Description("What is the name of the coffee?").
					Placeholder("").
					Suggestions(coffeeSuggestions).
					Validate(func(s string) error {
						for i, cof := range coffeeSuggestions {
							if strings.ToLower(s) == strings.ToLower(cof) {
								(*coffeeEntity) = coffeeEntitys[i]
								break
							}
						}

						return coffeeDAO.ValidateField((*coffeeEntity), "Name")
					}),
			),
		).WithAccessible(accessible).
			WithTheme(theme).
			WithKeyMap(tui.HuhKeyMap())
		helpers.HandleFormError(form.Run())
	}

	if coffeeEntity.ID >= 0 {
		bagEntity.CoffeeID = coffeeEntity.ID
	} else if coffeeEntity.ID == -1 {
		// If we don't have a pre-existing coffee, ask about the details:
		// Origin:         "Djimmah, Ethiopia",
		if coffeeEntity.Origin == "" {
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Value(&coffeeEntity.Origin).
						Title("Origin").
						Description("What origin is the coffee of?").
						Placeholder("").
						Validate(func(s string) error {
							return coffeeDAO.ValidateField((*coffeeEntity), "Origin")
						}),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
		// AltitudeLowerM: 1700,
		if coffeeEntity.AltitudeLowerM == 0 {
			var alt string
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Value(&alt).
						Title("MASL (lower)").
						Description("What lower altitude was the coffee farmed at?").
						Placeholder("").
						Validate(func(s string) error {
							x, err := strconv.ParseUint(alt, 10, 16)
							if err != nil {
								return err
							}
							coffeeEntity.AltitudeLowerM = uint16(x)
							return coffeeDAO.ValidateField((*coffeeEntity), "AltitudeLowerM")
						}),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
		// AltitudeUpperM: 2200,
		if coffeeEntity.AltitudeUpperM == 0 {
			var alt string
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Value(&alt).
						Title("MASL (upper)").
						Description("What upper altitude was the coffee farmed at?").
						Placeholder("").
						Validate(func(s string) error {
							x, err := strconv.ParseUint(alt, 10, 16)
							if err != nil {
								return err
							}
							coffeeEntity.AltitudeUpperM = uint16(x)
							return coffeeDAO.ValidateField((*coffeeEntity), "AltitudeUpperM")
						}),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
		// Level:          "medium",
		if coffeeEntity.Level == "" {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Value(&coffeeEntity.Level).
						Title("Roast level").
						Description("What roast level is the coffee?").
						Options(
							huh.NewOption("Light", "light").Selected(true),
							huh.NewOption("Medium", "medium"),
							huh.NewOption("Dark", "dark"),
						),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
		// Flavors:        "Pumpkin Yeot, Green Tangerine, Maplesyrup",
		if coffeeEntity.Flavors == "" {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Value(&coffeeEntity.Flavors).
						Title("Flavors").
						Description("What are the cupping notes/flavors?").
						Placeholder("").
						Validate(func(s string) error {
							return coffeeDAO.ValidateField((*coffeeEntity), "Flavors")
						}),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
		// Info:           "Long Aftertaste, Mountain Water Process Washed",
		if coffeeEntity.Info == "" {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Value(&coffeeEntity.Info).
						Title("Info").
						Description("What other information should be added?").
						Placeholder("").
						Validate(func(s string) error {
							return coffeeDAO.ValidateField((*coffeeEntity), "Info")
						}),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
		// Decaf:          false,
		if coffeeEntity.Decaf != true {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[bool]().
						Value(&coffeeEntity.Decaf).
						Title("Decaf").
						Description("Is the coffee decaf?").
						Options(
							huh.NewOption("There is coffee without caffeine? (No)", false).
								Selected(true),
							huh.NewOption("Yes", true),
						),
				),
			).WithAccessible(accessible).WithTheme(theme)
			helpers.HandleFormError(form.Run())
		}
	}
}

func FormBag(
	bagDAO *bag.DAO,
	bagEntity *bag.Bag,
	title string,
	description string,
	accessible bool,
) {
	var form *huh.Form
	var err error

	if bagEntity.WeightG == 0 {
		var weight string
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title(title).
					Description(description).
					Next(false).
					NextLabel("Let's go!"),
			),
			huh.NewGroup(
				// WeightG: 450,
				huh.NewInput().
					Value(&weight).
					Title("Weight").
					Description("What is the bag weight in grams?").
					Placeholder("").
					Validate(func(s string) error {
						if bagEntity.WeightG, err = strconv.ParseInt(weight, 10, 64); err != nil {
							return err
						}
						return bagDAO.ValidateField((*bagEntity), "WeightG")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	if bagEntity.Grind == "" {
		form = huh.NewForm(
			huh.NewGroup(
				// Grind:   "beans",
				huh.NewSelect[string]().
					Value(&bagEntity.Grind).
					Title("Grind").
					Description("What grind is the coffee?").
					Options(
						huh.NewOption("Beans (not ground)", "beans").Selected(true),
						huh.NewOption("Filter", "filter"),
						huh.NewOption("French Press", "frenchpress"),
						huh.NewOption("Stovetop", "stovetop"),
						huh.NewOption("Espresso", "espresso"),
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	if roastDate == "" {
		form = huh.NewForm(
			huh.NewGroup(
				// RoastDate:    roast,
				huh.NewInput().
					Value(&roastDate).
					Title("Roast date").
					Description("When was the coffee in the bag roasted?").
					Placeholder("YYYY-MM-DD").
					Validate(func(s string) error {
						if bagEntity.RoastDate, err = time.Parse("2006-01-02", s); err != nil {
							return err
						}

						if bagEntity.RoastDate.After(time.Now()) {
							return errors.
								New("Hol'up time traveller, this coffee seems **too** fresh.")
						} else if bagEntity.RoastDate.Before(time.Now().AddDate(-3, 0, 0)) {
							return errors.
								New("Frankly, you really shouldn't be drinking this anymore.")
						}

						return bagDAO.ValidateField((*bagEntity), "RoastDate")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	} else {
		if bagEntity.RoastDate, err = time.Parse("2006-01-02", roastDate); err != nil {
			out.Die("%s", err)
		}
	}

	if purchaseDate == "" {
		form = huh.NewForm(
			huh.NewGroup(
				// PurchaseDate: roast,
				huh.NewInput().
					Value(&purchaseDate).
					Title("Purchase date").
					Description("When was the bag purchased?").
					Placeholder("YYYY-MM-DD").
					Validate(func(s string) error {
						if bagEntity.PurchaseDate, err = time.Parse("2006-01-02", s); err != nil {
							return err
						}

						if bagEntity.PurchaseDate.After(time.Now()) {
							return errors.
								New("Hol'up time traveller, you haven't bought this bag yet.")
						} else if bagEntity.PurchaseDate.Before(bagEntity.RoastDate.AddDate(-1, 0, 0)) {
							return errors.
								New("This looks like some serious stock market futures" +
									" trading. Are you sure?")
						}

						return bagDAO.ValidateField((*bagEntity), "PurchaseDate")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	} else {
		if bagEntity.PurchaseDate, err = time.Parse("2006-01-02", purchaseDate); err != nil {
			out.Die("%s", err)
		}
	}

	if price == "" {
		form = huh.NewForm(
			huh.NewGroup(
				// OpenDate:     roast,
				// ---
				// PriceUSDct: 14000,
				huh.NewInput().
					Value(&price).
					Title("Price").
					Description("What was the price of the bag?\n" +
						"Leave empty if unknown.\n\n" +
						"Note: If the price is entered in a currency other than USD, a" +
						" request to the ECB will be made to get the current exchange" +
						" rates.").
					Placeholder("14.50 USD").
					Validate(func(s string) error {
						if s == "" {
							return nil
						}

						var curr string
						bagEntity.PriceUSDct, curr, err = helpers.ParsePrice(s)
						if err != nil {
							return err
						}

						if curr != "USD" {
							convertPrice := func() {
								bagEntity.PriceUSDct, err = currency.ConvertCurrencyToUSDcts(
									bagEntity.PriceUSDct,
									curr,
								)
							}

							_ = spinner.
								New().
								Title("Converting price into USD ...").
								Accessible(accessible).
								// Theme(theme). // INFO: https://github.com/charmbracelet/huh/issues/240#issuecomment-2273855313
								Action(convertPrice).
								Run()
						}

						return bagDAO.ValidateField((*bagEntity), "PriceUSDct")
					}),
				// PriceSats:  0,
				// ---
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	} else {
		var curr string
		bagEntity.PriceUSDct, curr, err = helpers.ParsePrice(price)
		if err != nil {
			out.Die("%s", err)
		}

		if curr != "USD" {
			bagEntity.PriceUSDct, err = currency.ConvertCurrencyToUSDcts(
				bagEntity.PriceUSDct,
				curr,
			)
		}
	}
}
