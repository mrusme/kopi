package helpers

import "strings"

func QueryArgRepeat(c int) string {
	var strsl []string

	for i := 0; i < c; i++ {
		strsl = append(strsl, "?")
	}

	return strings.Join(strsl, ",")
}
