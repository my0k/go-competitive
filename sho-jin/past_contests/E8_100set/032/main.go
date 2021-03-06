/*
URL:
https://onlinejudge.u-aizu.ac.jp/problems/1166
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	w, h int
	V, H [][]int

	G [][]int
)

func main() {
	for {
		w, h = ReadInt2()
		V, H = [][]int{}, [][]int{}
		if w == 0 && h == 0 {
			return
		}

		for i := 0; i < 2*h-1; i++ {
			if i%2 == 0 {
				row := ReadIntSlice(w - 1)
				V = append(V, row)
			} else {
				row := ReadIntSlice(w)
				H = append(H, row)
			}
		}

		steps := [4][2]int{
			[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0},
		}
		toid := func(i, j int) int { return w*i + j }
		// toy := func(id int) int { return id / w }
		// tox := func(id int) int { return id % w }

		N := h * w
		G = make([][]int, N)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				cid := toid(i, j)

				for _, s := range steps {
					dy, dx := s[0], s[1]
					ny, nx := i+dy, j+dx
					if 0 <= ny && ny < h && 0 <= nx && nx < w {
						ok := true

						if dy == 0 {
							if (dx == 1 && vCheck(i, j) && V[i][j] == 1) || (dx == -1 && vCheck(i, j-1) && V[i][j-1] == 1) {
								ok = false
							}
						} else {
							if (dy == 1 && hCheck(i, j) && H[i][j] == 1) || (dy == -1 && hCheck(i-1, j) && H[i-1][j] == 1) {
								ok = false
							}
						}

						if ok {
							nid := toid(ny, nx)
							G[cid] = append(G[cid], nid)
						}
					}
				}
			}
		}

		sid, gid := toid(0, 0), toid(h-1, w-1)
		dp, _ := SSSPByBFS(sid, N, G[:N])
		if dp[gid] >= INF_BIT30 {
			fmt.Println(0)
		} else {
			fmt.Println(dp[gid] + 1)
		}
	}
}

func hCheck(y, x int) bool {
	return 0 <= y && y < h-1 && 0 <= x && x < w
}

func vCheck(y, x int) bool {
	return 0 <= y && y < h && 0 <= x && x < w-1
}

// verified by https://codeforces.com/contest/1320/problem/B
func SSSPByBFS(sid, n int, AG [][]int) (dp []int, visited []bool) {
	dp = make([]int, n)
	visited = make([]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = INF_BIT30
		visited[i] = false
	}

	Q := []Node{}
	dp[sid] = 0
	visited[sid] = true
	Q = append(Q, Node{id: sid, cost: dp[sid]})

	for len(Q) > 0 {
		cnode := Q[0]
		Q = Q[1:]

		for _, nid := range G[cnode.id] {
			// 訪問済みならパス
			if visited[nid] {
				continue
			}

			dp[nid] = cnode.cost + 1
			visited[nid] = true
			Q = append(Q, Node{id: nid, cost: dp[nid]})
		}
	}

	return dp, visited
}

type Node struct {
	id, cost int
}

/*******************************************************************/

/********** common constants **********/

const (
	// General purpose
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPHABET_NUM = 26
	INF_INT64    = math.MaxInt64
	INF_BIT60    = 1 << 60
	INF_INT32    = math.MaxInt32
	INF_BIT30    = 1 << 30
	NIL          = -1

	// for dijkstra, prim, and so on
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

/********** FAU standard libraries **********/

//fmt.Sprintf("%b\n", 255) 	// binary expression

/********** I/O usage **********/

//str := ReadString()
//i := ReadInt()
//X := ReadIntSlice(n)
//S := ReadRuneSlice()
//a := ReadFloat64()
//A := ReadFloat64Slice(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*********** Input ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func newReadString(ior io.Reader, sf bufio.SplitFunc) func() string {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	r.Split(sf)

	return func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
}

// ReadInt returns an integer.
func ReadInt() int {
	return int(readInt64())
}
func ReadInt2() (int, int) {
	return int(readInt64()), int(readInt64())
}
func ReadInt3() (int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64())
}
func ReadInt4() (int, int, int, int) {
	return int(readInt64()), int(readInt64()), int(readInt64()), int(readInt64())
}

// ReadInt64 returns as integer as int64.
func ReadInt64() int64 {
	return readInt64()
}
func ReadInt64_2() (int64, int64) {
	return readInt64(), readInt64()
}
func ReadInt64_3() (int64, int64, int64) {
	return readInt64(), readInt64(), readInt64()
}
func ReadInt64_4() (int64, int64, int64, int64) {
	return readInt64(), readInt64(), readInt64(), readInt64()
}

func readInt64() int64 {
	i, err := strconv.ParseInt(ReadString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// ReadIntSlice returns an integer slice that has n integers.
func ReadIntSlice(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt()
	}
	return b
}

// ReadInt64Slice returns as int64 slice that has n integers.
func ReadInt64Slice(n int) []int64 {
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadInt64()
	}
	return b
}

// ReadFloat64 returns an float64.
func ReadFloat64() float64 {
	return float64(readFloat64())
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(ReadString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// ReadFloatSlice returns an float64 slice that has n float64.
func ReadFloat64Slice(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = ReadFloat64()
	}
	return b
}

// ReadRuneSlice returns a rune slice.
func ReadRuneSlice() []rune {
	return []rune(ReadString())
}

/*********** Output ***********/

// PrintIntsLine returns integers string delimited by a space.
func PrintIntsLine(A ...int) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		str := strconv.Itoa(A[i])
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

// PrintIntsLine returns integers string delimited by a space.
func PrintInts64Line(A ...int64) string {
	res := []rune{}

	for i := 0; i < len(A); i++ {
		str := strconv.FormatInt(A[i], 10) // 64bit int version
		res = append(res, []rune(str)...)

		if i != len(A)-1 {
			res = append(res, ' ')
		}
	}

	return string(res)
}

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

/*********** Debugging ***********/

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// ZeroPaddingRuneSlice returns binary expressions of integer n with zero padding.
// For debugging use.
func ZeroPaddingRuneSlice(n, digitsNum int) []rune {
	sn := fmt.Sprintf("%b", n)

	residualLength := digitsNum - len(sn)
	if residualLength <= 0 {
		return []rune(sn)
	}

	zeros := make([]rune, residualLength)
	for i := 0; i < len(zeros); i++ {
		zeros[i] = '0'
	}

	res := []rune{}
	res = append(res, zeros...)
	res = append(res, []rune(sn)...)

	return res
}
