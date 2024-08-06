package open

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/helpers"
	"github.com/mrusme/kopi/helpers/currency"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfeId int64
var cfe coffee.Coffee = coffee.Coffee{}
var bg bag.Bag = bag.Bag{}
var roastDate string
var purchaseDate string
var openDate string
var price string

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "Open a new bag of coffee",
	Long: "Open a new bag of coffee for consumption, by entering all" +
		" necessary details, adding the specific coffee to the database. If the" +
		" coffee already exists, details can be applied from the pre-existing" +
		" one using the --coffee-id flag.",
	Run: func(cmd *cobra.Command, args []string) {
		var form *huh.Form

		db, err := dal.Open(viper.GetString("Database"))
		if err != nil {
			log.Fatalf("%s\n", err)
		}

		coffeeDAO := coffee.NewDAO(db)

		// -------------------------------------------------------------------------
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
				log.Fatalln(err)
			}
		}
		// -------------------------------------------------------------------------

		var cfes []coffee.Coffee
		var roasterSuggestions []string
		var coffeeSuggestions []string

		if cfeId != 0 {
			if cfe, err = coffeeDAO.GetByID(context.Background(), cfeId); err != nil {
				log.Fatalf("Coffee could not be found: %s\n")
			}
		} else {
			cfes, err = coffeeDAO.List(context.Background())
			if err != nil {
				log.Fatalln(err)
			}

			for _, cfe := range cfes {
				roasterSuggestions = append(roasterSuggestions, cfe.Roaster)
				coffeeSuggestions = append(coffeeSuggestions, cfe.Name)
			}
		}

		accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

		if cfeId == 0 {
			form := huh.NewForm(
				huh.NewGroup(huh.NewNote().
					Title("Open a new bag of coffee").
					Description("That beautiful moment when opening a fresh"+
						"bag of coffee... mmh!\n\n"+
						"Let's start with a few details about the coffee.\n\n").
					Next(true).
					NextLabel("Next"),
				),

				huh.NewGroup(
					huh.NewInput().
						Value(&cfe.Roaster).
						Title("Coffee roaster").
						Description("What is the name of the coffee roaster?").
						Placeholder("").
						Suggestions(roasterSuggestions).
						Validate(func(s string) error {
							for i, roaster := range roasterSuggestions {
								if strings.ToLower(s) == strings.ToLower(roaster) {
									cfe = cfes[i]
									break
								}
							}

							return coffeeDAO.ValidateField(cfe, "Roaster")
						}),

					huh.NewInput().
						Value(&cfe.Name).
						Title("Coffee name").
						Description("What is the name of the coffee?").
						Placeholder("").
						Suggestions(coffeeSuggestions).
						Validate(func(s string) error {
							for i, cof := range coffeeSuggestions {
								if strings.ToLower(s) == strings.ToLower(cof) {
									cfe = cfes[i]
									break
								}
							}

							return coffeeDAO.ValidateField(cfe, "Name")
						}),
				),
			).WithAccessible(accessible)
			err = form.Run()
		}

		// If we don't have a pre-existing coffee, ask about the details:
		if cfe.ID == 0 {
			// Origin:         "Djimmah, Ethiopia",
			if cfe.Origin == "" {
				form = huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Value(&cfe.Origin).
							Title("Origin").
							Description("What origin is the coffee of?").
							Placeholder("").
							Validate(func(s string) error {
								return coffeeDAO.ValidateField(cfe, "Origin")
							}),
					),
				).WithAccessible(accessible)
				err = form.Run()
			}
			// AltitudeUpperM: 2200,
			if cfe.AltitudeUpperM == 0 {
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
								cfe.AltitudeUpperM = uint16(x)
								return coffeeDAO.ValidateField(cfe, "AltitudeUpperM")
							}),
					),
				).WithAccessible(accessible)
				err = form.Run()
			}
			// AltitudeLowerM: 1700,
			if cfe.AltitudeLowerM == 0 {
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
								cfe.AltitudeLowerM = uint16(x)
								return coffeeDAO.ValidateField(cfe, "AltitudeLowerM")
							}),
					),
				).WithAccessible(accessible)
				err = form.Run()
			}
			// Level:          "medium",
			if cfe.Level == "" {
				form := huh.NewForm(
					huh.NewGroup(
						huh.NewSelect[string]().
							Value(&cfe.Level).
							Title("Roast level").
							Description("What roast level is the coffee?").
							Options(
								huh.NewOption("Light", "light").Selected(true),
								huh.NewOption("Medium", "medium"),
								huh.NewOption("Dark", "dark"),
							),
					),
				).WithAccessible(accessible)
				err = form.Run()
			}
			// Flavors:        "Pumpkin Yeot, Green Tangerine, Maplesyrup",
			if cfe.Flavors == "" {
				form := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Value(&cfe.Flavors).
							Title("Flavors").
							Description("What are the cupping notes/flavors?").
							Placeholder("").
							Validate(func(s string) error {
								return coffeeDAO.ValidateField(cfe, "Flavors")
							}),
					),
				).WithAccessible(accessible)
				err = form.Run()
			}
			// Info:           "Long Aftertaste, Mountain Water Process Washed",
			if cfe.Info == "" {
				form := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Value(&cfe.Info).
							Title("Info").
							Description("What other information should be added?").
							Placeholder("").
							Validate(func(s string) error {
								return coffeeDAO.ValidateField(cfe, "Info")
							}),
					),
				).WithAccessible(accessible)
				err = form.Run()
			}
		}
		// </coffee details>
		bagDAO := bag.NewDAO(db)

		if bg.WeightG == 0 {
			var weight string
			form = huh.NewForm(
				huh.NewGroup(
					// WeightG: 450,
					huh.NewInput().
						Value(&weight).
						Title("Weight").
						Description("What is the bag weight in grams?").
						Placeholder("").
						Validate(func(s string) error {
							if bg.WeightG, err = strconv.ParseInt(weight, 10, 64); err != nil {
								return err
							}
							return bagDAO.ValidateField(bg, "WeightG")
						}),
				),
			).WithAccessible(accessible)
			err = form.Run()
		}

		if bg.Grind == "" {
			form = huh.NewForm(
				huh.NewGroup(
					// Grind:   "beans",
					huh.NewSelect[string]().
						Value(&bg.Grind).
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
			).WithAccessible(accessible)
			err = form.Run()
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
							if bg.RoastDate, err = time.Parse("2006-01-02", s); err != nil {
								return err
							}
							return bagDAO.ValidateField(bg, "RoastDate")
						}),
				),
			).WithAccessible(accessible)
			err = form.Run()
		} else {
			if bg.RoastDate, err = time.Parse("2006-01-02", roastDate); err != nil {
				log.Fatalln(err)
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
							if bg.PurchaseDate, err = time.Parse("2006-01-02", s); err != nil {
								return err
							}
							return bagDAO.ValidateField(bg, "PurchaseDate")
						}),
				),
			).WithAccessible(accessible)
			err = form.Run()
		} else {
			if bg.PurchaseDate, err = time.Parse("2006-01-02", purchaseDate); err != nil {
				log.Fatalln(err)
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
							"Note: If the price is entered in a curr other than USD, a" +
							" request to the ECB will be made to get the current exchange" +
							" rates.").
						Placeholder("14.50 USD").
						Validate(func(s string) error {
							if s == "" {
								return nil
							}

							var curr string
							bg.PriceUSDct, curr, err = helpers.ParsePrice(s)
							if err != nil {
								return err
							}

							if curr != "USD" {
								// TODO: Convert to USD using
								// https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml
								rates, err := currency.FetchRates()
								if err != nil {
									log.Fatalln(err)
								}

								_ = rates["USD"]
							}

							return bagDAO.ValidateField(bg, "PriceUSDct")
						}),
					// PriceSats:  0,
					// ---
				),
			).WithAccessible(accessible)
			err = form.Run()
		} else {
			var curr string
			bg.PriceUSDct, curr, err = helpers.ParsePrice(price)
			if err != nil {
				log.Fatalln(err)
			}

			if curr != "USD" {
				// TODO: Convert to USD using
				// https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml
			}
		}
	},
}

