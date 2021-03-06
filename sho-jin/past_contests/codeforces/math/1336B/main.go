/*
URL:
https://codeforces.com/problemset/problem/1336/B
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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
	MOD          = 1000000000 + 7
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
	t       int
	r, g, b int
	R, G, B []int64
)

func main() {
	t = ReadInt()
	for i := 0; i < t; i++ {
		r, g, b = ReadInt3()
		R, G, B = ReadInt64Slice(r), ReadInt64Slice(g), ReadInt64Slice(b)
		R = append(R, -(1 << 30), 1<<31)
		G = append(G, -(1 << 30), 1<<31)
		B = append(B, -(1 << 30), 1<<31)

		solve()
	}
}

func solve() {
	sort.SliceStable(R, func(i, j int) bool { return R[i] < R[j] })
	sort.SliceStable(G, func(i, j int) bool { return G[i] < G[j] })
	sort.SliceStable(B, func(i, j int) bool { return B[i] < B[j] })
	// PrintfDebug("R: %v\n", R)
	// PrintfDebug("B: %v\n", B)
	// PrintfDebug("G: %v\n", G)

	ans := int64(INF_INT64)

	// for i := 1; i <= g; i++ {
	// 	gw := G[i]
	// 	uridx := BinarySearch(r+1, -1, func(mid int) bool { return R[mid] >= gw })
	// 	lridx := BinarySearch(0, r+2, func(mid int) bool { return R[mid] < gw })
	// 	RT := []int64{R[uridx], R[lridx]}
	// 	ubidx := BinarySearch(b+1, -1, func(mid int) bool { return B[mid] >= gw })
	// 	lbidx := BinarySearch(0, b+2, func(mid int) bool { return B[mid] < gw })
	// 	BT := []int64{B[ubidx], B[lbidx]}

	// 	for j := 0; j < len(RT); j++ {
	// 		for k := 0; k < len(BT); k++ {
	// 			rw, bw := RT[j], BT[k]
	// 			val := (gw-rw)*(gw-rw) + (gw-bw)*(gw-bw) + (rw-bw)*(rw-bw)
	// 			ChMin(&ans, val)
	// 		}
	// 	}
	// }
	// for i := 1; i <= r; i++ {
	// 	rw := R[i]
	// 	ugidx := BinarySearch(g+1, -1, func(mid int) bool { return G[mid] >= rw })
	// 	lgidx := BinarySearch(0, g+2, func(mid int) bool { return G[mid] < rw })
	// 	GT := []int64{G[ugidx], G[lgidx]}
	// 	ubidx := BinarySearch(b+1, -1, func(mid int) bool { return B[mid] >= rw })
	// 	lbidx := BinarySearch(0, b+2, func(mid int) bool { return B[mid] < rw })
	// 	BT := []int64{B[ubidx], B[lbidx]}

	// 	for j := 0; j < len(GT); j++ {
	// 		for k := 0; k < len(BT); k++ {
	// 			gw, bw := GT[j], BT[k]
	// 			val := (gw-rw)*(gw-rw) + (gw-bw)*(gw-bw) + (rw-bw)*(rw-bw)
	// 			ChMin(&ans, val)
	// 		}
	// 	}
	// }
	// for i := 1; i <= b; i++ {
	// 	bw := B[i]
	// 	uridx := BinarySearch(r+1, -1, func(mid int) bool { return R[mid] >= bw })
	// 	lridx := BinarySearch(0, r+2, func(mid int) bool { return R[mid] < bw })
	// 	RT := []int64{R[uridx], R[lridx]}
	// 	ugidx := BinarySearch(g+1, -1, func(mid int) bool { return G[mid] >= bw })
	// 	lgidx := BinarySearch(0, g+2, func(mid int) bool { return G[mid] < bw })
	// 	GT := []int64{G[ugidx], G[lgidx]}

	// 	for j := 0; j < len(RT); j++ {
	// 		for k := 0; k < len(GT); k++ {
	// 			rw, gw := RT[j], GT[k]
	// 			val := (gw-rw)*(gw-rw) + (gw-bw)*(gw-bw) + (rw-bw)*(rw-bw)
	// 			ChMin(&ans, val)
	// 		}
	// 	}
	// }
	ChMin(&ans, sub(r, g, b, R, G, B))
	ChMin(&ans, sub(r, b, g, R, B, G))
	ChMin(&ans, sub(g, r, b, G, R, B))

	fmt.Println(ans)
}

func sub(r, g, b int, R, G, B []int64) int64 {
	ans := int64(INF_INT64)

	for i := 1; i <= g; i++ {
		gw := G[i]
		uridx := BinarySearch(r+1, -1, func(mid int) bool { return R[mid] >= gw })
		lridx := BinarySearch(0, r+2, func(mid int) bool { return R[mid] <= gw })
		RT := []int64{R[uridx], R[lridx]}
		ubidx := BinarySearch(b+1, -1, func(mid int) bool { return B[mid] >= gw })
		lbidx := BinarySearch(0, b+2, func(mid int) bool { return B[mid] <= gw })
		BT := []int64{B[ubidx], B[lbidx]}

		for j := 0; j < len(RT); j++ {
			for k := 0; k < len(BT); k++ {
				rw, bw := RT[j], BT[k]

				if IsProductOverflow(gw-rw, gw-rw, INF_INT64) {
					continue
				}
				if IsProductOverflow(gw-bw, gw-bw, INF_INT64) {
					continue
				}
				if IsProductOverflow(rw-bw, rw-bw, INF_INT64) {
					continue
				}

				val := (gw-rw)*(gw-rw) + (gw-bw)*(gw-bw) + (rw-bw)*(rw-bw)
				ChMin(&ans, val)
			}
		}
	}

	return ans
}

func IsProductOverflow(i, j, m int64) bool {
	return float64(i)*float64(j) > float64(m)
}

func IsAddOverflow(i, j, m int64) bool {
	return float64(i)+float64(j) > float64(m)
}

// ChMin accepts a pointer of integer and a target value.
// If target value is SMALLER than the first argument,
//	then the first argument will be updated by the second argument.
func ChMin(updatedValue *int64, target int64) bool {
	if *updatedValue > target {
		*updatedValue = target
		return true
	}
	return false
}

// PowInt is integer version of math.Pow
// PowInt calculate a power by Binary Power (二分累乗法(O(log e))).
func PowInt(a, e int64) int64 {
	if a < 0 || e < 0 {
		panic(errors.New("[argument error]: PowInt does not accept negative integers"))
	}

	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := PowInt(a, halfE)
		return half * half
	}

	return a * PowInt(a, e-1)
}

func BinarySearch(initOK, initNG int, isOK func(mid int) bool) (ok int) {
	ng := initNG
	ok = initOK
	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if isOK(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

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
