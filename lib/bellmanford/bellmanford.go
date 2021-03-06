package bellmanford

type Edge struct {
	from, to, cost int
}

// Bellmanford calculates shortest paths from node sid.
// Node ids are 0-based.
// dist[i]: shortest path from sid to i
// roop[i]: whether node i is in a negative roop.
// Time complexity: O(|V| * |E|)
func Bellmanford(sid, n int, L []Edge) (dist []int, negaRoop []bool) {
	const BEL_INF = 1 << 60

	dist, negaRoop = make([]int, n), make([]bool, n)

	// Initialization
	for i := 0; i < n; i++ {
		dist[i], negaRoop[i] = BEL_INF, false
	}
	dist[sid] = 0

	// main algorithm (at least n-1 times updates)
	for c := 0; c < n; c++ {
		isUpdate := false

		for i := 0; i < len(L); i++ {
			e := L[i]
			if dist[e.from] != BEL_INF && dist[e.to] > dist[e.from]+e.cost {
				dist[e.to] = dist[e.from] + e.cost
				isUpdate = true
			}
		}

		if !isUpdate {
			break
		}
	}

	// check negative roops (at least n-1 times updates)
	for c := 0; c < n; c++ {
		for _, e := range L {
			// abnormal updates
			if dist[e.from] != BEL_INF && dist[e.to] > dist[e.from]+e.cost {
				dist[e.to] = dist[e.from] + e.cost
				negaRoop[e.to] = true
			}

			if negaRoop[e.from] {
				negaRoop[e.to] = true
			}
		}
	}

	return
}