func init() {
	Cmd.Flags().Int64Var(
		&cfeId,
		"coffee-id",
		0,
		"ID of existing coffee in the database",
	)
	Cmd.Flags().StringVar(
		&cfe.Roaster,
		"roaster",
		"",
		"Name of the coffee roaster",
	)
	Cmd.Flags().StringVar(
		&cfe.Name,
		"name",
		"",
		"Name of the coffee",
	)
	Cmd.Flags().StringVar(
		&cfe.Origin,
		"origin",
		"",
		"Origin of the coffee, e.g. \"Djimmah, Ethiopia\"",
	)
	Cmd.Flags().Uint16Var(
		&cfe.AltitudeLowerM,
		"masl-low",
		0,
		"Lower meters above sea level (masl), e.g. 1700",
	)
	Cmd.Flags().Uint16Var(
		&cfe.AltitudeUpperM,
		"masl-up",
		0,
		"Upper meters above sea level (masl), e.g. 2200",
	)
	Cmd.Flags().StringVar(
		&cfe.Level,
		"level",
		"",
		"Roasting level of the coffee, e.g. \"medium\"",
	)
	Cmd.Flags().StringVar(
		&cfe.Flavors,
		"flavors",
		"",
		"Cupping notes/flavors of the coffee, e.g. \"Green Tangerine, Maplesyrup\"",
	)
	Cmd.Flags().StringVar(
		&cfe.Info,
		"info",
		"",
		"Additional info on the coffee, e.g. \"Mountain water process washed\"",
	)

	Cmd.Flags().Int64Var(
		&bg.WeightG,
		"weight",
		0,
		"Bag weight in grams, e.g. 450",
	)
	Cmd.Flags().StringVar(
		&bg.Grind,
		"grind",
		"",
		"Grind"+
			" (possible \"beans\", \"filter\", \"frenchpress\", \"stovetop\","+
			" \"espresso\")",
	)
	Cmd.Flags().StringVar(
		&roastDate,
		"roast-date",
		"",
		"Date of roast, e.g. 2024-01-01",
	)
	Cmd.Flags().StringVar(
		&purchaseDate,
		"purchase-date",
		"",
		"Date of purchase, e.g. 2024-01-01",
	)
	now := time.Now()
	Cmd.Flags().StringVar(
		&openDate,
		"open-date",
		now.Format("2006-01-02"),
		"Date of opening the bag, e.g. 2024-01-01",
	)
	Cmd.Flags().StringVar(
		&price,
		"price",
		"",
		"Price of bag, including ISO 4217 currencty code, e.g. \"14.50 USD\"",
	)
}
