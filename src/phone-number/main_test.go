package main

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"04-456-78920", "0445678920"},
		{"(04) 2456 7890", "0424567890"},
		{"043.456.7890", "0434567890"},
		{"04334567890", "04334567890"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := normalize(test.input)
			if result != test.expected {
				t.Errorf("normalize(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}
