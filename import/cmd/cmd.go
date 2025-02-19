package importCmd

import (
	"errors"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/mrusme/kopi/bag"
	"github.com/mrusme/kopi/coffee"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/equipment"
	"github.com/mrusme/kopi/helpers/ocr"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "import",
	Short: "Import logged coffees, cups and equipment use from a file",
	Long: "The import command lets you import coffees, cups and equipment use" +
		" from a file.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var mtypes []mimetype.MIME
		for _, arg := range args {
			mtype, err := mimetype.DetectFile(arg)
			out.NilOrDie(err)
			mtypes = append(mtypes, *mtype)
		}
		for i, arg := range args {
			ftype := mtypes[i].String()
			fcat := ftype[:strings.IndexByte(ftype, '/')]
			ImportFile(arg, fcat, ftype, mtypes[i].Extension())
		}
	},
}

func ImportFile(file string, fcat string, ftype string, fext string) error {
	switch fcat {
	case "image":
		return ImportImageFile(file)
	default:
		return errors.New("Unsupported file format")
	}
}

func ImportImageFile(file string) error {
	od, err := ocr.GetDataFromPhoto(file)
	if err != nil {
		return err
	}

	for _, ode := range od {
		cfe := coffee.Coffee{}
		ode.ToCoffee(&cfe)
	}

	return nil
}

func ProcessEntities(cfe *coffee.Coffee, bg *bag.Bag, cp *cup.Cup, eq *equipment.Equipment) {
	/* TODO: Workflow
	1) Check if equipment exists in database
		-> Yes: Keep ref, continue
		-> No: Call add equipment form, save to db, keep ref
	2) Check if coffee exists in database
		-> Yes: Keep ref, continue
		-> No: Call add coffee form, save to db, keep ref
	3) Check if bag exists in database
		-> Yes: Keep ref, continue
		-> No: Call add bag form, save to db , keep ref
	4) Check if cup exists in database
		-> Yes: Disregard all previous refs, continue
		-> No: Call add drink form, save to db, disregard all refs
	*/
}
