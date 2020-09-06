# AGC047 感想

1年ぶりに参加したが楽しかった。冷えたけど。

## A問題

最初に数え上げ方法を間違えたのはすぐに修正できたが、
そもそも入力を受け取る部分で誤差を考慮しきれていなかった。

**浮動小数点を受け取るときは常に誤差を意識する**
**特に、浮動小数点を処理して整数で答えるようなものは、必ず何かしらの特殊な処理が要求される、と思ったほうが良い。**

## B問題

ローリングハッシュとtrie木なるものを使う方法2種類あるっぽい。

どちらもちゃんと復習したい。

### 解き方

公式解説がないのと、若干他の人のブログの説明も不明瞭（に感じられる）ので、書き残しておく。

#### 観察

すぐわかることとして、以下の2点がある。

- 長い方の文字列に対してFirst Second処理を行うことで、最初からより短い他方の文字列を得ることができる。
- 長い方から作れる文字列は「長さ0以上の接尾辞＋接尾辞以外のその文字列中に存在する文字種1文字」というものに限られる。

1つ目の特徴から、平面走査法やzero-sum-rangesよろしく、この問題も「時間を固定」して考えることが重要。
より具体的には、「短い文字列から処理していき、ある文字列に着目しながら条件を満たすものを加算する」という方針をベースに考える。

また、2つ目の特徴において、接尾辞については文字列全体を舐める必要があるものの、
文字種については26文字だけ（しかも、さらに存在するもののみ）を考えれば良い、ということに注意。

以上の観察を元に、Trieとローリングハッシュを用いた解法がそれぞれ考えられる。
どちらも重要なので、やってみるのが良い。

※文字列系の問題に置いて「文字種26は定数」という認識を持っておくことは重要かもしれない。

#### 共通の前処理

これまでの考察から分かる通り、すべての文字列を反転させて、文字列長に対して昇順にソートしておくのが良い。

※辞書式順序は特に木にしなくて良い。

#### Trieを使った解法

すでにスキャンした文字列がTrieに登録されている、と仮定して考える。

今注目中の単語でTrieを検索するにあたって、現在探索中のノードごとに集約処理を考える。

> 「次のノードを見たときに、何らかの文字で受理状態になっている（つまり、ある単語が登録されている）」、
> かつ、「検索文字列の以降にその文字が存在する」という条件を満たすときに、集約値をインクリメントする。

これは、注目中の単語からすでに見てきた短い文字列のうち、生成可能なものの個数を数え上げる操作にほかならない。

この計算回数をざっくり見積もると、だいたい「注目中の文字列の長さ×26」ぐらいになる。
制約で、すべての文字列の長さの和は `10^6` 以下とあるので、これは十分に高速。

※この集約処理は新たなメソッドとしてアドホックに対応したが、このあたりも一般化する方法も検討したい。

#### ローリングハッシュを使った解法

上述のTrieで行った方法が高速にシミュレートできれば良い。

「接尾辞＋それとは分断されたある適当な文字」のような文字列に対するローリングハッシュが計算したくなるが、
これをいちいち高速に計算するのは難しい（多分できないと思う、できても非常にめんどくさいしバグに注意が必要そう）。

そこで、接尾辞に該当するハッシュ値と、末尾の1文字を別個に扱い、以下のようなミュータブルな構造体を考える。

```go
type V struct {
  hash uint64
  lc rune
}
```

これを普通の `map` で管理し、注目中の文字列に関しても、各接尾辞ごとに都度求めてやれば、同様の処理を実現できる。
