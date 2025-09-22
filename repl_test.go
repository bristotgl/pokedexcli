package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello, wonderful world!",
			expected: []string{"hello,", "wonderful", "world!"},
		},
		{
			input: "hElLO WORLd",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{},
		},

	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: %v vs %v", c.expected, actual)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%s) == %s, but should be %s", c.input, word, expectedWord)	
			}
		}
	}
}
