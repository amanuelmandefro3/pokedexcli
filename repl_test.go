package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"  hello world  ", []string{"hello", "world"}},
		{"\tfoo bar\t", []string{"foo", "bar"}},
		{"\n\nbaz qux\n\n", []string{"baz", "qux"}},
		{"   multiple   spaces   ", []string{"multiple", "spaces"}},
		{"Charmander Bulbasaur PIKACHU", []string{"charmander", "bulbasaur", "pikachu"}},
		{"explore pastoria-city-area", []string{"explore", "pastoria-city-area"}},
		{"catch pikachu", []string{"catch", "pikachu"}},
		{"inspect pikachu", []string{"inspect", "pikachu"}},
	}

	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", result, c.expected)
			continue
		}
		for i := range result {
			if result[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) == %q, expected %q", c.input, result, c.expected)
			}
		}
	}
}
