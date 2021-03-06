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

/*********** I/O ***********/

var (
	// ReadString returns a WORD string.
	ReadString func() string
	stdout     *bufio.Writer
)

func init() {
	ReadString = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}

func newReadString(ior io.Reader) func() string {
	r := bufio.NewScanner(ior)
	// r.Buffer(make([]byte, 1024), int(1e+11)) // for AtCoder
	r.Buffer(make([]byte, 1024), int(1e+9)) // for Codeforces
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.
	r.Split(bufio.ScanWords)

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

// PrintDebug is wrapper of fmt.Fprintf(os.Stderr, format, a...)
func PrintDebug(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
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

/*******************************************************************/

const MOD = 1000000000 + 7
const ALPHABET_NUM = 26
const INF_INT64 = math.MaxInt64
const INF_BIT60 = 1 << 60
const INF_INT32 = math.MaxInt32
const INF_BIT30 = 1 << 30

var q, k int
var queries [][]int

func main() {
	q, k = ReadInt2()

	for i := 0; i < q; i++ {
		c := ReadInt()
		if c == 1 {
			v := ReadInt()
			queries = append(queries, []int{c, v})
		} else {
			queries = append(queries, []int{c})
		}
	}

	elems := []int{}
	for i := 0; i < q; i++ {
		if len(queries[i]) == 1 {
			continue
		}

		elems = append(elems, queries[i][1])
	}

	_, otop, ptoo := ZaAtsu1Dim(elems, 1)

	bit := NewBITSet(200000 + 5)
	for i := 0; i < q; i++ {
		que := queries[i]
		if que[0] == 1 {
			v := que[1]
			vv := otop[v]
			bit.Insert(vv, 1)
		} else {
			if bit.Count(bit.n) < k {
				fmt.Println(-1)
			} else {
				v := bit.Kth(k)
				vv := ptoo[v]
				fmt.Println(vv)
				bit.Delete(v, 1)
			}
		}
	}
}

// Public methods
// bit := NewBITSet(200000 + 5)
// c := bit.Count(b.n)
// bit.Insert(val, 1)
// bit.Delete(val, 1)
// ans := bit.Kth(k)

type BITSet struct {
	bit     []int
	n       int
	minPow2 int
}

// n(>=1) is maximum integer for the set.
func NewBITSet(n int) *BITSet {
	newBit := new(BITSet)

	newBit.bit = make([]int, n+1)
	newBit.n = n

	newBit.minPow2 = 1
	for {
		if (newBit.minPow2 << 1) > n {
			break
		}
		newBit.minPow2 <<= 1
	}

	return newBit
}

// Count returns number of elements less or equal than e in the set.
// b.Count(b.n) returns total number of elements in the set.
// O(logN)
func (b *BITSet) Count(e int) int {
	s := 0

	for e > 0 {
		s += b.bit[e]
		e -= e & (-e)
	}

	return s
}

// Insert e(1<=e<=n) for num(>= 1) times.
func (b *BITSet) Insert(e, num int) {
	for e <= b.n {
		b.bit[e] += num
		e += e & (-e)
	}
}

// Delete e(1<=e<=n) for num(>= 1) times.
func (b *BITSet) Delete(e, num int) {
	num *= -1
	for e <= b.n {
		b.bit[e] += num
		e += e & (-e)
	}
}

// Kth returns kth element in the set
func (b *BITSet) Kth(kth int) int {
	if kth <= 0 {
		return 0
	}

	x := 0
	for k := b.minPow2; k > 0; k /= 2 {
		if x+k <= b.n && b.bit[x+k] < kth {
			kth -= b.bit[x+k]
			x += k
		}
	}

	return x + 1
}

// ---

// ZaAtsu1Dim returns 3 values.
// pressed: pressed slice of the original slice
// orgToPress: map for translating original value to pressed value
// pressToOrg: reverse resolution of orgToPress
// O(nlogn)
func ZaAtsu1Dim(org []int, initVal int) (pressed []int, orgToPress, pressToOrg map[int]int) {
	pressed = make([]int, len(org))
	copy(pressed, org)
	sort.Sort(sort.IntSlice(pressed))

	orgToPress = make(map[int]int)
	for i := 0; i < len(org); i++ {
		if i == 0 {
			orgToPress[pressed[0]] = initVal
			continue
		}

		if pressed[i-1] != pressed[i] {
			initVal++
			orgToPress[pressed[i]] = initVal
		}
	}

	for i := 0; i < len(org); i++ {
		pressed[i] = orgToPress[org[i]]
	}

	pressToOrg = make(map[int]int)
	for k, v := range orgToPress {
		pressToOrg[v] = k
	}

	return
}
