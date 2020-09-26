# グラフ

---

## 無向グラフの用語

隣接: 2つの頂点間に辺があるとき、その2つの頂点を隣接している、という。

パス: 隣接している頂点の列

閉路: 始点と終点が同じようなパス

連結グラフ: 任意の2頂点間にパスが存在するようなグラフ

頂点の次数: 頂点につながっている辺の数

木: 連結グラフで、閉路を持たないようなもの

森: 連結でないグラフで、閉路を持たないようなもの

※木は辺の数がちょうど(頂点数-1)。

※逆に、辺の数が(頂点数-1)に等しい連結なグラフは木となる。

### 根付き木について

> 根（root）と呼ばれる頂点を1つ選ぶと、根を一番上に、根から遠い頂点ほど下になるように木を配置することができます。
> このような木を根付き木といいます。
> **根付き木でない普通の木も、適当に頂点を選んで根付き木にすることで扱いやすくなる場合もあります。**
> 根付き木を家系図のように見立てると、頂点の間に親子関係を与える事ができます。
> **これは、辺に向きを与えたと考えることもできます。**

---

## 有向グラフの用語

ノードvの出次数: vから出ていく辺の集合の大きさ（＝辺の数）

ノードvの入次数: vへと入ってくる辺の集合の大きさ（＝辺の数）

DAG: 有向グラフで閉路を持たないグラフ（Directed Acyclic Graph）

**※例えば、整数を頂点とし、「nがmを割り切るときに、nからmに辺を引いたグラフ」を考えると、これはDAGになる。**
（nとmの大小関係を考えると閉路はできなさそう。）

トポロジカル順序: DAGの各頂点への番号付で、頂点 `v[i]` から `v[j]` に向かって辺があるときに、 `i<j` が成り立つような番号付のこと！

**※DAGの問題をDPで解くための準備方法！**

トポロジカルソート: トポロジカル順序を求めること。

---

## グラフの表現

### 隣接行列（有益そうなポイントだけ）

- `i -> j` の辺がない場合は `g[i][j] = INF` とする
  - よく `-1` としてしまうので、この癖は直しておきたい。
- 多重辺や自己ループも、重み付きでないなら `g[i][j]` に辺の個数を入れれば良い。
  - 重み付きの場合は無理。

---

## 二部グラフ判定

グラフの彩色問題: 隣接頂点同士を違う色になるようにする問題。

グラフの彩色数: 彩色問題で、必要な色の数の最小値。

二部グラフ: 彩色数が2であるグラフ。

---

## グラフのDFSでできること

1. 連結判定
2. 木判定
3. トポロジカルソート

### DFSの計算量

**すべての頂点と辺を一度ずつ見るため、 `O(|V| + |E|)` ！**

---

## 最短路問題

### 単一始点最短路問題1: ベルマンフォード法

1. 各点のコストを、始点は0、それ以外はINFに初期化する。
2. 1回のループですべての辺に対して、「選択した辺の、一方の点から他方の点をたどる場合のコストをそれぞれ計算する」ということを繰り返す。
3. 負の閉路がない場合は少なくとも `|v|-1` 回で収束する（最も遠い点でも最大で `|v|-1` 個の辺しか通らないため）。

#### ポイント

- 計算量はループの上限から分かる通り `O(|v|×|e|)`
- `|v|` 回目のループでも更新が発生する場合は負の閉路が存在している
- **すべての `i` について `d[i] = 0` と初期化を行えば、すべての不平路の検出を行うこともできる**

#### 実装ポイント

隣接リストではなく、始点と終点とコストの情報をもった辺の構造体をリストで管理する。

#### 負の閉路の検出方法について

> すべての `i` について `dist[i] = 0` と初期化を行えば、すべての負の閉路の検出を行う事もできます。

これの原理がよくわからなかったが、単一始点最短路問題では始点の距離を0で初期化することから、
↑の方法は最初すべての点を始点とみなしているのかもしれない（？）

### 単一始点最短路問題2: ダイクストラ法

ベルマンフォード法の欠点を修正。

1. 最短距離が確定した頂点と隣接している頂点を更新する
2. 1で使った「最短距離が確定した頂点」はもう使わない

計算量は `O(|e|log|v|)` （ちゃんとノードを取り出したときに最短が確定している場合は捨てないとこの計算量にならないことに注意）

（捨てない場合は `O(|e|log|e|)` ？）

#### 実装ポイント

隣接リストで管理すれば良い。

### 全点対最短路問題: ワーシャルフロイド法

#### ポイント

