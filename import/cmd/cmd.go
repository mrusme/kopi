package importCmd

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/mrusme/kopi/bag"
	bagOpenCmd "github.com/mrusme/kopi/bag/cmd/open"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/cup"
	cupDrinkCmd "github.com/mrusme/kopi/cup/cmd/drink"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/drink"
	"github.com/mrusme/kopi/equipment"
	equipmentAddCmd "github.com/mrusme/kopi/equipment/cmd/add"
	"github.com/mrusme/kopi/helpers/fuma"
	"github.com/mrusme/kopi/helpers/ocr"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type EntitySet struct {
	EquipmentEntity equipment.Equipment
	CoffeeEntity    coffee.Coffee
	BagEntity       bag.Bag
	CupEntity       cup.Cup
}

var Cmd = &cobra.Command{
	Use:   "import",
	Short: "Import logged coffees, cups and equipment use from a file",
	Long: "The import command lets you import coffees, cups and equipment use" +
		" from a file.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var withErrors bool = false
		var mtypes []mimetype.MIME
		var entitySetsPerFile [][]EntitySet

		var devMode bool = viper.GetBool("Developer")

		db, err := dal.Open(
			viper.GetString("Database"),
			devMode,
		)
		if err != nil {
			out.Die("%s", err)
		}

		for _, arg := range args {
			mtype, err := mimetype.DetectFile(arg)
			out.NilOrDie(err)
			mtypes = append(mtypes, *mtype)
		}

		for i, arg := range args {
			ftype := mtypes[i].String()
			fcat := ftype[:strings.IndexByte(ftype, '/')]
			entitySets, err := ImportFile(arg, fcat, ftype, mtypes[i].Extension())
			out.NilOrDie(err)

			entitySetsPerFile = append(entitySetsPerFile, entitySets)
		}

		for fileIndex, entitySets := range entitySetsPerFile {
			for setIndex, entitySet := range entitySets {
				out.Put("Importing set %d/%d from file %d/%d",
					(setIndex + 1), len(entitySets),
					(fileIndex + 1), len(entitySetsPerFile),
				)
				err := ProcessEntities(db, &entitySet)
				if out.NilOrErr(err) {
					withErrors = true
				}
			}
		}

		if withErrors {
			out.Put("Import completed with errors")
		} else {
			out.Put("Import completed")
		}
	},
}

func ImportFile(file string, fcat string, ftype string, fext string) ([]EntitySet, error) {
	switch fcat {
	case "image":
		return ImportImageFile(file)
	default:
		return []EntitySet{}, errors.New("Unsupported file format")
	}
}

func ImportImageFile(file string) ([]EntitySet, error) {
	var entitySets []EntitySet

	od, err := ocr.GetDataFromPhoto(file)
	if err != nil {
		return entitySets, err
	}

	for _, ode := range od {
		var entitySet EntitySet
		entitySet.EquipmentEntity = equipment.Equipment{ID: -1}
		ode.ToEquipment(&entitySet.EquipmentEntity)

		entitySet.CoffeeEntity = coffee.Coffee{ID: -1}
		ode.ToCoffee(&entitySet.CoffeeEntity)

		entitySet.BagEntity = bag.Bag{ID: -1, CoffeeID: -1}
		ode.ToBag(&entitySet.BagEntity)

		entitySet.CupEntity = cup.Cup{ID: -1, BagID: -1}
		ode.ToCup(&entitySet.CupEntity)

		out.Debug("%v", ode)

		entitySets = append(entitySets, entitySet)
	}

	return entitySets, nil
}

