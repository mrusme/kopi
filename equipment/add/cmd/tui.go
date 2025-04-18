package equipmentAddCmd

import (
	"errors"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers"
	"github.com/mrusme/kopi/helpers/currency"
	"github.com/mrusme/kopi/helpers/out"
)

var theme *huh.Theme = huh.ThemeBase()

func FormEquipment(
	equipmentDAO *equipment.DAO,
	equipmentEntity *equipment.Equipment,
	title string,
	description string,
	accessible bool,
) {
	var form *huh.Form
	var err error

	if equipmentEntity.Name == "" {
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
					Value(&equipmentEntity.Name).
					Title("Name").
					Description("What is the name of the equipment?").
					Placeholder("e.g. Rancilio Silvia").
					Validate(func(s string) error {
						return equipmentDAO.ValidateField(*equipmentEntity, "Name")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	if equipmentEntity.Description == "" {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&equipmentEntity.Description).
					Title("Description").
					Description("How do you describe the equipment?").
					Placeholder("").
					Validate(func(s string) error {
						return equipmentDAO.ValidateField(*equipmentEntity, "Description")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	if equipmentEntity.Type == "" {
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&equipmentEntity.Type).
					Title("Type").
					Description("What type of equipment is it?").
					Options(
						huh.NewOption("Espresso maker", "espresso_maker").Selected(true),
						huh.NewOption("Coffee maker", "coffee_maker"),
						huh.NewOption("Filter", "filter"),
						huh.NewOption("Grinder", "grinder"),
						huh.NewOption("Frother", "frother"),
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	if purchaseDate == "" {
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&purchaseDate).
					Title("Purchase date").
					Description("When was the equipment purchased?").
					Placeholder("YYYY-MM-DD").
					Validate(func(s string) error {
						if equipmentEntity.PurchaseDate, err = time.Parse("2006-01-02", s); err != nil {
							return err
						}

						if equipmentEntity.PurchaseDate.After(time.Now()) {
							return errors.
								New("Hol'up time traveller, you haven't bought this" +
									" equipment yet.")
						}

						return equipmentDAO.ValidateField(*equipmentEntity, "PurchaseDate")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	} else {
		if equipmentEntity.PurchaseDate, err = time.Parse("2006-01-02", purchaseDate); err != nil {
			out.Die("%s", err)
		}
	}

	if price == "" {
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&price).
					Title("Price").
					Description("What was the price of the equipment?\n" +
						"Leave empty if unknown.\n\n" +
						"Note: If the price is entered in a currency other than USD, a" +
						" request to the ECB will be made to get the current exchange" +
						" rates.").
					Placeholder("420.00 USD").
					Validate(func(s string) error {
						if s == "" {
							return nil
						}

						var curr string
						equipmentEntity.PriceUSDct, curr, err = helpers.ParsePrice(s)
						if err != nil {
							return err
						}

						if curr != "USD" {
							convertPrice := func() {
								equipmentEntity.PriceUSDct, err = currency.ConvertCurrencyToUSDcts(
									equipmentEntity.PriceUSDct,
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

						return equipmentDAO.ValidateField(*equipmentEntity, "PriceUSDct")
					}),
				// PriceSats:  0,
				// ---
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	} else {
		var curr string
		equipmentEntity.PriceUSDct, curr, err = helpers.ParsePrice(price)
		if err != nil {
			out.Die("%s", err)
		}

		if curr != "USD" {
			equipmentEntity.PriceUSDct, err = currency.ConvertCurrencyToUSDcts(
				equipmentEntity.PriceUSDct,
				curr,
			)
		}
	}
}
