package strTools

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct{ input, expected_output string }{
		{"hello!", "!olleh"},
		{"Input", "tupnI"},
		{"", ""},
	}
	for _, testcase := range cases {
		output := ReverseString(testcase.input)
		if output != testcase.expected_output {
			t.Errorf("ReverseString(%q) == %q, expected: %q", testcase.input, output, testcase.expected_output)
		}
	}
}
