package tsort

// TSort returns a node ids list in topological order.
// node id is 0-based.
// Time complexity: O(|E| + |V|)
func TSort(nn int, AG [][]int) (ok bool, tsortedIDs []int) {
	tsortedIDs = []int{}

	inDegrees := make([]int, nn)
	for s := 0; s < nn; s++ {
		for _, t := range AG[s] {
			inDegrees[t]++
		}
	}

	stack := []int{}
	for nid := 0; nid < nn; nid++ {
		if inDegrees[nid] == 0 {
			stack = append(stack, nid)
		}
	}

	for len(stack) > 0 {
		cid := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		tsortedIDs = append(tsortedIDs, cid)

		for _, nid := range AG[cid] {
			inDegrees[nid]--
			if inDegrees[nid] == 0 {
				stack = append(stack, nid)
			}
		}
	}

	if len(tsortedIDs) != nn {
		return false, nil
	}

	return true, tsortedIDs
}

// LongestPath returns a length of longest path of a given graph.
// This function assumes that all costs of edges are 1.
// Time complexity: O(|E| + |V|)
func LongestPath(tsortedIDs []int, AG [][]int) (maxLength int, dp []int) {
	_chmax := func(updatedValue *int, target int) bool {
		if *updatedValue < target {
			*updatedValue = target
			return true
		}
		return false
	}

	dp = make([]int, len(tsortedIDs)+1)

	for i := 0; i < len(tsortedIDs); i++ {
		cid := tsortedIDs[i]
		for _, nid := range AG[cid] {
			_chmax(&dp[nid], dp[cid]+1)
		}
	}

	maxLength = 0
	for i := 0; i < len(tsortedIDs); i++ {
		_chmax(&maxLength, dp[i])
	}

	return maxLength, dp
}