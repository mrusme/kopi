package coffeeRankingCmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mrusme/kopi/helpers/out"
	"golang.org/x/term"
)

func tuiOutput(rankedCoffees *[]RankedCoffee) {
	doc := strings.Builder{}

	// TODO: Clean up styles, consollidate somewhere centrally (theme.go?)
	lineRank := lipgloss.NewStyle().
		Foreground(lipgloss.ANSIColor(7)).
		Padding(0, 1)

	lineBarStyle := lipgloss.NewStyle().
		Foreground(lipgloss.ANSIColor(7))

	lineStyle := lipgloss.NewStyle().
		Inherit(lineBarStyle).
		Foreground(lipgloss.ANSIColor(7)).
		Background(lipgloss.ANSIColor(3)).
		Padding(0, 1).
		MarginRight(1)

	infoStyle := lineRank.
		Foreground(lipgloss.ANSIColor(6)).
		Align(lipgloss.Right)

	lineText := lipgloss.NewStyle().Inherit(lineBarStyle)

	avgRatingStyle := lineRank.Background(lipgloss.ANSIColor(2))

	termWidth, _, err := term.GetSize(0)
	out.NilOrDie(err)

	if termWidth < 80 {
		out.Die("Terminal must be at least 80 characters wide")
	}

	w := lipgloss.Width

	docStyle := lipgloss.NewStyle().Padding(0, 0, 0, 0)

	for i, rankedCoffee := range *rankedCoffees {

		lineKey := lineStyle.Render(
			fmt.Sprintf("#%d", rankedCoffee.Ranking.Ranking))
		info := infoStyle.Render(rankedCoffee.Coffee.Roaster)
		avgRating := avgRatingStyle.Render(
			fmt.Sprintf("%.1f", rankedCoffee.Ranking.AvgRating))
		lineVal := lineText.
			Width(termWidth - w(lineKey) - w(info) - w(avgRating)).
			Render(rankedCoffee.Coffee.Name)

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			lineKey,
			lineVal,
			info,
			avgRating,
		)

		doc.WriteString(lineBarStyle.Width(termWidth).Render(bar))
		if i < (len(*rankedCoffees) - 1) {
			doc.WriteString("\n")
		}
	}
	out.Put(docStyle.Render(doc.String()))
}
