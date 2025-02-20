package fuma

import (
	"errors"
	"math"
	"reflect"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func FindMatch[T any](slice *[]T, fieldName string, matchValue string) (*T, error) {
	sliceValue := reflect.ValueOf(*slice)

	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("Input is not a slice")
	}

	var highestRank int = -1
	var highestRanked *T = nil
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		if item.Kind() == reflect.Struct {
			field := item.FieldByName(fieldName)

			if field.IsValid() && field.Kind() == reflect.String {
				r := GetRank(matchValue, field.String())
				if r > 0 && r > highestRank {
					highestRank = r
					highestRanked = item.Addr().Interface().(*T)
				}
			} else {
				return nil, errors.New("Field " + fieldName +
					" does not exist in struct")
			}
		} else {
			return nil, errors.New("Input slice doesn't contain structs")
		}
	}

	if highestRank > 0 && highestRanked != nil {
		return highestRanked, nil
	}
	return nil, nil
}

func GetRank(target string, query string) int {
	tL := strings.ReplaceAll(strings.ToLower(target), " ", "")
	qL := strings.ReplaceAll(strings.ToLower(query), " ", "")
	queryLength := len(qL)

	distance := fuzzy.LevenshteinDistance(tL, qL)
	// fmt.Printf("%s : %s = %d\n", target, query, distance)

	threshold := int(math.Floor(float64(queryLength) / 2.0))
	rank := threshold - distance
	if rank < 0 {
		rank = -1
	}

	return rank
}
