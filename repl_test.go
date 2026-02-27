package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"empty input":     {input: "", expected: []string{}},
		"clean input":     {input: "hello world", expected: []string{"hello", "world"}},
		"trailing spaces": {input: " hello world ", expected: []string{"hello", "world"}},
		"table spacing":   {input: "hello\tworld", expected: []string{"hello", "world"}},
		"new line":        {input: "\nhello\nworld\n", expected: []string{"hello", "world"}},
		"uppercase word":  {input: "HELLO World", expected: []string{"hello", "world"}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if len(c.expected) != len(actual) {
				t.Errorf("lenghts don't match: got %#v but want %#v", actual, c.expected)
				return
			}

			for i := range actual {
				actualWord := actual[i]
				expectedWord := c.expected[i]
				if actualWord != expectedWord {
					t.Errorf("cleanInput(%#v) == %#v, expected: %#v", c.input, actual, c.expected)
				}
			}
		})
	}
}
