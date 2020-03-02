# ARC021 過去問感想

Last Change: 2020-02-23 12:05:28.

## A問題（@2020-02-08）

懐かしのあのスマホのミニゲーム。

すべてのマス目の隣り合うマス目を調べて、同じ数字があるところがあればOK。

バグが怖いので、グリッドのBFS的なものを書いて全探索した。

## B問題（@2020-02-23）

XORの性質を問う問題。一応構築問題に分類される？

解の判定は全部のXORを取ったら0になるかどうか？でOK（のはず。。嘘解法ではないと思う。。）

辞書順最小のものについては `A[0]` は必ず `0` に設定でき、それが決まれば連鎖的にすべての `A` が決定する。
（累積XORとか計算しつつ考える必要があると思ったが、よく考えたら不要だった。）
