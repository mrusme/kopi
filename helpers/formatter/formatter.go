package formatter

import (
	"fmt"
	"reflect"
	"strings"
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

	var header = "|"
	var sep = "|"
	var rows []string
	for li, entity := range entities {
		var row = "|"

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
				stVal.FieldByName(stField).Type()

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
