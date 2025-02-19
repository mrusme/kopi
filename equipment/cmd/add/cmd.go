package equipmentAddCmd

import (
	"context"

	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	globEquipmentEntity equipment.Equipment = equipment.Equipment{ID: -1}
	purchaseDate        string
	decommissionDate    string
	price               string
)

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Add equipment to your inventory",
	Long: "Add coffee equipment that you use to prepare cups, including all" +
		" the relevant details.",
	Run: func(cmd *cobra.Command, args []string) {
		var devMode bool = viper.GetBool("Developer")

		db, err := dal.Open(
			viper.GetString("Database"),
			devMode,
		)
		if err != nil {
			out.Die("%s", err)
		}

		if devMode {
		}

		accessible := viper.GetBool("TUI.Accessible")

		equipmentDAO := equipment.NewDAO(db)
		FormEquipment(equipmentDAO, globEquipmentEntity, accessible)

		// Add equipment to database
		globEquipmentEntity, err = equipmentDAO.Create(context.Background(), globEquipmentEntity)
		if err != nil {
			out.Die("%s", err)
		} else {
			out.Put("Equipment logged!")
		}
	},
}

func init() {
	Cmd.Flags().StringVar(
		&globEquipmentEntity.Name,
		"name",
		"",
		"Name of equipment",
	)
	Cmd.Flags().StringVar(
		&globEquipmentEntity.Description,
		"description",
		"",
		"Description of equipment",
	)
	Cmd.Flags().StringVar(
		&globEquipmentEntity.Type,
		"type",
		"",
		"Type of equipment",
	)
	Cmd.Flags().StringVar(
		&purchaseDate,
		"purchase-date",
		"",
		"Date of purchase, e.g. 2024-01-01",
	)
	Cmd.Flags().StringVar(
		&decommissionDate,
		"decommission-date",
		"",
		"Date of decommission, e.g. 2024-01-01",
	)
	Cmd.Flags().StringVar(
		&price,
		"price",
		"",
		"Price of equipment, including ISO 4217 currencty code, e.g. \"14.50 USD\"",
	)
}
