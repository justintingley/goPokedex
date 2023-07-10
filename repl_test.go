package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Expected %v, got %v", len(cs.expected), len(actual))
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("Expected %v, got %v", cs.expected[i], actual[i])
			}
		}
	}

}
