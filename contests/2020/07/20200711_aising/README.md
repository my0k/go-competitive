# エイシングプログラミングコンテスト2020 感想

ワースト2位のパフォを叩き出してしまう。
ちゃんと向き合って次に活かす。

## A問題

企業コンにあるまじきやるだけその1。

## B問題

企業コンにあるまじきやるだけその2。

## C問題

難しかったので飛ばしてしまった。

過去の結果を使う発想は合ったので、すぐに答えに至りたかった。
無駄なく計算するには、1つ1つ求めるのではなく、範囲をすべてなめて答えに含まれる可能性のあるものを、
すべてカウントアップしておけば良い。

出力時に求められているものだけを出力すれば良い。

## D問題

水色とは思えないやらかしをした。

`1~n` のforループ中のサブルーチンで `1~n` の愚直計算をやらかしてしまい、
なんと60分以上気づくことができなかった。

以下、個人的な反省点。

- そもそも色々な箇所で注意力が必要な汚い実装になってしまったのが、注意力がそれた原因かもしれない。
- **ループ中でのサブルーチンは特に計算量に注意が必要！**
  - 結局の所これぐらいしか言えることがないが。。

