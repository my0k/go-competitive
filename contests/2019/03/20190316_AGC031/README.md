# AGC031感想

感想書くのを忘れていたが、A問題を15分程度で解いて、
あとは2時間以上椅子に座って考えていただけだったので、翌日頭痛が止まらなかった。

- A問題は整理すると単純な場合の数の問題になる。
- B問題は後日考えた。
  - 前から見ていったときの通り数を記録する1次元DPで解けてしまう。
  - 反転する際は、前の方を調べて一番近いものだけをペアとすれば良い。
    - この位置を調べるのは前処理として `O(N)` でできるので、全体の計算量も `O(N)` になる。
    - 過去に算出した結果の再利用の際の、添字の付け方にはよく注意する必要がある。
    - **また、MODとるのも忘れずに！！！**
  - この問題もそうだったが、やはり **小規模な問題を「DPっぽい方法」で具体的に紙に書いて解いてみる** ことが大事だと感じる。
    - 一般化はそれをヒントにやっていけば良い（それこそ「動的」に）。
  - **樹形図も大事だと感じる（経路や文字列に1文字加えただけのものは通り数が増えたりはしない、という感覚とか）**
  - 一応、今後につなげるために[けんちょんさんの詳細な解説記事](http://drken1215.hatenablog.com/entry/2019/03/17/130800)も拝読しておく。
    - 。。読んでみたが今回はあまり響かなかった（区間を扱う経験値がまだ足りないせいな気がする）。