func ProcessEntities(db *dal.DAL, entitySet *EntitySet) error {
	ctx := context.Background()
	accessible := viper.GetBool("TUI.Accessible")

	out.Debug("%v", entitySet.CupEntity.Drink)
	out.Debug("%v", entitySet.CupEntity.EquipmentIDs)

	equipmentEntity := &(*entitySet).EquipmentEntity
	coffeeEntity := &(*entitySet).CoffeeEntity
	bagEntity := &(*entitySet).BagEntity
	cupEntity := &(*entitySet).CupEntity
	// -------------------------------------------------------------------------//

	// 1. Process equipment if provided
	if equipmentEntity != nil && equipmentEntity.Name != "" {
		equipmentDAO := equipment.NewDAO(db)
		equipmentList, err := equipmentDAO.List(ctx, true)
		out.NilOrDie(err)

		// Check if equipment exists by matching name
		var existingEquipmentEntity *equipment.Equipment
		existingEquipmentEntity, err = fuma.FindMatch(&equipmentList, "Name", equipmentEntity.Name)

		if existingEquipmentEntity != nil {
			equipmentEntity.ID = existingEquipmentEntity.ID
		} else {
			// Equipment doesn't exist, add it
			equipmentAddCmd.FormEquipment(
				equipmentDAO,
				equipmentEntity,
				"Import Equipment",
				"This wizard will guide through the steps to import a new"+
					" piece of coffee equipment into the database.",
				accessible,
			)
			newEquipmentEntity, err := equipmentDAO.Create(ctx, *equipmentEntity)
			out.NilOrDie(err)
			equipmentEntity = &newEquipmentEntity
		}
	}

	// -------------------------------------------------------------------------//

	// 2. Process coffee if provided
	if coffeeEntity != nil && coffeeEntity.Name != "" {
		coffeeDAO := coffee.NewDAO(db)
		coffeeList, err := coffeeDAO.List(ctx)
		out.NilOrDie(err)

		// Check if coffee exists by matching name and roaster
		var existingCoffeeEntity *coffee.Coffee = nil
		existingCoffeeEntity, err = fuma.FindMatch(&coffeeList, "Name", coffeeEntity.Name)

		if existingCoffeeEntity != nil {
			coffeeEntity.ID = existingCoffeeEntity.ID
		} else {
			// Coffee doesn't exist, add it
			bagOpenCmd.FormCoffee(
				coffeeDAO,
				coffeeEntity,
				bagEntity,
				"Import "+coffeeEntity.Name,
				"This wizard will guide through the steps to import data on the "+
					coffeeEntity.Name+" coffee that was used.",
				accessible,
			)
			newCoffee, err := coffeeDAO.Create(ctx, *coffeeEntity)
			out.NilOrDie(err)
			coffeeEntity.ID = newCoffee.ID
		}
	}

	// -------------------------------------------------------------------------//

	// 3. Process bag if provided
	if bagEntity != nil && coffeeEntity != nil && coffeeEntity.ID != 0 {
		bagDAO := bag.NewDAO(db)
		bagEntity.CoffeeID = coffeeEntity.ID

		// TODO: Find coffee bag that is open and not empty
		var requiredCoffeeG int8 = 0
		if cupEntity != nil && cupEntity.Drink != "" {
			if cupEntity.CoffeeG > 0 {
				requiredCoffeeG = int8(cupEntity.CoffeeG)
			} else {
				drinkDAO := drink.NewDAO(db)
				drinkList, err := drinkDAO.List(ctx)
				out.NilOrDie(err)

				var existingDrinkEntity *drink.Drink = nil
				existingDrinkEntity, err = fuma.FindMatch(&drinkList, "Name", cupEntity.Drink)

				if existingDrinkEntity != nil {
					requiredCoffeeG = int8(existingDrinkEntity.CoffeeG)
				}
			}
		}

		var existingBagEntity *bag.Bag = nil
		var bagList []bag.Bag
		var err error
		if requiredCoffeeG > 0 {
			bagList, err = bagDAO.FindOpenByCoffeeIDWithAtLeast(ctx, coffeeEntity.ID, int(requiredCoffeeG))
		} else {
			bagList, err = bagDAO.FindOpenByCoffeeID(ctx, coffeeEntity.ID)
		}
		out.NilOrDie(err)

		if len(bagList) > 0 {
			existingBagEntity = &bagList[0]
			bagEntity.ID = existingBagEntity.ID
		} else {
			// Bag doesn't exist, add it
			bagEntity.OpenDate = time.Now()
			bagOpenCmd.FormBag(
				bagDAO,
				bagEntity,
				"Import Bag",
				"This wizard will guide through the steps to import data on a"+
					" bag of coffee that was opened.",
				accessible,
			)
			newBag, err := bagDAO.Create(ctx, *bagEntity)
			out.NilOrDie(err)
			bagEntity.ID = newBag.ID
			out.Debug("newBag.ID: %d", newBag.ID)
			out.Debug("bagEntity.ID: %d", bagEntity.ID)
		}
	}

	// -------------------------------------------------------------------------//

	// 4. Process cup if provided
	if cupEntity != nil && cupEntity.Drink != "" {
		cupDAO := cup.NewDAO(db)

		// Link bag ID
		cupEntity.BagID = bagEntity.ID
		out.Debug("cupEntity.BagID: %d", cupEntity.BagID)

		// Link equipment IDs
		if equipmentEntity != nil && equipmentEntity.ID > 0 {
			cupEntity.AddEquipmentID(equipmentEntity.ID)
		}

		// Show drink form and save cup
		cupDrinkCmd.FormCup(
			cupDAO,
			cupEntity,
			"Import Cup",
			"This wizard will guide through the steps to import data on a cup"+
				" of coffee that was consumed.",
			accessible,
		)
		_, err := cupDAO.Create(ctx, *cupEntity)
		out.NilOrDie(err)
	}

	return nil
}
