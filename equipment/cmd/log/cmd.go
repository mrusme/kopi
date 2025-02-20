package equipmentLogCmd

import (
	"context"

	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	equipmentLog "github.com/mrusme/kopi/equipment/log"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var globEquipmentLogEntity equipmentLog.Log = equipmentLog.Log{ID: -1, EquipmentID: -1}

var Cmd = &cobra.Command{
	Use:   "log",
	Short: "Log equipment use",
	Long: "Log coffee equipment use, including all" +
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
			_, err := developer.InjectDummyEquipment(db)
			if err != nil {
				out.Die("%s", err)
			}
		}

		accessible := viper.GetBool("TUI.Accessible")

		equipmentLogDAO := equipmentLog.NewDAO(db)
		FormEquipmentLog(
			equipmentLogDAO,
			globEquipmentLogEntity,
			"Log Equipment Data",
			"This wizard will guide through the steps to log new"+
				" data to a piece of equipment in the database.",
			accessible,
		)

		// Add equipment to database
		globEquipmentLogEntity, err = equipmentLogDAO.Create(context.Background(), globEquipmentLogEntity)
		if err != nil {
			out.Die("%s", err)
		} else {
			out.Put("Equipment %s logged!", globEquipmentLogEntity.Key)
		}
	},
}

func init() {
	Cmd.Flags().Int64Var(
		&globEquipmentLogEntity.EquipmentID,
		"equipment-id",
		-1,
		"ID of existing equipment in the database",
	)
	Cmd.Flags().StringVar(
		&globEquipmentLogEntity.Key,
		"key",
		"",
		"Key",
	)
	Cmd.Flags().StringVar(
		&globEquipmentLogEntity.Value,
		"value",
		"",
		"Value",
	)
}
