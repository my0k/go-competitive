/*
URL:
https://atcoder.jp/contests/abc178/tasks/abc178_c
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
)

func main() {
	defer stdout.Flush()

	n = readi()
	if n == 1 {
		fmt.Println(0)
		return
	}

	ten := modpow(10, n, MOD)
	eight := modpow(8, n, MOD)
	nine := modpow(9, n, MOD)

	all := modi(ten-eight, MOD)

	tmp := modi(nine-eight, MOD)
	tmp *= 2
	tmp %= MOD

	ans := modi(all-tmp, MOD)
	fmt.Println(ans)
}

// ModInv returns $a^{-1} mod m$ by Fermat's little theorem.
// O(1), but C is nearly equal to 30 (when m is 1000000000+7).
func ModInv(a, m int) int {
	return modpow(a, m-2, m)
}

func modpow(a, e, m int) int {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := modpow(a, halfE, m)
		return half * half % m
	}

	return a * modpow(a, e-1, m) % m
}

// cf := NewCombFactorial(2000000) // maxNum == "maximum n" * 2 (for H(n,r))
// res := cf.C(n, r) 	// 組み合わせ
// res := cf.H(n, r) 	// 重複組合せ
// res := cf.P(n, r) 	// 順列

type CombFactorial struct {
	factorial, invFactorial []int
	maxNum                  int
}

func NewCombFactorial(maxNum int) *CombFactorial {
	cf := new(CombFactorial)
	cf.maxNum = maxNum
	cf.factorial = make([]int, maxNum+50)
	cf.invFactorial = make([]int, maxNum+50)
	cf.initCF()

	return cf
}
func (c *CombFactorial) modInv(a int) int {
	return c.modpow(a, MOD-2)
}
func (c *CombFactorial) modpow(a, e int) int {
	if e == 0 {
		return 1
	}

	if e%2 == 0 {
		halfE := e / 2
		half := c.modpow(a, halfE)
		return half * half % MOD
	}

	return a * c.modpow(a, e-1) % MOD
}
func (c *CombFactorial) initCF() {
	var i int

	for i = 0; i <= c.maxNum; i++ {
		if i == 0 {
			c.factorial[i] = 1
			c.invFactorial[i] = c.modInv(c.factorial[i])
			continue
		}

		num := i * c.factorial[i-1]
		num %= MOD
		c.factorial[i] = num
		c.invFactorial[i] = c.modInv(c.factorial[i])
	}
}
func (c *CombFactorial) C(n, r int) int {
	var res int

	res = 1
	res *= c.factorial[n]
	res %= MOD
	res *= c.invFactorial[r]
	res %= MOD
	res *= c.invFactorial[n-r]
	res %= MOD

	return res
}
func (c *CombFactorial) P(n, r int) int {
	var res int

	res = 1
	res *= c.factorial[n]
	res %= MOD
	res *= c.invFactorial[n-r]
	res %= MOD

	return res
}
func (c *CombFactorial) H(n, r int) int {
	return c.C(n-1+r, r)
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
