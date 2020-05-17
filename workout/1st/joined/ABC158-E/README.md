# ABC158 E. Divisible Substring

Last Change: 2020-05-16 19:58:17.

## @2020-05-16

ちょっとサンプルを通すのに手間取ったが一発AC。
（そもそも累積的に観てあまりが0のものを加算し忘れていた。）

### 合同式の性質

この問題で一番重要なのは、合同式に関する以下の性質（[参考](https://drken1215.hatenablog.com/entry/2020/03/08/020200)）。

`x = y (mod m) => ax = ay (mod m)` は一般的に成り立つが `ax = ay (mod m) => x = y (mod m)` は成り立たない。
後者が成り立つのは `m, a` が互いに素の場合のみ。
