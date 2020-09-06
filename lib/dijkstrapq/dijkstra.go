package dijkstrapq

import "container/heap"

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

// DP value type
type V struct {
	// {{
	gas, times int
	// }}
}

// weight of edge
type EdgeWeight struct {
	// {{
	gas int
	// }}
}

// edge of graph
type Edge struct {
	to int
	ew EdgeWeight
}

// for initializing start points of dijkstra algorithm
type StartPoint struct {
	id    int
	vzero V
}

type DijkstraSolver struct {
	vinf  V
	Less  func(l, r V) bool          // Less returns l < r, and shared with pq.
	NextV func(cv *Vertex, e Edge) V // NextV returns next value considered by transition.
}

func NewDijkstraSolver(
	vinf V, Less func(l, r V) bool, NextV func(cv *Vertex, e Edge) V,
) *DijkstraSolver {
	ds := new(DijkstraSolver)

	// shared with priority queue
	__less = Less

	ds.vinf, ds.Less, ds.NextV = vinf, Less, NextV

	return ds
}

// verified by [ABC143-E](https://atcoder.jp/contests/abc143/tasks/abc143_e)
func (ds *DijkstraSolver) Dijkstra(S []StartPoint, n int, AG [][]Edge) []V {
	// initialize data
	dp, colors := ds.initAll(n)

	// configure about start points (some problems have multi start points)
	pq := ds.initStartPoint(S, dp, colors)

	// body of dijkstra algorithm
	for pq.Len() > 0 {
		pop := pq.pop()
		colors[pop.id] = BLACK
		if ds.Less(dp[pop.id], pop.v) {
			continue
		}

		// to next node
		for _, e := range AG[pop.id] {
			if colors[e.to] == BLACK {
				continue
			}

			// update optimal value of the next node
			nv := ds.NextV(pop, e)

			if ds.Less(nv, dp[e.to]) {
				dp[e.to] = nv
				pq.push(&Vertex{id: e.to, v: nv})
				colors[e.to] = GRAY
			}
		}
	}

	return dp
}

// initAll returns initialized dp and colors slices.
func (ds *DijkstraSolver) initAll(n int) (dp []V, colors []int) {
	dp, colors = make([]V, n), make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = ds.vinf
		colors[i] = WHITE
	}

	return dp, colors
}

// initStartPoint returns initialized priority queue, and update dp and colors slices.
// *This function update arguments (side effects).*
func (ds *DijkstraSolver) initStartPoint(S []StartPoint, dp []V, colors []int) *VertexPQ {
	pq := NewVertexPQ()

	for _, sp := range S {
		pq.push(&Vertex{id: sp.id, v: sp.vzero})
		dp[sp.id] = sp.vzero
		colors[sp.id] = GRAY
	}

	return pq
}

// Less function is shared with a priority queue.
var __less func(l, r V) bool

// Definitions of a priority queue
type Vertex struct {
	id int
	v  V
}
type VertexPQ []*Vertex

func NewVertexPQ() *VertexPQ {
	temp := make(VertexPQ, 0)
	pq := &temp
	heap.Init(pq)

	return pq
}
func (pq *VertexPQ) push(target *Vertex) {
	heap.Push(pq, target)
}
func (pq *VertexPQ) pop() *Vertex {
	return heap.Pop(pq).(*Vertex)
}

func (pq VertexPQ) Len() int { return len(pq) }
func (pq VertexPQ) Less(i, j int) bool {
	return __less(pq[i].v, pq[j].v)
}
func (pq VertexPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *VertexPQ) Push(x interface{}) {
	item := x.(*Vertex)
	*pq = append(*pq, item)
}
func (pq *VertexPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}