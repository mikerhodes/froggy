// This solution to the froggy problem takes the approach
// of setting up states dangling off each letter. As we
// progress across the river, see where we can get to with
// our next set of hops. Store the number of hops as a
// possible "state" at that location. At each location,
// process all the states we find there (i.e., each of
// the prior number of hops we took to get there).
//
// Here's the states we have at each point for the
// spaces further on than where we are looking at:
//  v
// "***~~*"
//  1
//   v
// "***~~*"
//   12     <- In the previous step, you could hop 1 or 2
//             forward, so you get 1 in the
//             2nd location and 2 in the third location.
//    v
// "***~~*"
//    22   <-- that second 2 is actually invalid, because it's in the river
//    1
//       v
// "***~~*"
//       3 <-- you have to jump 3 from state 2, to not fall
//
// "***~~*"
//           4  <- longest hop takes you off the end of the river

package froggy

func froggy(input []byte) bool {

	states := make([][]int, len(input))

	// The first move is always going to be 1 to the first place.
	states[0] = []int{1}

	// rock := '*'
	water := byte('~')

	for i, c := range input {
		// The frog can't start a hop from water
		if c == water {
			continue
		}

		for _, state := range states[i] {
			// Test whether the largest possible hop from this
			// location will take us off the end of the river.
			// This is our exit condition, and if we never hit
			// it, we can't find an answer.
			if i+state+1 >= len(input) {
				return true
			}

			// Otherwise, add the states that we can hop
			// to to our state map, then continue to the
			// next character
			for _, hops := range []int{state - 1, state, state + 1} {
				if hops > 0 { // state might be 1
					loc := i + hops
					// there's a question here of how likely
					// we are to end up with duplicates by
					// not checking for whether we already have
					// this state. And if we don't, we'll just
					// amplify that as we progress.
					states[loc] = append(states[loc], hops)
				}
			}
		}
	}

	return false
}
