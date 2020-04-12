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

/*
ASCII code

ASCII   10進数  ASCII   10進数  ASCII   10進数
!       33      "       34      #       35
$       36      %       37      &       38
'       39      (       40      )       41
*       42      +       43      ,       44
-       45      .       46      /       47
0       48      1       49      2       50
3       51      4       52      5       53
6       54      7       55      8       56
9       57      :       58      ;       59
<       60      =       61      >       62
?       63      @       64      A       65
B       66      C       67      D       68
E       69      F       70      G       71
H       72      I       73      J       74
K       75      L       76      M       77
N       78      O       79      P       80
Q       81      R       82      S       83
T       84      U       85      V       86
W       87      X       88      Y       89
Z       90      [       91      \       92
]       93      ^       94      _       95
`       96      a       97      b       98
c       99      d       100     e       101
f       102     g       103     h       104
i       105     j       106     k       107
l       108     m       109     n       110
o       111     p       112     q       113
r       114     s       115     t       116
u       117     v       118     w       119
x       120     y       121     z       122
{       123     |       124     }       125
~       126             127
*/

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

var (
	S     []rune
	num   int
	combo int

	isIgnored [1000000 + 5]bool
	memoCombo [1000000 + 5]int
)

func main() {
	S = ReadRuneSlice()

	// temp := make(EventPQ, 0)
	// pq := &temp
	// heap.Init(pq)

	L := make(EventList, 0)

	for i := 0; i < len(S); i++ {
		if S[i] == '-' {
			continue
		}

		time := float64(i)
		if S[i] == 'N' {
			// throw := &Event{pri: time + 0.5, kind: S[i], diff: -1, eid: i}
			throw := &Event{key: time, kind: S[i], diff: -1, eid: i}
			damage := &Event{key: time + 1.5, kind: S[i], diff: 0, eid: i, damage: 1}
			back := &Event{key: time + 6.5, kind: S[i], diff: 1, eid: i}

			// heap.Push(pq, throw)
			// heap.Push(pq, damage)
			// heap.Push(pq, back)
			L = append(L, throw)
			L = append(L, damage)
			L = append(L, back)
		} else {
			// throw := &Event{pri: time + 2.5, kind: S[i], diff: -3, eid: i}
			throw := &Event{key: time, kind: S[i], diff: -3, eid: i}
			damage := &Event{key: time + 3.5, kind: S[i], diff: 0, eid: i, damage: 5}
			back := &Event{key: time + 8.5, kind: S[i], diff: 3, eid: i}

			// heap.Push(pq, throw)
			// heap.Push(pq, damage)
			// heap.Push(pq, back)
			L = append(L, throw)
			L = append(L, damage)
			L = append(L, back)
		}
	}

	sort.Stable(byKey{L})

	num = 5
	combo = 0
	ans := 0
	t := 0.0
	// for pq.Len() > 0 {
	for i := 0; i < len(L); i++ {
		// pop := heap.Pop(pq).(*Event)
		pop := L[i]

		if pop.diff < 0 {
			// 投げイベント
			if num+pop.diff >= 0 && pop.key >= t {
				// 投げる
				num += pop.diff
				isIgnored[pop.eid] = true
				memoCombo[pop.eid] = combo

				if pop.kind == 'N' {
					t = pop.key + 0.5
				} else {
					t = pop.key + 2.5
				}
			}
		} else if isIgnored[pop.eid] {
			// 投げられていたのならば処理する
			if pop.diff == 0 {
				// ダメージ計算
				c := memoCombo[pop.eid]
				c /= 10
				c += 10
				ans += pop.damage * c

				// コンボは一律で1インクリメント
				combo++
			} else {
				// 数を戻す
				num += pop.diff
			}
		}
	}

	fmt.Println(ans)
}

type Event struct {
	key               float64
	eid, diff, damage int
	kind              rune
}
type EventList []*Event
type byKey struct {
	EventList
}

func (l EventList) Len() int {
	return len(l)
}
func (l EventList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byKey) Less(i, j int) bool {
	return l.EventList[i].key < l.EventList[j].key
}

// how to use
// L := make(EventList, 0, 200000+5)
// L = append(L, &Event{key: intValue})
// sort.Stable(byKey{ L })                // Stable ASC
// sort.Stable(sort.Reverse(byKey{ L }))  // Stable DESC

// type Event struct {
// 	pri    float64
// 	eid    int
// 	kind   rune
// 	diff   int
// 	damage int
// }
// type EventPQ []*Event

// func (pq EventPQ) Len() int           { return len(pq) }
// func (pq EventPQ) Less(i, j int) bool { return pq[i].pri < pq[j].pri } // <: ASC, >: DESC
// func (pq EventPQ) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }
// func (pq *EventPQ) Push(x interface{}) {
// 	item := x.(*Event)
// 	*pq = append(*pq, item)
// }
// func (pq *EventPQ) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	*pq = old[0 : n-1]
// 	return item
// }

// how to use
// temp := make(EventPQ, 0, 100000+1)
// pq := &temp
// heap.Init(pq)
// heap.Push(pq, &Event{pri: intValue})
// popped := heap.Pop(pq).(*Event)

/*
- まずは全探索を検討しましょう
- MODは最後にとりましたか？
- 負のMODはちゃんと関数を使って処理していますか？
- ループを抜けた後も処理が必要じゃありませんか？
- 和・積・あまりを求められたらint64が必要ではありませんか？
- いきなりオーバーフローはしていませんか？
- MOD取る系はint64必須ですよ？
- 後ろ・逆・ゴールから考えましたか？
- 3者のうち真ん中に着目しましたか？
*/

/*******************************************************************/