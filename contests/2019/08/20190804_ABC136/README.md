# ABC136 感想

1300弱のパフォでレートが1下がってしまう。

42分でノーペナ4完なら大丈夫だと踏んで提出したが、かなり厳しい。
D問題は変な解き方をしてしまったようなので、解説放送をみてベストな方針を学びたい。

E問題は苦手な整数・約数の問題だが、これも次につなげるような復習をしたい。

- A問題は簡単な算数。ちょっと慎重になりすぎたかも。。
- B問題は10進数整数の桁数を計算する問題だが、かなりもたついてしまった。
  - それなりに登場しそうだし、ライブラリ化しておいてもいいかもしれない。。
- C問題は貪欲をシミュレーションすればよい。
- D問題は累積和と二分探索を併用したが、かなりバグが怖かったため、慎重になりすぎ時間を費やしてしまった。
  - もっとシンプルな方法があるらしいので、それを学ぶこと。
- E問題は制約も特徴的な問題。
  - 全然わからなかったので解説を見る。
  - 解説放送を見てもPDFを読んでも難しかった。

## 動画で復習（2019/08/18）

### D問題

`R...RL...L` のブロックに分けて、境目の部分に人が集中することに着目する。

実装としては、

- ブロック単位でスキャンしながらRの個数をカウントする
- カウントしたRの数を2つの境目のマスに切り上げと切り捨てをうまく使って振り分ける
- Lについてはもとの配列を反転してもう一度ループしてやれば良い。

。。結構テクニカルに見える。
自分の方法のほうが面倒だけどシンプルな気もした。

### E問題

- まず、 **総和が不変であることに着目する必要がある。**
- 数列のGCDは、数列の総和も割り切る。
  - よって、和の約数が探索範囲といえる。
  - チェックすべき約数の個数が `O(sqrt(N * maxA))` なので判定に `O(NlogN)` ぐらいかかっても大丈夫
    - 。。という思考ができる。
- 判定は「GCDをXにできるか？」と考える。
  - 「GCDをXにするために必要な操作回数（の最小値）」を考え、これが制約以下ならばOK。
  - 更に言い換えて「全部をXの倍数にするために必要な操作回数」を考えるのが良い。
    - 剰余に帰着できる。
- 剰余の問題に帰着したあともかなり難しいと思う。
  - 画像にしたので、主たる性質はそちらで確認する。

#### どう考えるのが良かったか？

- 典型: 不変量に着目する（？）
  - 典型としてしまってよいのか微妙。。
- 典型: 数列の和に着目する。
  - とりあえずの思考として持っておくのはいいかもしれない。
  - 今回だったら、連鎖的に次の思考に繋げられた可能性はある。

#### 実装面の注意

- 約数を考えるときは `sqrt(N)` の上限に引っ張られて、片方の大きい方の約数を見落とすミスを、しばしば自分はやってしまいがちなので、そこは十分に注意したい。
  - この問題でも結構な時間ハマった。

