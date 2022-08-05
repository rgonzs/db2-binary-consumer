package utils

import (
	"time"
)

func Utf8Decode(str string) string {
	var result string
	for i := range str {
		result += string(str[i])
	}
	return result
}

func ParseDateToLocal(str string) (time.Time, error) {
	date, err := time.ParseInLocation(time.RFC3339, str, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
