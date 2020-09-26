# AGC010過去問感想

- A問題はAGC010から030までのA問題だと最易だった。
  - しかしながらもしもN=1が制約に入っていたらコーナーケースで沈んでいたので、本番ではココらへんも集中して見ていきたい。
- B問題は久々の500点自力で嬉しい。
  - まずは操作回数が定まることはすぐに見抜けた。
  - 部分和を矛盾なく構成できるか？は計算量的にやばそうなので考えるのをやめた。
    - もっと早くやめたいところだった。
  - 操作の差分が単純なことから、隣接項間の差分に目をつけたらうまくいった。
  - 十分条件が見抜ききれずに1回WAを出してしまったのが悔やまれるところ。
  - **典型チック: 隣接項間の差分に着目、操作回数の固定、問題の単純な部分に目をつける**