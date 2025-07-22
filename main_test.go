package main_test

import (
	"testing"

	gount "github.com/ComicShrimp/gount"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "new lines",
			input: "one two\nthree four",
			wants: 4,
		},
		{
			name:  "multiple spaces",
			input: "one two  three four",
			wants: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := gount.CountWords([]byte(tc.input))

			if result != tc.wants {
				t.Logf("Expected: %v | Got: %v", tc.wants, result)
				t.Fail()
			}
		})
	}
}
