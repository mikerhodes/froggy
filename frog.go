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

	states := make([]map[int]bool, len(input))
	for i := 0; i < len(input); i++ {
		states[i] = make(map[int]bool)
	}

	// The first move is always going to be 1 to the first place,
	// so seed the states list with that first move.
	states[0][1] = true

	// rock := '*'
	water := byte('~')

	for i, c := range input {
		// The frog can't start a hop from water
		if c == water {
			continue
		}

		for state, _ := range states[i] {
			// Test whether the largest possible hop from this
			// location will take us off the end of the river.
			// This is our exit condition, and if we never hit
			// it, we can't find an answer.
			if i+state+1 >= len(input) {
				return true
			}

			// We can hop one fewer places, the same places or one
			// more space. Figure out where that puts us on the
			// path, and store the new state (i.e., the (location,hops)
			// pair, though the location is implicit in where we
			// maintain the states set in the states array.
			for _, hops := range []int{state - 1, state, state + 1} {
				if hops > 0 { // state might be 1
					states[i+hops][hops] = true
				}
			}
		}
	}

	return false
}
