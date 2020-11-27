# 確率DP

Last Change: 2020-11-26 23:26:02.

期待値DPが一番の勉強目的だったが、確率DPや期待値の線形性といったトピックも一緒に学ぶ。

## 参考

- [確率DPを極めよう](https://compro.tsutaj.com//archive/180220_probability_dp.pdf)
- [けんちょん氏のABC078-C HSIの解説](https://drken1215.hatenablog.com/entry/2019/03/23/175300)
- [「〜がすべて終わるまでの試行回数の期待値」を求める一般的なフレームワーク](https://drken1215.hatenablog.com/entry/2019/03/23/214500)
- [「期待値の線形性」についての解説と、それを用いる問題のまとめ](https://qiita.com/drken/items/3fe1613c44de1f3bfbe1)
  - かなり易しいところから始めてくれるので、迷ったら読み返すと良い。

## スタータパック

[ZRK+94さんのツイートより](https://twitter.com/Zen_Re_Kkyo/status/1135152582194651136)より

- [AtCoder社の給料](https://atcoder.jp/contests/abc003/tasks/abc003_1)
- [HSI](https://atcoder.jp/contests/abc078/tasks/arc085_a)
- [Theme Color](https://atcoder.jp/contests/code-festival-2018-final/tasks/code_festival_2018_final_b)
- [デフレゲーム](https://atcoder.jp/contests/tkppc3/tasks/tkppc3_e)
- [Ordinary Beauty](https://atcoder.jp/contests/soundhound2018-summer-qual/tasks/soundhound2018_summer_qual_c)
- [Removing Blocks](https://atcoder.jp/contests/agc028/tasks/agc028_b)
- [Modulo Operations](https://atcoder.jp/contests/exawizards2019/tasks/exawizards2019_d)

## コツ？

- 状態遷移図を書いてみることは意味がありそう。

## 必要な知識

知識を頭に常においておくのが良いと思う。  
できれば証明も併せて頭に入れておきたい（証明における考え方自体も応用範囲が広いように思えるので）。

- 確率 `p` で成功する試行について、成功がでるまでの期待値は `1/p`
  - 超典型
  - 「 `m` 回成功が出るまでの回数」というのを考える場合もある。これについてはシンプルに `m*(1/p)` でよい。
- コンプガチャ問題
  - この考え方がそのまま応用できるものも少なくないように見える。
    - https://qiita.com/drken/items/3fe1613c44de1f3bfbe1#4-1-%E3%82%B3%E3%83%B3%E3%83%97%E3%82%AC%E3%83%81%E3%83%A3%E5%95%8F%E9%A1%8C
    - https://mathtrain.jp/completegacha
- 余事象
  - 知識というよりは、常に武器として取り出せるようにしておきたい、という心構え。

---

## 問題例

### [yukicoder No.144 エラトステネスのざる](https://yukicoder.me/problems/no/144)

「期待値の線形性」を理解するのにとてもよい題材。  
解説はリンク先のPDFが非常にわかりやすいので、適宜そちらを参照すべし。

**「主客転倒」の概念も絡んでいる気がする。**

とはいえ、「そうなるから、そう」という感覚が正直拭いきれない。  
`X(S)` の定義が巧くて、自分でできるようになるのだろうか？という気持ちになる。

### [yukicoder No.108 トリプルカードコンプ](https://yukicoder.me/problems/no/108)

まずはDPテーブルの定義の仕方に壁がある（とはいえものすごい典型なのだとは思う）。 
個人的には、その次の「3枚以上引いてるカードは引かない場合の条件付き期待値、および条件付き確率」を考えている部分に壁がある。

ほぼ写経でコードを書いてしまったため、DPテーブルの定義は独自のものにアレンジしてみたが、自分のコードのほうが直感的でわかりやすい気がする。

**同じ状態にループしてしまうような遷移は省いて考える」というのが肝なのか？**

**「確率 `p` で起こる事象について、それが起こるまでの期待値は `1/p` 」というのは超典型なので暗記。**

参考: https://drken1215.hatenablog.com/entry/2019/03/23/175300

### [EDPC J.Sushi](https://atcoder.jp/contests/dp/tasks/dp_j)

トリプルカードコンプと全く同じ問題と思って良い。

### [ABC184 D.increment of coins](https://atcoder.jp/contests/abc184/tasks/abc184_d)

これもトリプルカードコンプやSushiと同じ。  
ただし、条件付き期待値や条件付き確率を考えなくてもよいあたり、この問題は確かに易しめだった。

### [SoundHound Inc. Programming Contest 2018 C.Ordinary Beauty](https://atcoder.jp/contests/soundhound2018-summer-qual/tasks/soundhound2018_summer_qual_c)

「期待値の線形性」の最初の練習に。  
`m-1` 個の区間1つ1つが01をとる確率変数だと思って、独立に計算して最後に足し合わせる。

### [ABC008 C.コイン](https://atcoder.jp/contests/abc008/tasks/abc008_3)

エラトステネスのざると一緒に勉強したい問題。  
ただし、ざるよりも確率の計算部分で頭を柔軟にする必要がある。

コイン一つ一つについて01をとる確率変数だと考えればよいのはすぐわかる。  
それぞれについて、「自身を除いた約数について、自分より左側に偶数個並ぶ確率」を求めることになるが、
これを素直に確率の定義に従って「分母は `n!` で分子は〜」と考え始めると泥沼。  
確率を正確に捉えてもっとシンプルに考えれば良い。

この場合は、自身および約数以外はないものとして考えればよい。  
また、約数についてもすべて同質のものと考えて1列に並べてしまい、それらの間に自身を挿入する、という状況を考える。  
すると約数の数を `m` としたときに `Ceil(m+1, 1) / (m+1)` で求まることがわかる。

場合の数に比べて、確率は考え方によっていくらでも簡単になったり難しくなったりするので、そこは訓練が必要。

※なんか2018年7月とかに通してた。こんな教育的問題を雑に消費してしまっていたなんて。。

### [ABC078 C.HSI](https://atcoder.jp/contests/abc078/tasks/arc085_a)

「確率 `p` で起こる事象について、それが起こるまでの期待値は `1/p` 」という超典型。  
できれば `e = (e+1)*(1-p) + p` という漸化式ももう暗記してしまうのが良い。

### [AGC049 A.Erasing Vertices](https://atcoder.jp/contests/agc049/tasks/agc049_a)

期待値の線形性だが、難しい。  
多分「この問題は期待値の線形性を使います！」って宣言されてても難しい。

[けんちょんさんの解説](https://drken1215.hatenablog.com/entry/2020/11/15/104400)が、思考過程も書かれており、勉強になる。  
最終的には、以下の確率変数が定義できると（なんとか）理解できる。

`X[v] := 頂点vが操作過程で削除対象として選ばれるとき1、そうでないとき0`  
こうすると求めるものは `E(X[0] + X[1] + ... + X[n-1]) = E(X[0]) + E(X[1]) + ...` となる。  
「ノードごとに「それが1回の操作としてカウントされる割合」」を求めていると、なんとか合点がいく。  
とはいえやはり難しい。。

「一本道のグラフ＝パスグラフ」というのは、地味に知識として漏れていたので覚えておく。  
構築などでも極端な例としてすぐ出てくるようにしたい。

参考: [基本的なグラフの種類](https://malibu-bulldog.hatenadiary.org/entry/20090516/1242455204)

※なぜか再帰関数と使って `v` に到達するノードの数を数えたら、半分以上のケースが通らなかった。  
グラフの形状によってはこれでは巧く数えられないのか？？

### [第6回ドワンゴからの挑戦状予選 B.Fusing Slimes](https://atcoder.jp/contests/dwacon6th-prelims/tasks/dwacon6th_prelims_b)

期待値の線形性。  
以前解いたことがあったが、理解度確認のために解き直した。

各区間について期待値を求めて足し合わせればよいが、期待値の具体的な計算方法をエスパーしてしまった。  
PDFの説明が簡潔で「スライム `i` がその区間を通るための条件」というのを、全てのスライムに考えてやるのが良い。  
ちょっと戸惑うが、例えば `i` を含めて右側に `4` 個スライムがいると考えると、全体が `4!` 通りで、
`(4-1)` 個のスライムが先に選ばれる通り数は `3!` 通り、よって確率は `3!/4! = 1/4` というふうに求まる。

**この問題のある区間の期待値部分は「コンプガチャ問題」が絡んでいると思われるので、やはりあの有名問題についても暗記してしまったほうが良さそう。**

※こういった問題で、最後にかけるべきものをかけるのを忘れる、といった凡ミスをなくしたい。。  
※解説放送のりんごさんの考え方も面白い（期待値漸化式から導出している）。

### [M-SOLUTIONSプロコンオープン C.Best-of-(2n-1)](https://atcoder.jp/contests/m-solutions2019/tasks/m_solutions2019_c)

純粋な期待値を求める問題。  
しかも（基本的には）定義どおりに期待値を求めれば良い、という問題。  
※りんごさんは↑のように解説放送で説明していたが、厳密には「条件付き期待値」の計算が求められているように思える。  
。。いや、条件付き期待値とは別のなにかの計算が求められているようにも見える。

以下の発想が重要となる。

1. 「引き分けがない場合」をまずは考える。
2. 次に「引き分けでないゲームが `m` 回行われた」という前提で、「引き分けを含めて何回試合が行われるかの期待値」を考える。
3. 「確率 `p` で成功する試行について `m` 回成功するまでに要するトータルの試行回数の期待値」は `m * (1/p)` である、というシンプルな事実

1,2の間で `(p, q) = (a/100, b/100) => (p', q') = (a/(a+b), b/(a+b))` というふうに、考える確率が（一見して）変わって見えるのが難しい。
