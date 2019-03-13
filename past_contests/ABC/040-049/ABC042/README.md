# ABC042感想

- A問題は辞書を使って解いた。
  - ソートよりも早いと思うのでこれで良いと思う。
- B問題は一瞬戸惑うが、文字列スライスをソートして、それらを連結していけば良い。
- C問題も辞書式順序で解くように見えるが、真面目にやると多分かなりめんどくさい（できない？）
  - そこから制約に目を向けて、ちゃんと全探索に切り替えられたのは偉い。
  - Pythonっぽく、integerスライス中の存在判定（in演算）をするための関数を作りたくなるが、これは即席で辞書を使えば作れる。
- D問題は逆元の事前計算までは考えられたが、計算量の推定が間違っており、解答を読んでしまった。
  - 解答PDFより: **二分累乗法のオーダーはO(1)ではあるが定数倍が30倍程あることに注意せよ。**
    - **ある数値に対する逆元の計算はちょっと大きめの定数オーダーと再認識する。**
      - 二分累乗法は対数オーダーという認識が強く、`O(logN)` と勘違いしてしまい、計算量を過大評価してしまった。
        - 。。あとあと冷静に考えたら、仮に `O(logN)` だったとしても通りそうだったが。。
  - また、場合の数の求め方についても間違っており、解答を読んでも通すのに時間がかかってしまった。
    - **排反事象となるように切り分けないと、重複した数え上げが発生してしまい、正しい答えにならない。**
  - そもそも、当初グリッドではなく格子点で考えてしまったので、それも反省。
    - このミスは今後につなげたい。
  - **逆元の練習問題としてよい。**
    - ABC110のD問題でも逆元は必要だが、あちらは別の考察も必要となるため、逆元の練習を目的とするとこちらよりもしんどい。
