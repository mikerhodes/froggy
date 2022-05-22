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
		{"*~*", true},                       // 1, 2
		{"*~~*", false},                     // can't escape the water
		{"**~~*", false},                    // can't escape the water
		{"***~~*", true},                    // 1, 2, 3
		{"***~~~*", false},                  // can't escape the water
		{"*~*~~*", true},                    // 1, 2, 3
		{"**~*", true},                      // 1, 1, 2
		{"***~***********~~~~~****~", true}, // need 6 to cross water
		{"***~~*~~~*~~~~*~~~~~*~~~~", true}, // need 6 to cross water
		//1 2  3   4    5     6   <- position after each hop, adding 1 to length each time
		{"***~***********~~~~~~****~", false}, // need 7 to cross water
	}

	for _, c := range cases {
		if found := froggy([]byte(c.input)); found != c.expected {
			t.Errorf("Input: %s ; expected %v, got %v", c.input, c.expected, found)
		}
	}
}
