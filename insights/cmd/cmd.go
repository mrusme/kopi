package insightsCmd

import (
	"context"
	"math"
	"time"

	"github.com/markusmobius/go-dateparser"
	"github.com/mrusme/kopi/cup"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/developer"
	"github.com/mrusme/kopi/helpers/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	periodFrom  string = "yesterday"
	periodUntil string = "now"
	outputJSON  bool   = false
	outputMD    bool   = false
)

type Insights struct {
	PeriodFrom    time.Time
	PeriodUntil   time.Time
	PeriodInDays  int
	PeriodInHours int
	Cups          int64
	Caffeine      float64
	Water         int64
	Milk          int64
	RealMilk      int64
	PlantMilk     int64
}

var Cmd = &cobra.Command{
	Use:   "insights",
	Short: "Insights into the tracked data",
	Long: "The insights command shows metrics on tracked coffees, cups and" +
		" equipment use.",
	Run: func(cmd *cobra.Command, args []string) {
		var insightsEntity Insights
		var devMode bool = viper.GetBool("Developer")

		db, err := dal.Open(
			viper.GetString("Database"),
			devMode,
		)
		if err != nil {
			out.Die("%s", err)
		}

		if devMode {
			eqpt, err := developer.InjectDummyEquipment(db)
			if err != nil {
				out.Die("%s", err)
			}
			coffees, err := developer.InjectDummyCoffee(db)
			if err != nil {
				out.Die("%s", err)
			}
			bags, err := developer.InjectDummyBags(db, coffees)
			if err != nil {
				out.Die("%s", err)
			}
			_, err = developer.InjectDummyCups(db, eqpt, coffees, bags)
			if err != nil {
				out.Die("%s", err)
			}
		}

		dtFrom, err := dateparser.Parse(nil, periodFrom)
		out.NilOrErr(err)
		insightsEntity.PeriodFrom = dtFrom.Time

		dtUntil, err := dateparser.Parse(nil, periodUntil)
		out.NilOrErr(err)
		insightsEntity.PeriodUntil = dtUntil.Time

		insightsEntity.PeriodInDays = DaysBetween(
			insightsEntity.PeriodFrom,
			insightsEntity.PeriodUntil,
		)

		insightsEntity.PeriodInHours = HoursBetween(
			insightsEntity.PeriodFrom,
			insightsEntity.PeriodUntil,
		)

		cupDAO := cup.NewDAO(db)

		insightsEntity.Cups, err = cupDAO.GetCupsForPeriod(context.Background(),
			insightsEntity.PeriodFrom, insightsEntity.PeriodUntil)
		out.NilOrErr(err)

		insightsEntity.Caffeine, err = cupDAO.GetCaffeineForPeriod(context.Background(),
			insightsEntity.PeriodFrom, insightsEntity.PeriodUntil)
		out.NilOrErr(err)

		insightsEntity.Water, err = cupDAO.GetWaterForPeriod(context.Background(),
			insightsEntity.PeriodFrom, insightsEntity.PeriodUntil)
		out.NilOrErr(err)

		insightsEntity.Milk, err = cupDAO.GetMilkForPeriod(context.Background(),
			insightsEntity.PeriodFrom, insightsEntity.PeriodUntil)
		out.NilOrErr(err)

		insightsEntity.RealMilk, err = cupDAO.GetRealMilkForPeriod(context.Background(),
			insightsEntity.PeriodFrom, insightsEntity.PeriodUntil)
		out.NilOrErr(err)

		insightsEntity.PlantMilk, err = cupDAO.GetPlantMilkForPeriod(context.Background(),
			insightsEntity.PeriodFrom, insightsEntity.PeriodUntil)
		out.NilOrErr(err)

		// TODO: GetBagsForPeriod -> range bags -> GetCupsForPeriodByBagID
		//
		// cupsPerBag, err := cupDAO.GetCupsForPeriodByBagID(context.Background(),
		// 	insightsEntity.PeriodFrom, insightsEntity.PeriodUntil, bags[0].ID)
		// out.NilOrErr(err)
		// fmt.Printf("Cups per bag: %d\n", cupsPerBag)

		if outputJSON {
			jsonOutput(&insightsEntity)
		} else if outputMD {
			mdOutput(&insightsEntity)
		} else {
			tuiOutput(&insightsEntity)
		}
	},
}

func init() {
	Cmd.Flags().StringVar(
		&periodFrom,
		"from",
		"yesterday",
		"Show insights starting from a specific date/time",
	)
	Cmd.Flags().StringVar(
		&periodUntil,
		"until",
		"now",
		"Show insights until a specific date/time",
	)

	Cmd.Flags().BoolVar(
		&outputJSON,
		"json",
		false,
		"Output JSON",
	)
	Cmd.Flags().BoolVar(
		&outputMD,
		"markdown",
		false,
		"Output Markdown",
	)
	Cmd.MarkFlagsMutuallyExclusive("json", "markdown")
}

func DaysBetween(start, end time.Time) int {
	start = start.Truncate(24 * time.Hour)
	end = end.Truncate(24 * time.Hour)
	return int(end.Sub(start).Hours() / 24)
}

func HoursBetween(start, end time.Time) int {
	duration := end.Sub(start)
	truncatedDuration := duration.Truncate(time.Hour)
	return int(truncatedDuration.Hours())
}

const CaffeineHalfLife = 5.0

func CaffeineEliminationTime(initialCaffeine float64) float64 {
	if initialCaffeine <= 1 {
		return 0
	}
	return CaffeineHalfLife * (math.Log2(initialCaffeine / 1))
}
