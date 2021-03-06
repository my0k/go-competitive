/*
URL:
https://atcoder.jp/contests/abc167/tasks/abc167_c
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
	n, m, x int
	C       []int
	A       [][]int
)

func main() {
	defer stdout.Flush()

	n, m, x = readi3()
	C = make([]int, n)
	A = make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = readi()
		A[i] = readis(m)
	}

	ans := INF_B60
	BruteForceByBits01(n, func(B []int) {
		p := 0
		X := make([]int, m)
		for i, a := range B {
			if a == 0 {
				continue
			}
			p += C[i]
			for j := range A[i] {
				X[j] += A[i][j]
			}
		}

		for _, xx := range X {
			if xx < x {
				return
			}
		}

		if ans > p {
			ans = p
		}
	})

	if ans == INF_B60 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

// originated from ktateish@AtCoder
// e.g.: https://atcoder.jp/contests/abc167/submissions/13042361

// BruteForceByBits01(n, fn) calls fn with []int for each n-bit 0/1 pattern
// e.g.
// BruteForceByBits01(10, func(b []int) { fmt.Println(b) }
// -> [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
//    [1, 0, 0, 0, 0, 0, 0, 0, 0, 0]
//    [0, 1, 0, 0, 0, 0, 0, 0, 0, 0]
//    [1, 1, 0, 0, 0, 0, 0, 0, 0, 0]
//    [0, 0, 1, 0, 0, 0, 0, 0, 0, 0]
//    ...
func BruteForceByBits01(n int, fn func([]int)) {
	N := 1 << uint(n)
	a := make([]int, n)
	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			k := n - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = 0
			} else {
				a[k] = 1
			}
		}
		fn(a)
	}
}

// BruteForceByBitsTF(n, fn) calls fn with []bool for each n-bit true/false pattern
// e.g.
// BruteForceByBitsTF(10, func(b []bool) { fmt.Println(b) }
// -> [false, false, false, false, false, false, false, false, false, false]
//    [true, false, false, false, false, false, false, false, false, false]
//    [false, true, false, false, false, false, false, false, false, false]
//    ...
func BruteForceByBitsTF(n int, fn func([]bool)) {
	N := 1 << uint(n)
	a := make([]bool, n)
	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			k := n - j - 1
			if i&(1<<uint(j)) == 0 {
				a[k] = false
			} else {
				a[k] = true
			}
		}
		fn(a)
	}
}

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
