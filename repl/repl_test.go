package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected []string
	} {
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "hello\tworld",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Errorf("slices of different sizes!")
		}

		// continue...
	}
}