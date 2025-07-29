package common

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello     world      ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World Go LanG",
			expected: []string{"hello", "world", "go", "lang"},
		},
		{
			input:    " extra spaces here ",
			expected: []string{"extra", "spaces", "here"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "TEST",
			expected: []string{"test"},
		},
		{
			input:    "G0 1s Fun! @#$%",
			expected: []string{"g0", "1s", "fun!", "@#$%"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected len of words %d; got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]

			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected %s; actual %s", expectedWord, word)
			}
		}
	}
}
