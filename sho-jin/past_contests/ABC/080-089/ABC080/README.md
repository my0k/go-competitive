# ABC080過去問感想

- A, B問題は特になし。
- C問題はビット全探索。
  - 最近ABC114C問題で爆死したので非常にタイムリー。
  - 練習も兼ねて再帰で解いたが、問題的にもビット演算で解いたほうが早く解けそうな気がした。
- D問題は例題で図を描くことで割と早く解答を見つけられた。
  - ABC086のD問題で調べていて知ったことだが、imos法に関連（？）
  - 累積和をとるためにある区間を1で塗りつぶす（配列の各要素を1つずつ設定する）ということをしているが、これがimos法だと効率的に行える。
    - 見返してみたところ、嘘解法でACをとってしまったかと思ったが、今回の場合、塗りつぶしは最大でも `10^5*30` 程度であるため、十分に間に合った。
    - オーダーはこのようなナイーブな方法だと `O(N*C)` だが、imos法を使うと `O(C+N)` になる。
    - 1次元のimos法は塗りつぶしを最後にまとめて行うイメージ。
- トータルで40分程度だったので、この調子を維持したい。
