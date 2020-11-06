/*
URL:

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
	h, w           int
	sx, sy, tx, ty int
	S              [][]rune

	C [1000 + 50][1000 + 50]int

	steps [4][2]int

	next     [4][4][3]int // i: 今向いている方向, j: 進みたい方向 -> dy, dx
	priority [4]int
)

const (
	L, U, R, D = 0, 1, 2, 3
)

func main() {
	defer stdout.Flush()

	steps = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	next[U][L] = [3]int{0, -1, L}
	next[U][U] = [3]int{-1, 0, U}
	next[U][R] = [3]int{0, 1, R}
	next[U][D] = [3]int{1, 0, D}

	next[L][L] = [3]int{1, 0, D}
	next[L][U] = [3]int{0, -1, L}
	next[L][R] = [3]int{-1, 0, U}
	next[L][D] = [3]int{0, 1, R}

	next[D][L] = [3]int{0, 1, R}
	next[D][U] = [3]int{1, 0, D}
	next[D][R] = [3]int{0, -1, L}
	next[D][D] = [3]int{-1, 0, U}

	next[R][L] = [3]int{-1, 0, U}
	next[R][U] = [3]int{0, 1, R}
	next[R][R] = [3]int{1, 0, D}
	next[R][D] = [3]int{0, -1, L}

	priority = [4]int{L, U, R, D}

	h, w = readi2()
	// sx, sy, tx, ty = readi4()
	sy, sx, ty, tx = readi4()
	sx--
	sy--
	tx--
	ty--
	for i := 0; i < h; i++ {
		row := readrs()
		S = append(S, row)
	}

	ans := solve()

	fmt.Println(ans)
}

func solve() int {
	// 周囲が全て壁の場合は例外処理
	wall := 0
	for _, s := range steps {
		dy, dx := s[0], s[1]
		ny, nx := sy+dy, sx+dx
		if isOutofGrid(ny, nx) || S[ny][nx] == '#' {
			wall++
		}
	}
	if wall == 4 {
		// debugf("first\n")
		return -1
	}

	curState := State{y: sy, x: sx, time: 0, dir: U}
	for {
		// ゴールに到達
		if curState.y == ty && curState.x == tx {
			break
		}
		// 異常終了
		if C[curState.y][curState.x] >= 10 {
			break
		}

		debugf("curState: %v\n", curState)

		// カウントアップ
		C[curState.y][curState.x]++

		for _, d := range priority {
			dy, dx := next[curState.dir][d][0], next[curState.dir][d][1]
			ny, nx := curState.y+dy, curState.x+dx

			if isOutofGrid(ny, nx) || S[ny][nx] == '#' {
				continue
			}

			curState.y = ny
			curState.x = nx
			curState.time++
			curState.dir = next[curState.dir][d][2]

			break
		}
	}

	// debugf("curState: %v\n", curState)
	if C[curState.y][curState.x] >= 10 {
		return -1
	}

	return curState.time
}

type State struct {
	y, x int
	time int
	dir  int // LRUD
}

func isOutofGrid(y, x int) bool {
	return !(0 <= y && y < h && 0 <= x && x < w)
}

// steps := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
// toid := func(i, j int) int { return w*i + j }
// toy := func(id int) int { return id / w }
// tox := func(id int) int { return id % w }

// N := h * w
// G = make([][]int, N)
// for i := 0; i < h; i++ {
// 	for j := 0; j < w; j++ {
// 		cid := toid(i, j)

// 		for _, s := range steps {
// 			dy, dx := s[0], s[1]
// 			ny, nx := i+dy, j+dx
// 			if 0 <= ny && ny < h && 0 <= nx && nx < w {
// 				nid := toid(ny, nx)

// 				G[cid] = append(G[cid], nid)
// 			}
// 		}
// 	}
// }

/*******************************************************************/

/********** common constants **********/

const (
	MOD = 1000000000 + 7
	// MOD          = 998244353
	ALPH_N  = 26
	INF_I64 = math.MaxInt64
	INF_B60 = 1 << 60
	INF_I32 = math.MaxInt32
	INF_B30 = 1 << 30
	NIL     = -1
	EPS     = 1e-10
)

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

// mod can calculate a right residual whether value is positive or negative.
func mod(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// min returns the min integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

// max returns the max integer among input set.
// This function needs at least 1 argument (no argument causes panic).
func max(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m < integer {
			m = integer
		}
	}
	return m
}

// chmin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func chmin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// chmax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func chmax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
}

// sum returns multiple integers sum.
func sum(integers ...int) int {
	var s int
	s = 0

	for _, i := range integers {
		s += i
	}

	return s
}

// abs is integer version of math.Abs
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/********** FAU standard libraries **********/

//fmt.Sprintf("%b\n", 255) 	// binary expression

/********** I/O usage **********/

//str := reads()
//i := readi()
//X := readis(n)
//S := readrs()
//a := readf()
//A := readfs(n)

//str := ZeroPaddingRuneSlice(num, 32)
//str := PrintIntsLine(X...)

/*********** Input ***********/

var (
	// reads returns a WORD string.
	reads  func() string
	stdout *bufio.Writer
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

// readi returns an integer.
func readi() int {
	return int(_readInt64())
}
func readi2() (int, int) {
	return int(_readInt64()), int(_readInt64())
}
func readi3() (int, int, int) {
	return int(_readInt64()), int(_readInt64()), int(_readInt64())
}
func readi4() (int, int, int, int) {
	return int(_readInt64()), int(_readInt64()), int(_readInt64()), int(_readInt64())
}

// readll returns as integer as int64.
func readll() int64 {
	return _readInt64()
}
func readll2() (int64, int64) {
	return _readInt64(), _readInt64()
}
func readll3() (int64, int64, int64) {
	return _readInt64(), _readInt64(), _readInt64()
}
func readll4() (int64, int64, int64, int64) {
	return _readInt64(), _readInt64(), _readInt64(), _readInt64()
}

func _readInt64() int64 {
	i, err := strconv.ParseInt(reads(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

// readis returns an integer slice that has n integers.
func readis(n int) []int {
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = readi()
	}
	return b
}

// readlls returns as int64 slice that has n integers.
func readlls(n int) []int64 {
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		b[i] = readll()
	}
	return b
}

// readf returns an float64.
func readf() float64 {
	return float64(_readFloat64())
}

func _readFloat64() float64 {
	f, err := strconv.ParseFloat(reads(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

// ReadFloatSlice returns an float64 slice that has n float64.
func readfs(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = readf()
	}
	return b
}

// readrs returns a rune slice.
func readrs() []rune {
	return []rune(reads())
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

// Printf is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func printf(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}

/*********** Debugging ***********/

// debugf is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func debugf(format string, a ...interface{}) {
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