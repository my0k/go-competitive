/*
URL:
https://atcoder.jp/contests/joisc2014/tasks/joisc2014_h
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
	n int
	S []rune

	J, O, I []int
	JO, JI  []int

	A, B []V

	P [500000 + 50][]int

	memo map[Coord][]int
)

type V struct {
	v, idx int
}

type Coord struct {
	jo, ji int
}

func main() {
	defer stdout.Flush()

	n = readi()
	S = readrs()

	J, O, I = make([]int, n+1), make([]int, n+1), make([]int, n+1)
	JO, JI = make([]int, n+1), make([]int, n+1)

	for i := 0; i < n; i++ {
		J[i+1] = J[i]
		O[i+1] = O[i]
		I[i+1] = I[i]
		if S[i] == 'J' {
			J[i+1]++
		} else if S[i] == 'O' {
			O[i+1]++
		} else {
			I[i+1]++
		}
	}

	for i := 0; i <= n; i++ {
		JO[i] = J[i] - O[i]
		JI[i] = J[i] - I[i]
	}
	debugf("JO: %v\n", JO)
	debugf("JI: %v\n", JI)

	memo = make(map[Coord][]int)
	for i := 0; i <= n; i++ {
		c := Coord{JO[i], JI[i]}
		if _, ok := memo[c]; !ok {
			memo[c] = make([]int, 0)
		}
		memo[c] = append(memo[c], i)
	}

	ans := 0
	for _, P := range memo {
		if len(P) > 1 {
			ChMax(&ans, P[len(P)-1]-P[0])
		}
	}
	fmt.Println(ans)

	// for i := 0; i <= n; i++ {
	// 	A = append(A, V{JO[i], i})
	// 	B = append(B, V{JI[i], i})
	// }
	// sort.Slice(A, func(i, j int) bool {
	// 	if A[i].v < A[j].v {
	// 		return true
	// 	} else if A[i].v > A[j].v {
	// 		return false
	// 	}
	// 	return A[i].idx < A[j].idx
	// })
	// sort.Slice(B, func(i, j int) bool {
	// 	if B[i].v < B[j].v {
	// 		return true
	// 	} else if B[i].v > B[j].v {
	// 		return false
	// 	}
	// 	return B[i].idx < B[j].idx
	// })
	// debugf("A: %v\n", A)
	// debugf("B: %v\n", B)

}

// ChMax accepts a pointer of integer and a target value.
// If target value is LARGER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMax(updatedValue *int, target int) bool {
	if *updatedValue < target {
		*updatedValue = target
		return true
	}
	return false
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
)

// modi can calculate a right residual whether value is positive or negative.
func modi(val, m int) int {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

// modll can calculate a right residual whether value is positive or negative.
func modll(val, m int64) int64 {
	res := val % m
	if res < 0 {
		res += m
	}
	return res
}

/********** bufio setting **********/

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	reads = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
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
