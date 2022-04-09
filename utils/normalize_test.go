package utils_test

import (
	"phone/utils"
	"reflect"
	"sort"
	"testing"
)

func TestRemoveDuplicatePhoneNumbers(t *testing.T) {

	var phones = []string{
		"123 456 7891", "(123) 456 7892", "(123) 456-7893",
		"123-456-7894", "123-456-7890", "1234567892", "(123)456-7892",
	}

	var normalizedPhones []string

	for _, phone := range phones {
		normalizedPhones = append(normalizedPhones, utils.Normalize(phone))
	}

	normalizedPhones = utils.RemoveDuplicatePhoneNumbers(normalizedPhones)

	valid := []string{"1234567890", "1234567891", "1234567892", "1234567893", "1234567894"}
	sort.Strings(valid)
	if !reflect.DeepEqual(valid, normalizedPhones) {
		t.Errorf("Not properly normalized or contain duplicate [ expect : %v , got :%v]", valid, normalizedPhones)
	}

}

func TestNormalize(t *testing.T) {

	testcases := []struct {
		input  string
		expect string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			actual := utils.Normalize(tc.input)
			if actual != tc.expect {
				t.Errorf("got %s : want %s", actual, tc.expect)
			}
		})
	}
}

func TestRgNormalize(t *testing.T) {

	testcases := []struct {
		input  string
		expect string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			actual := utils.RgNormalize(tc.input)
			if actual != tc.expect {
				t.Errorf("got %s : want %s", actual, tc.expect)
			}
		})
	}
}
