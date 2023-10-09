package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator that stores errors
//
//	Validator.CheckField(Validator.NotBlank("abc"), "name", "name should be not blank")
//	Validator.Valid()
type Validator struct {
	NonFieldErrors []string
	FieldErrors    map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = map[string]string{}
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func NotBlank(v string) bool {
	return strings.TrimSpace(v) != ""
}

func MaxChars(v string, n int) bool {
	return utf8.RuneCountInString(v) <= n
}

func MinChars(v string, n int) bool {
	return utf8.RuneCountInString(v) >= n
}

func Matches(v string, rx *regexp.Regexp) bool {
	return rx.MatchString(v)
}

func PermittedValue[T comparable](v T, permittedVals ...T) bool {
	for i := range permittedVals {
		if v == permittedVals[i] {
			return true
		}
	}
	return false
}
