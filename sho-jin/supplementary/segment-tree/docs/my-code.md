# 完成品(@2020-01-18)

自分で書いてみたGoのコードについての説明とか注意点とか。

## 通常（遅延伝搬なし）

基本的にはモノイドの定義を適切に書き換えることに注力すればいいようにしたつもり。

強いて留意点を列挙するなら以下の点。

- インスタンス化してからの初期化は `Set -> Build` の順序で行う。
- `Get` メソッドはセグ木の葉ノードを直接返すようにしている。
  - 後述するように、遅延伝搬ありバージョンの方はそうではない。

```go
type T int // (T, f): Monoid

type SegmentTree struct {
  sz   int              // minimum power of 2
  data []T              // elements in T
  f    func(lv, rv T) T // T <> T -> T
  ti   T                // identity element of Monoid
}

func NewSegmentTree(
  n int, f func(lv, rv T) T, ti T,
) *SegmentTree {
  st := new(SegmentTree)
  st.ti = ti
  st.f = f

  st.sz = 1
  for st.sz < n {
    st.sz *= 2
  }

  st.data = make([]T, 2*st.sz-1)
  for i := 0; i < 2*st.sz-1; i++ {
    st.data[i] = st.ti
  }

  return st
}

func (st *SegmentTree) Set(k int, x T) {
  st.data[k+(st.sz-1)] = x
}

func (st *SegmentTree) Build() {
  for i := st.sz - 2; i >= 0; i-- {
    st.data[i] = st.f(st.data[2*i+1], st.data[2*i+2])
  }
}

func (st *SegmentTree) Update(k int, x T) {
  k += st.sz - 1
  st.data[k] = x

  for k > 0 {
    k = (k - 1) / 2
    st.data[k] = st.f(st.data[2*k+1], st.data[2*k+2])
  }
}

func (st *SegmentTree) Query(a, b int) T {
  return st.query(a, b, 0, 0, st.sz)
}

func (st *SegmentTree) query(a, b, k, l, r int) T {
  if r <= a || b <= l {
    return st.ti
  }

  if a <= l && r <= b {
    return st.data[k]
  }

  lv := st.query(a, b, 2*k+1, l, (l+r)/2)
  rv := st.query(a, b, 2*k+2, (l+r)/2, r)
  return st.f(lv, rv)
}

func (st *SegmentTree) Get(k int) T {
  return st.data[k+(st.sz-1)]
}
```

## 遅延伝搬あり

`Assumption: T == E` みたいなことを書いてしまったが、
別にそうとも限らない（割と早い段階で `T != E` な問題に出会ってしまった）。
が、まずはシンプルなケースに慣れたいため、原則 `T == E` と思っておく。

こちらもモノイド、作用素モノイド、
そしてそれらに関する関数オブジェクトを適切に決定することに注力しさえすればいいように書いたつもり。

`Get` メソッドは注意が必要！

こちらは遅延伝搬を考慮しなければならない関係で、
通常バージョンのように直接セグメントツリーの葉ノードを参照すると真の値が得られない。
よって、1要素のみの区間を取得するクエリをラッピングしている。

ここを忘れるとデバッグ時にハマったりするので注意
（生現在の葉ノードの要素が知りたいのに一見おかしな内容が返ってくるように見える）。

```go
// Assumption: T == E
type T int // (T, f): Monoid
type E int // (E, h): Operator Monoid

type LazySegmentTree struct {
  sz   int
  data []T
  lazy []E
  f    func(lv, rv T) T        // T <> T -> T
  g    func(to T, from E) T    // T <> E -> T (assignment operator)
  h    func(to, from E) E      // E <> E -> E (assignment operator)
  p    func(e E, length int) E // E <> N -> E
  ti   T
  ei   E
}

func NewLazySegmentTree(
  n int,
  f func(lv, rv T) T, g func(to T, from E) T,
  h func(to, from E) E, p func(e E, length int) E,
  ti T, ei E,
) *LazySegmentTree {
  lst := new(LazySegmentTree)
  lst.f, lst.g, lst.h, lst.p = f, g, h, p
  lst.ti, lst.ei = ti, ei

  lst.sz = 1
  for lst.sz < n {
    lst.sz *= 2
  }

  lst.data = make([]T, 2*lst.sz-1)
  lst.lazy = make([]E, 2*lst.sz-1)
  for i := 0; i < 2*lst.sz-1; i++ {
    lst.data[i] = lst.ti
    lst.lazy[i] = lst.ei
  }

  return lst
}

func (lst *LazySegmentTree) Set(k int, x T) {
  lst.data[k+(lst.sz-1)] = x
}

func (lst *LazySegmentTree) Build() {
  for i := lst.sz - 2; i >= 0; i-- {
    lst.data[i] = lst.f(lst.data[2*i+1], lst.data[2*i+2])
  }
}

func (lst *LazySegmentTree) propagate(k, length int) {
  if lst.lazy[k] != lst.ei {
    if k < lst.sz-1 {
      lst.lazy[2*k+1] = lst.h(lst.lazy[2*k+1], lst.lazy[k])
      lst.lazy[2*k+2] = lst.h(lst.lazy[2*k+2], lst.lazy[k])
    }
    lst.data[k] = lst.g(lst.data[k], lst.p(lst.lazy[k], length))
    lst.lazy[k] = lst.ei
  }
}

func (lst *LazySegmentTree) Update(a, b int, x E) T {
  return lst.update(a, b, x, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) update(a, b int, x E, k, l, r int) T {
  lst.propagate(k, r-l)

  if r <= a || b <= l {
    return lst.data[k]
  }

  if a <= l && r <= b {
    lst.lazy[k] = lst.h(lst.lazy[k], x)
    lst.propagate(k, r-l)
    return lst.data[k]
  }

  lv := lst.update(a, b, x, 2*k+1, l, (l+r)/2)
  rv := lst.update(a, b, x, 2*k+2, (l+r)/2, r)
  lst.data[k] = lst.f(lv, rv)
  return lst.data[k]
}

func (lst *LazySegmentTree) Query(a, b int) T {
  return lst.query(a, b, 0, 0, lst.sz)
}

func (lst *LazySegmentTree) query(a, b, k, l, r int) T {
  lst.propagate(k, r-l)

  if r <= a || b <= l {
    return lst.ti
  }

  if a <= l && r <= b {
    return lst.data[k]
  }

  lv := lst.query(a, b, 2*k+1, l, (l+r)/2)
  rv := lst.query(a, b, 2*k+2, (l+r)/2, r)
  return lst.f(lv, rv)
}

func (lst *LazySegmentTree) Get(k int) T {
  return lst.Query(k, k+1)
}
```
