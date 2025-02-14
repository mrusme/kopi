package cmd

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mrusme/kopi/helpers"
	"github.com/mrusme/kopi/helpers/out"
)

var theme *huh.Theme = huh.ThemeBase()

func formWelcome(cfgfile string, dbfile string, accessible bool) {
	var form *huh.Form

	var yesno bool = true
	form = huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Kopi").
				Description(
					"Welcome to _Kopi_, your command-line coffee journal.\n\n"+
						"The configuration has just been created in '"+cfgfile+"' and the"+
						" database has been initialized under '"+dbfile+"'."),
			huh.NewConfirm().
				Title("Would you like a quick introduction on how to use Kopi?").
				Value(&yesno).
				Affirmative("Yes").
				Negative("No"),
		),
	).WithAccessible(accessible).WithTheme(theme)
	helpers.HandleFormError(form.Run())

	if yesno == false {
		out.Put("Alright, enjoy using Kopi!")
		os.Exit(0)
	}

	codeBg := lipgloss.NewStyle().Background(lipgloss.Color("8"))
	form = huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("What is Kopi?").
				Description(
					"_Kopi_ is a command-line coffee journal that allows you to track"+
						" every aspect of your coffee consumption. With _Kopi_, you can"+
						" log and rate the coffees you make, and get an overview of your"+
						" favorite roasts, as well as manage your caffeine and dairy"+
						" intake.\n\n"+
						"On top of that, _Kopi_ allows you to add and track the use of your"+
						" coffee equipment, like grinders or foamers.",
				).
				Next(true).
				NextLabel("Interesting, tell me more!"),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Kopi respects your privacy!").
				Description(
					"_Kopi_ is a local-first tool that prioritizes full ownership of"+
						" your data. Hence, it stores everything that you log in a"+
						" local, universally readable SQLite3 database and it supports"+
						" various ways to export the data, e.g. as JSON or Markdown.\n\n"+
						"_Kopi_ is open-source and it respects your privacy. No data is "+
						" ever sent to servers on the internet and whenever _Kopi_ needs"+
						" to access an external service, it will inform you, so you have"+
						" the choice whether to allow that or not.",
				).
				Next(true).
				NextLabel("Awesome!"),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Quick-start").
				Description(
					"To start tracking your coffee consumptionm you need to first "+
						" _open a coffee bag_. This can be done using the following"+
						" command:\n\n"+
						codeBg.Render(" kopi bag open ")+"\n\n"+
						"Opening a bag means that _Kopi_ will ask you a few details"+
						" about the coffee in order to add it to your database.",
				).
				Next(true).
				NextLabel("Next"),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Quick-start").
				Description(
					"With your first coffee bag open, you can begin tracking individual"+
						" cups of coffee:\n\n"+
						codeBg.Render(" kopi cup drink ")+"\n\n"+
						"_Kopi_ will ask you a few details about the cup and log it to"+
						" your database.",
				).
				Next(true).
				NextLabel("Next"),
		),
	).WithAccessible(accessible).WithTheme(theme)
	helpers.HandleFormError(form.Run())
}
