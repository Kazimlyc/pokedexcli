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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    " selam dunya",
			expected: []string{"selam", "dunya"},
		}, {
			input:    " selam dunya naber",
			expected: []string{"selam", "dunya", "naber"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("fail")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("fail")
				break
			}
		}
	}

}
