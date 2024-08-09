package bagListCmd

import (
	"fmt"
	"reflect"
	"strings"

	bagLabel "github.com/mrusme/kopi/bag/label"
	"github.com/mrusme/kopi/helpers/out"
)

func mdList(labels []bagLabel.Label, fields []string) {
	var fieldsStr string

	if len(fields) == 0 {
		fieldsStr = ""
	} else {
		fieldsStr = strings.Join(fields, " ")
	}

	var header = "|"
	var sep = "|"
	var rows []string
	for li, label := range labels {
		var row = "|"

		st := reflect.TypeOf(label)
		for i := 0; i < st.NumField(); i++ {

			stField := st.Field(i).Name

			if len(fieldsStr) == 0 || strings.Contains(fieldsStr, stField) {
				if li == 0 {
					header += " " + stField + " |"
					sep += "---|"
				}
				stVal := reflect.ValueOf(label)
				val := stVal.FieldByName(stField).Interface()
				stVal.FieldByName(stField).Type()

				row += " " + fmt.Sprintf("%v", val) + " |"
			}
		}

		rows = append(rows, row)
	}

	out.Put("%s", header)
	out.Put("%s", sep)
	for _, row := range rows {
		out.Put("%s", row)
	}
}
