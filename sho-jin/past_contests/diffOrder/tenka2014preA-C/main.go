/*
URL:
https://atcoder.jp/contests/tenka1-2014-quala/tasks/tenka1_2014_qualA_c
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

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

/*******************************************************************/

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

func init() {
	// bufio.ScanWords <---> bufio.ScanLines
	ReadString = newReadString(os.Stdin, bufio.ScanWords)
	stdout = bufio.NewWriter(os.Stdout)
}

var (
	n, m int
	P    [][]rune
	M    [20][20]bool

	dp [1000000 + 50]int
)

func main() {
	n, m = ReadInt2()
	for i := 0; i < n; i++ {
		row := ReadRuneSlice()
		P = append(P, row)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			M[i][j] = match(P[i], P[j])
		}
	}

	// for S := 0; S < 1<<uint(n); S++ {
	// 	dp[S] = INF_BIT60
	// }

	// 初期化
	for S := 0; S < 1<<uint(n); S++ {
		dp[S] = 1
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if NthBit(S, i) == 1 && NthBit(S, j) == 1 && !M[i][j] {
					dp[S] = INF_BIT60
				}
			}
		}
	}

	// 計算: 4^n
	// for i := 0; i < 1<<uint(n); i++ {
	// 	for j := 0; j < 1<<uint(n); j++ {
	// 		ChMin(&dp[i|j], dp[i]+dp[j])
	// 	}
	// }

	// 計算: 3^n
	for S := 0; S < 1<<uint(n); S++ {
		for T := S; ; T = (T - 1) & S {
			ChMin(&dp[S], dp[T]+dp[S^T])

			if T == 0 {
				break
			}
		}
	}

	fmt.Println(dp[1<<uint(n)-1])
}

func match(S, T []rune) bool {
	for i := 0; i < m; i++ {
		if S[i] == '*' || T[i] == '*' {
			continue
		}

		if S[i] != T[i] {
			return false
		}
	}
	return true
}

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int, target int) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// NthBit returns nth bit value of an argument.
// n starts from 0.
func NthBit(num int, nth int) int {
	return num >> uint(nth) & 1
}

// OnBit returns the integer that has nth ON bit.
// If an argument has nth ON bit, OnBit returns the argument.
func OnBit(num int, nth int) int {
	return num | (1 << uint(nth))
}

// OffBit returns the integer that has nth OFF bit.
// If an argument has nth OFF bit, OffBit returns the argument.
func OffBit(num int, nth int) int {
	return num & ^(1 << uint(nth))
}

// PopCount returns the number of ON bit of an argument.
func PopCount(num int) int {
	res := 0

	for i := 0; i < 70; i++ {
		if ((num >> uint(i)) & 1) == 1 {
			res++
		}
	}

	return res
}

// func main() {
// 	n, m = ReadInt2()
// 	for i := 0; i < n; i++ {
// 		row := ReadRuneSlice()
// 		P = append(P, row)
// 	}

// 	ans := 0
// 	Q := []Node{}

// 	ini := []int{}
// 	for i := 0; i < n; i++ {
// 		ini = append(ini, i)
// 	}
// 	Q = append(Q, Node{G: ini, cur: 0})

// 	for len(Q) > 0 {
// 		pop := Q[0]
// 		Q = Q[1:]

// 		if pop.cur >= m {
// 			ans++
// 			continue
// 		}

// 		memo := [40][]int{}
// 		for _, i := range pop.G {
// 			if P[i][pop.cur] == '*' {
// 				memo[30] = append(memo[30], i)
// 			} else {
// 				memo[P[i][pop.cur]-'a'] = append(memo[P[i][pop.cur]-'a'], i)
// 			}
// 		}

// 		// アスタリスクは適当な文字に寄せる, 1以上が1つも無ければ本当に適当に
// 		ok := false
// 		for i := 0; i < ALPHABET_NUM; i++ {
// 			if len(memo[i]) > 0 {
// 				memo[i] = append(memo[i], memo[30]...)
// 				ok = true
// 				break
// 			}
// 		}
// 		if !ok {
// 			memo[0] = append(memo[0], memo[30]...)
// 		}

// 		// queueにpush
// 		for i := 0; i < ALPHABET_NUM; i++ {
// 			if len(memo[i]) > 0 {
// 				node := Node{G: memo[i], cur: pop.cur + 1}
// 				Q = append(Q, node)
// 			}
// 		}
// 	}

// 	fmt.Println(ans)
// }

// type Node struct {
// 	G   []int
// 	cur int
// }

/*******************************************************************/

/*********** I/O ***********/

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

/*********** Debugging ***********/

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

// Strtoi is a wrapper of strconv.Atoi().
// If strconv.Atoi() returns an error, Strtoi calls panic.
func Strtoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(errors.New("[argument error]: Strtoi only accepts integer string"))
	} else {
		return i
	}
}

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

// PrintfDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintfDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// PrintfBufStdout is function for output strings to buffered os.Stdout.
// You may have to call stdout.Flush() finally.
func PrintfBufStdout(format string, a ...interface{}) {
	fmt.Fprintf(stdout, format, a...)
}