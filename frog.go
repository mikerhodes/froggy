package froggy

func froggy(input string) bool {

	states := make([][]int, len(input))

	// The first move is always going to be 1 to the first place.
	states[0] = []int{1}

	// rock := '*'
	water := '~'

	for i, c := range input {
		// The frog can't start a hop from water
		if c == water {
			continue
		}

		// Check whether we've managed to hop here
		if len(states[i]) == 0 {
			continue
		}

		// Test whether the largest of the possible numbers
		// of hops from this location will take us
		// off the end of the river. This is our exit condition,
		// and if we never hit it, we can't find an answer.
		for _, state := range states[i] {
			if i+state+1 >= len(input) {
				return true
			}
		}

		// Otherwise, add the states that we can hop
		// to to our state map, then continue to the
		// next character
		for _, state := range states[i] {
			for _, hops := range []int{state - 1, state, state + 1} {
				if hops > 0 { // state might be 1
					loc := i + hops
					states[loc] = append(states[loc], hops)
				}
			}
		}
	}

	return false
}
