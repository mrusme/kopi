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
	"github.com/mrusme/kopi/equipment"
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
		labels, err := bagLabelDAO.List(context.Background(), true)
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
	methodDAO := method.NewDAO(cupDAO.DB())
	if cp.Method == "" {
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
	} else {
		method, err := methodDAO.GetByID(context.Background(), cp.Method)
		if err != nil {
			out.Die("%s", err)
		}
		if !strings.Contains(theDrink.CompatibleMethods, method.ID) {
			out.Die("There's no way you can prepare %s using %s.",
				theDrink.Name, method.Name)
		}
	}
	// EquipmentIDs string `validate:"is_idslist"`
	equipmentDAO := equipment.NewDAO(cupDAO.DB())
	if cp.EquipmentIDs == "" {
		eqpt, err := equipmentDAO.List(context.Background(), true)
		if err != nil {
			out.Die("%s", err)
		}

		var opts []huh.Option[string]
		for _, eq := range eqpt {
			if !strings.Contains(theDrink.CompatibleEquipment, eq.Type) {
				continue
			}
			opt := huh.NewOption[string](eq.Name, strconv.FormatInt(eq.ID, 10))
			opts = append(opts, opt)
		}

		var vals []string
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Value(&vals).
					Title("Equipment").
					Description("What equipment was used for preparing the cup?").
					Options(
						opts...,
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
		cp.EquipmentIDs = strings.Join(vals, " ")
	} else {
		eqIDs := strings.Split(cp.EquipmentIDs, " ")
		var IDs []int64
		for _, eqID := range eqIDs {
			id, err := strconv.ParseInt(eqID, 10, 64)
			if err != nil {
				out.Die("%s", err)
			}
			IDs = append(IDs, id)
		}
		eqpt, err := equipmentDAO.FindByIDs(context.Background(), IDs)
		if err != nil {
			out.Die("%s", err)
		}

		if len(eqpt) < len(eqIDs) {
			out.Die("Not all specified equipment IDs could be found!")
		}
	}
	// CoffeeG uint8 `validate:"gt=0,lte=200"`
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
	// BrewMl  uint16 `validate:"gt=0,lte=1000,ltefield=WaterMl"`
	if cp.BrewMl == 0 {
		var val string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&val).
					Title("Milliliters of brew").
					Description("How much brew was used to prepare the drink, in milliliters?").
					Placeholder(strconv.FormatUint(uint64(theDrink.BrewMl), 10)).
					Validate(func(s string) error {
						if val == "" {
							cp.BrewMl = theDrink.BrewMl
							return nil
						}
						x, err := strconv.ParseUint(val, 10, 16)
						if err != nil {
							return err
						}
						cp.BrewMl = uint16(x)
						return cupDAO.ValidateField(cp, "BrewMl")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
	// WaterMl uint16 `validate:"gt=0,lte=1000,gtefield=BrewMl"`
	if cp.WaterMl == 0 {
		var val string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&val).
					Title("Milliliters of water").
					Description("How much water was used to prepare the drink, in milliliters?").
					Placeholder(strconv.FormatUint(uint64(theDrink.WaterMl), 10)).
					Validate(func(s string) error {
						if val == "" {
							cp.WaterMl = theDrink.WaterMl
							return nil
						}
						x, err := strconv.ParseUint(val, 10, 16)
						if err != nil {
							return err
						}
						cp.WaterMl = uint16(x)
						return cupDAO.ValidateField(cp, "WaterMl")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
	// MilkMl  uint16 `validate:"gte=0,lte=1000"`
	if cp.MilkMl == 0 &&
		theDrink.IsAlwaysVegan == false {
		var val string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&val).
					Title("Milliliters of milk").
					Description("How much milk was used to prepare the drink, in milliliters?").
					Placeholder(strconv.FormatUint(uint64(theDrink.MilkMl), 10)).
					Validate(func(s string) error {
						if val == "" {
							cp.MilkMl = theDrink.MilkMl
							return nil
						}
						x, err := strconv.ParseUint(val, 10, 16)
						if err != nil {
							return err
						}
						cp.MilkMl = uint16(x)
						return cupDAO.ValidateField(cp, "MilkMl")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
	// Vegan   bool   `validate:""`
	if cp.Vegan != true &&
		cp.MilkMl > 0 &&
		theDrink.IsAlwaysVegan == false &&
		theDrink.CanBeVegan == true {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[bool]().
					Value(&cp.Vegan).
					Title("Vegan").
					Description("Is the milk real or plant-based?").
					Options(
						huh.NewOption("It's the real stuff", false).
							Selected(true),
						huh.NewOption("Plant-based", true),
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
	// SugarG  uint16 `validate:"gte=0,lte=100"`
	if cp.SugarG == 0 {
		var val string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&val).
					Title("Grams of sugar").
					Description("How much sugar was used to prepare the drink, in grams?").
					Placeholder(strconv.FormatUint(uint64(theDrink.SugarG), 10)).
					Validate(func(s string) error {
						if val == "" {
							cp.SugarG = theDrink.SugarG
							return nil
						}
						x, err := strconv.ParseUint(val, 10, 8)
						if err != nil {
							return err
						}
						cp.SugarG = uint8(x)
						return cupDAO.ValidateField(cp, "SugarG")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
}
