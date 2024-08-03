package helpers

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func QueryArgRepeat(c int) string {
	var strsl []string

	for i := 0; i < c; i++ {
		strsl = append(strsl, "?")
	}

	return strings.Join(strsl, ",")
}

func IDsListValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile(`^(([0-9]+)( ){0,1}){0,}$`)
	return re.MatchString(value)
}
