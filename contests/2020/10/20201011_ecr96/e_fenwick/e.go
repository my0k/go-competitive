/*
URL:
https://codeforces.com/contest/1430/problem/E
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

func Reverse(A []rune) []rune {
	res := []rune{}

	n := len(A)
	for i := n - 1; i >= 0; i-- {
		res = append(res, A[i])
	}

	return res
}

var (
	n int
	S []rune

	C [30][]int
)

func main() {
	defer stdout.Flush()

	n = readi()
	S = readrs()
	revS := Reverse(S)

	for i, rr := range S {
		r := rr - 'a'
		C[r] = append(C[r], i)
	}

	A := []int{}
	for _, rr := range revS {
		r := rr - 'a'
		A = append(A, C[r][0])
		C[r] = C[r][1:]
	}

	ans := InversionNumber(A)

	fmt.Println(ans)
}

type FenwickTree struct {
	dat     []int
	n       int
	minPow2 int
}

// n(>=1) is number of elements of original data
func NewFenwickTree(n int) *FenwickTree {
	ft := new(FenwickTree)

	ft.dat = make([]int, n+1)
	ft.n = n

	ft.minPow2 = 1
	for {
		if (ft.minPow2 << 1) > n {
			break
		}
		ft.minPow2 <<= 1
	}

	return ft
}

// Add x to i.
// i is 0-index.
func (ft *FenwickTree) Add(i int, x int) {
	ft._add(i+1, x)
}

// RangeSum calculates a range sum of [l, r).
// l, r are 0-index.
func (ft *FenwickTree) RangeSum(l, r int) int {
	res := ft._cumsum(r) - ft._cumsum(l)

	return res
}

// LowerBound returns minimum i such that bit.Sum(i) >= w.
// LowerBound returns n, if the total sum is less than w.
func (ft *FenwickTree) LowerBound(w int) int {
	if w <= 0 {
		return 0
	}

	x := 0
	for k := ft.minPow2; k > 0; k /= 2 {
		if x+k <= ft.n && ft.dat[x+k] < w {
			w -= ft.dat[x+k]
			x += k
		}
	}

	return x
}

// private: Sum of [1, i](1-based)
func (ft *FenwickTree) _cumsum(i int) int {
	s := 0

	for i > 0 {
		s += ft.dat[i]
		i -= i & (-i)
	}

	return s
}

// private: Add x to i(1-based)
func (ft *FenwickTree) _add(i, x int) {
	for i <= ft.n {
		ft.dat[i] += x
		i += i & (-i)
	}
}

// InversionNumber returns 転倒数 of slice A.
// Time complexity: O(NlogN)
func InversionNumber(A []int) (res int64) {
	_max := func(integers ...int) int {
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

	res = 0

	maximum := _max(A...)
	ft := NewFenwickTree(maximum + 1)

	for i := 0; i < len(A); i++ {
		res += int64(i - ft.RangeSum(0, A[i]+1))
		ft.Add(A[i], 1)
	}

	return res
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