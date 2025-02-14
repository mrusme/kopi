package formatter

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/cdfmlr/ellipsis"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

func ListToMarkdown[T any](entities []T, fields []string) string {
	var fieldsStr string

	if len(entities) == 0 {
		return ""
	}

	if len(fields) == 0 {
		fieldsStr = ""
	} else {
		fieldsStr = strings.Join(fields, " ")
	}

	header := "|"
	sep := "|"
	var rows []string
	for li, entity := range entities {
		row := "|"

		st := reflect.TypeOf(entity)
		for i := 0; i < st.NumField(); i++ {

			stField := st.Field(i).Name

			if len(fieldsStr) == 0 || strings.Contains(fieldsStr, stField) {
				if li == 0 {
					header += " " + stField + " |"
					sep += "---|"
				}
				stVal := reflect.ValueOf(entity)
				val := stVal.FieldByName(stField).Interface()
				tpe := stVal.FieldByName(stField).Type()

				if tpe.String() == "string" {
					val = strings.ReplaceAll(val.(string), "%", "%%")
				}

				row += " " + fmt.Sprintf("%v", val) + " |"
			}
		}

		rows = append(rows, row)
	}

	buf := ""
	for _, row := range rows {
		buf += "\n" + row
	}
	ret := fmt.Sprintf(
		"%s\n%s%s",
		header,
		sep,
		buf,
	)

	return ret
}

func ListToTUI[T any](entities []T, fields []string) string {
	var fieldsStr string
	var fieldsNr int

	if len(entities) == 0 {
		return ""
	}

	fieldsNr = len(fields)
	if fieldsNr == 0 {
		fieldsStr = ""
	} else {
		fieldsStr = strings.Join(fields, " ")
	}

	baseStyle := lipgloss.NewStyle().Padding(0, 1).
		BorderForeground(lipgloss.Color("240"))

	termWidth, _, err := term.GetSize(0)
	if err != nil {
		return ""
	}

	if termWidth < 80 {
		return "" // TODO: Error message
	}

	listedFields := []string{}
	rows := [][]string{}
	for eidx, entity := range entities {
		var row []string

		st := reflect.TypeOf(entity)
		tf := st.NumField()

		actualFieldsNr := (tf - 1)
		if fieldsNr > 0 {
			actualFieldsNr = fieldsNr
		}

		for i := 0; i < tf; i++ {

			stField := st.Field(i).Name

			if fieldsNr == 0 || strings.Contains(fieldsStr, stField) {
				stVal := reflect.ValueOf(entity)
				val := stVal.FieldByName(stField).Interface()
				tpe := stVal.FieldByName(stField).Type()

				if tpe.String() == "string" {
					val = strings.ReplaceAll(val.(string), "%", "%%")
				}
				size := int(math.Floor(float64(termWidth / actualFieldsNr)))

				row = append(row, ellipsis.Ending(fmt.Sprintf("%v", val), size))

				if eidx == 0 {
					listedFields = append(listedFields, stField)
				}
			}

		}

		rows = append(rows, row)
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("255"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row%2 == 1:
				return baseStyle.Foreground(lipgloss.Color("245"))
			default:
				return baseStyle.Foreground(lipgloss.Color("252"))
			}
		}).
		Headers(listedFields...).
		Rows(rows...)

	return t.String()
}
