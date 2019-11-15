package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var notnum = regexp.MustCompile(`[^0-9]+`)
var phone = regexp.MustCompile(`^1?([2-9][0-9]{2}[2-9][0-9]{6})$`)

func clean(s string) (string, error) {
	cleaned := notnum.ReplaceAllString(s, "")
	if len(cleaned) > 11 || len(cleaned) < 10 {
		return "", errors.New("invalid phone number")
	}
	matches := phone.FindStringSubmatch(cleaned)
	if len(matches) == 2 {
		return matches[1], nil
	}
	return "", errors.New("invalid phone number")
}

// Number returns a cleaned version of phone number s
func Number(s string) (string, error) {
	num, err := clean(s)
	if err != nil {
		return "", err
	}
	return num, nil
}

// AreaCode returns the area code of phone number s
func AreaCode(s string) (string, error) {
	num, err := clean(s)
	if err != nil {
		return "", err
	}
	return num[0:3], nil
}

// Format returns a nicely formatted representation of phone number s
func Format(s string) (string, error) {
	num, err := clean(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:10]), nil

}
