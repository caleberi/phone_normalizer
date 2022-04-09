package utils

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var phoneRegex = "\\D"

func Normalize(phone string) string {
	var sb strings.Builder
	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

func RgNormalize(phone string) string {
	regx := regexp.MustCompile(phoneRegex)
	return regx.ReplaceAllString(phone, "")
}

func RemoveDuplicatePhoneNumbers(phones []string) []string {
	var keys []string
	m := make(map[string]bool)
	for _, phone := range phones {
		if _, ok := m[phone]; !ok {
			keys = append(keys, phone)
			m[phone] = true
		}
	}
	sort.Strings(keys)
	return keys
}
