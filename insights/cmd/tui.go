package insightsCmd

import (
	"time"

	"github.com/mrusme/kopi/helpers/out"
)

func tuiOutput(insightsEntity *Insights) {
	caffeinePerCup := insightsEntity.Caffeine / float64(insightsEntity.Cups)

	out.Put(
		"Over the period of %d day(s) (or %d hour(s)) you consumed %d cup(s) of"+
			" coffee. Your total caffeine intake amounts to %.1fmg - that's"+
			" %.1fmg per cup on average. Your body hence required on average %.2f"+
			" hours per cup to eliminate the caffeine.",
		insightsEntity.PeriodInDays,
		insightsEntity.PeriodInHours,
		insightsEntity.Cups,

		insightsEntity.Caffeine,
		caffeinePerCup,
		CaffeineEliminationTime(caffeinePerCup),
	)

	out.Put("")

	out.Put(
		"With the %d cup(s) of coffee you also ingested %dml of water, %dml of"+
			" dairy and %dml of plant-based milk (which adds to your water intake).",
		insightsEntity.Cups,
		insightsEntity.Water,
		insightsEntity.RealMilk,
		insightsEntity.PlantMilk,
	)

	if insightsEntity.LastCupHoursAgo >= 0 {
		out.Put("")

		out.Pit(
			"Your last cup of coffee was %d hours ago.",
			insightsEntity.LastCupHoursAgo,
		)

		if insightsEntity.LastCupCaffeine > 0 {
			out.Pit(
				" It contained around %.1fmg of caffeine.",
				insightsEntity.LastCupCaffeine,
			)

			cet := CaffeineEliminationTime(insightsEntity.LastCupCaffeine)
			if cet > float64(insightsEntity.LastCupHoursAgo) {
				cetDiff := cet - float64(insightsEntity.LastCupHoursAgo)
				out.Put(
					" Your body will still need approximately %.1f hours (%s) to"+
						" elimiate all caffeine.",
					cetDiff,
					time.Now().Add(time.Hour*time.Duration(cetDiff)).Format(time.RFC822),
				)
			}
		} else {
			out.Put("")
		}
	}
}