- 実装が簡単なので `O(n^3)` という計算量に耐えられるのであれば積極的に使ったほうがいい。
- **負の辺があっても動作する！**
  - 負の閉路があるかないかは `d[i][i]` が負になる頂点 `i`があるかどうかで判別できる。

#### 考え方: DP

wikipediaなどの解説も詳しかったが、「`k` について1周で十分なのはなぜか？」みたいな疑問が生じていて、なんとなく腑に落ちていなかった。

**DPだと言われると納得しやすかったので、書いて理解を深める。**

3次元のDPテーブルを考える。

```golang
dp[k+1][i][j] := 頂点0~kとi,jのみを使う場合の、iからjへの最短路
dp[0][i][j] = e[i][j] or INFで初期化
```

と定義する。

頂点 `0~k, i, j` のみを使うとき、 `i->j` への最短路は、

1. 頂点 `k` をちょうど一度通る場合
2. 頂点 `k` をまったく通らない場合

の2つの排反なパターンに分かれる。

それぞれ、

1. `dp[k+1][i][j] = dp[k-1][i][j]`: `k` までの最短路が採用される
2. `dp[k+1][i][j] = dp[k][i][k] + dp[k][k][j]`: `i->k` の最短路と `k->j` の最短路に `k` が分解する

**このようなDPの配列再利用によって、2次元テーブルで済ませたものが一般的なもの。**

※DPだと言われると、それまでの部分問題の正しい答えがわかっているのだから、正当性について納得しやすい。

#### 実装ポイント

計算量の都合上、単純に隣接行列的に管理してやれば良い。
（隣接リスト等は不要。）

※DPの配列再利用をしてしまうと、問題に寄っては対応が難しくなる気がするので、3次元のDPテーブルを持っておくのも悪くない気がする（AtCoderでしか使えないかもしれないが）。

---

## 最小全域木

### 定義

**無向グラフ** が与えられたときに、その部分グラフで任意の2頂点を連結にするような木を全域木（Spanning Tree）という。

辺にコストがある場合に、使われる辺のコストの和を最小にする全域木を最小全域木（MST: Minimum Spanning Tree）という。

※全域木が存在することと、連結であることは同値。

### 最小全域疑問題1: プリム法

**ダイクストラ法に似ている！** : ある頂点から始めて少しずつ辺を追加していく方法。

※証明は貪欲法っぽい。

#### 計算量

プライオリティキューを使うと、計算量はダイクストラ法と同じ `O(|E|log|V|)` になる。

#### 実装ポイント

ダイクストラと非常に似ているので、実装も自然と似てくる。

- 隣接リストでグラフを持つ。
- 優先度付きキューを使う。

### 最小全域木問題2: クラスカル法

**閉路（二重辺なども含む）ができないように** 辺をコストの小さい順に追加していく方法。

※プリム法は（ダイクストラ法のように）頂点中心に考え、クラスカル法は辺中心に考える、とおぼえておくといいかもしれない。

#### 閉路判定はどうするか？

Union Find木を使う！

頂点uと頂点vを結ぶ辺eを追加しようと考えるとき、

1. 追加前にuとvが同じ連結成分に属していなければ、eを追加しても閉路はできない。
2. 逆にuとvが同じ連結成分に属していれば、必ず閉路ができる。

#### 計算量

辺をソートする部分がネックになるため `O(|E|log|V|)` となる。

（`O(|E|log|E|)` では？と思ったが、 `|E|` の数は最大でも `|V|^2` 程度かつ `log(|V|^2) = 2log|V|` であることから上のようになる）

#### 実装ポイント

- 始点と終点とコストを持った辺の構造体のリストでグラフを管理する（ベルマンフォード法のような感じ）。
- さらにその辺の構造体リストをコストで昇順ソートする。
- Union-Find木を使う。

Union-Find木の実装方法次第だが、コアとなるロジック部分はプリム法よりもスッキリする気がする。

※クラスカル法では、無向グラフでも逆向きの辺はリストに加えなくても大丈夫？
（UF木で連結判定を行うだけなため。）

---

## グラフアルゴリズムのverify問集

### ベルマンフォード法

- [Score Attack](https://atcoder.jp/contests/abc061/tasks/abc061_d)

### ダイクストラ法


### ワーシャルフロイド法

- [joisino's travel](https://atcoder.jp/contests/abc073/tasks/abc073_d)

### プリム法、クラスカル法

- [AOJ Minimum Spanning Tree](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=GRL_2_A&lang=jp)
  - 入力が隣接リスト形式
- [AOJ Graph2 Minimum Spanning Tree](http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_12_A&lang=jp)
  - 入力が隣接行列
