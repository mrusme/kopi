package equipmentLogCmd

import (
	"context"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/mrusme/kopi/equipment"
	equipmentLog "github.com/mrusme/kopi/equipment/log"
	"github.com/mrusme/kopi/helpers"
	"github.com/mrusme/kopi/helpers/out"
)

var theme *huh.Theme = huh.ThemeBase()

func FormEquipmentLog(
	equipmentLogDAO *equipmentLog.DAO,
	equipmentLogEntity equipmentLog.Log,
	title string,
	description string,
	accessible bool,
) {
	var form *huh.Form

	equipmentDAO := equipment.NewDAO(equipmentLogDAO.DB())

	if equipmentLogEntity.EquipmentID == -1 {
		eqpt, err := equipmentDAO.List(context.Background(), true)
		if err != nil {
			out.Die("%s", err)
		}

		var opts []huh.Option[string]
		for _, eq := range eqpt {
			opt := huh.NewOption[string](eq.Name, strconv.FormatInt(eq.ID, 10))
			opts = append(opts, opt)
		}

		var val string
		form = huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title(title).
					Description(description).
					Next(false).
					NextLabel("Let's go!"),
			),
			huh.NewGroup(
				huh.NewSelect[string]().
					Value(&val).
					Title("Equipment").
					Description("For which equipment would you like to log something?").
					Options(
						opts...,
					),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
		equipmentLogEntity.EquipmentID, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			out.Die("%s", err)
		}
	} else {
		_, err := equipmentDAO.GetByID(context.Background(), equipmentLogEntity.EquipmentID)
		if err != nil {
			out.Die("%s", err)
		}
	}

	if equipmentLogEntity.Key == "" {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&equipmentLogEntity.Key).
					Title("Key").
					Description("Under what key would you like to log?").
					Placeholder("e.g. grind_level").
					Validate(func(s string) error {
						return equipmentLogDAO.ValidateField(equipmentLogEntity, "Key")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}

	if equipmentLogEntity.Value == "" {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Value(&equipmentLogEntity.Value).
					Title("Value").
					Description("What value would you like to log?").
					Placeholder("").
					Validate(func(s string) error {
						return equipmentLogDAO.ValidateField(equipmentLogEntity, "Value")
					}),
			),
		).WithAccessible(accessible).WithTheme(theme)
		helpers.HandleFormError(form.Run())
	}
}
