package froggy

import "testing"

type testCase struct {
	input    string
	expected bool
}

func TestFroggy(t *testing.T) {
	cases := []testCase{
		{"*", true},
		{"~", false},
		{"***********************", true},
		{"~~~~~~~~~~~~~~~~~~~~~~~", false},
		{"*~*", true},      // 1, 2
		{"*~~*", false},    // can't escape the water
		{"**~~*", false},   // can't escape the water
		{"***~~*", true},   // 1, 2, 3
		{"***~~~*", false}, // can't escape the water
		{"*~*~~*", true},   //  1, 2, 3
		{"**~*", true},     //  1, 1, 2
	}

	for _, c := range cases {
		if found := froggy(c.input); found != c.expected {
			t.Errorf("Input: %s ; expected %v, got %v", c.input, c.expected, found)
		}
	}
}
