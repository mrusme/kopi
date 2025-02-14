package helpers

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
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

func Validate(val *validator.Validate, entity interface{}) error {
	if err := val.Struct(entity); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

func ValidateField(val *validator.Validate, entity interface{}, field string) error {
	if err := val.Struct(entity); err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			if er.StructField() == field {
				return er
			}
		}
	}
	return nil
}

func ParsePrice(s string) (int64, string, error) {
	re := regexp.MustCompile(`(\d{1,3}(?:[\,\.]\d{3})*(?:[\.,]\d{1,2})?)\s*([a-zA-Z]{3})|\b([a-zA-Z]{3})\s*(\d{1,3}(?:,\d{3})*(?:\.\d{2})?)\b`)
	match := re.FindStringSubmatch(s)
	if len(match) < 3 {
		return 0, "", errors.New("Please enter a price in the format '10.50 XYZ'")
	}

	value := match[1]
	currency := match[2]

	dotIdx := strings.Index(value, ".")
	commaIdx := strings.Index(value, ",")
	if dotIdx < commaIdx {
		// If . appears before , or doesn't appear at all, we are using
		// 1.000,00 or 10,00
		value = strings.Replace(value, ".", "", -1)
		value = strings.Replace(value, ",", "", 1)
	} else if dotIdx > commaIdx {
		// Otherwise we are using
		// 1,000.00 or 10.00
		value = strings.Replace(value, ".", "", 1)
		value = strings.Replace(value, ",", "", -1)
	} else if dotIdx == commaIdx && dotIdx == -1 {
		// If both are not to be found, we are using
		// 12
		value = value + "00"
	}
	cents, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, "", err
	}

	return cents, strings.ToUpper(currency), nil
}

func HandleFormError(err error) {
	if err == nil {
		return
	}

	if errors.Is(err, huh.ErrUserAborted) {
		os.Exit(1)
	}

	// TODO: Else?
	return
}
