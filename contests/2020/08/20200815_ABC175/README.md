# ABC175 感想

かなりギリギリだったが、なんとか95分くらいでEまで解けてパフォは1700ちょいで生還。

なんかスポンサードの関係か企業コンABC級という感じがして、全体的に面倒な問題が揃っていた。

## A問題

3文字だけとはいえ場合分け8通りはだるすぎるので、ランレングス圧縮した。

## B問題

3重ループだが、よく見ると辺の長さがすべて異なる場合の通り数を数えよとあり、気づくのが少し遅れた。

## C問題

難しいと思った。確信が持てず、提出時にかなり祈った。
一応、自分の考察は正しかった。

一般性を失わずに、初期位置を非負にしてしまって良い。

できるだけ原点に向かって移動した後は、残りの回数の偶奇性で決まる。

## D問題

フルフィードバックじゃないとおそらく正解できなかった。

一周したときにプラスになるときはできるだけそれを使ったほうがいいが、
実は一周だけ余らせると、より先に進むことができて得をする場合がある。

ここは、admin解答にあるように、「着地点について全探索（到達できないのであればスキップ）」とすると、
もれなく答えが調べられていいように思えた。

### 解説放送で勉強

- 順列なので、グラフはいくつかのループがある形になる
  - あるノードの行き先は1つだし、順列なのであるノードに入ってくるものも1つなので。
  - **このような設定の問題はほかにもありそうなため、これはちゃんと抑えておきたい。**

### MMNMMさんの解法

ダブリングで桁DPっぽくやってるとのことなので、参考にしたい。

https://atcoder.jp/contests/abc175/submissions/15960077

なんとなく想像はつくが、Kを2進数解釈したりして結構めんどくさそう。。

## E問題

`dp[i][j][l] := (i, j)の品物を拾うことを考慮した上での、(i, j)まででi行で合計l個拾った状態における最大値`

`(0, 0)` の例外処理が面倒だった。
20分ぐらいかかってしまったが、もう少し素早く解けるようになりたい。

## F問題

なんかダイクストラらしいので、ライブラリチェックを兼ねて解いてみたい。

