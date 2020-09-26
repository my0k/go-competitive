# ABC006 過去問感想

Last Change: 2020-04-12 16:36:23.

## D問題（@2020-04-10）

転倒数に着目する問題。。と思ったらそれは嘘解法だった。
多分、反例を示すサンプルは簡単につくれる。

解説通り、よくよく考えると答えはLISの長さそのものだった。

### 学習目的にやってみたら。。

LISの理解を深めるために、二分探索で高速化するバージョンのDPテーブルの定義のまま、
`O(n^2)` のロジックで更新させてみたら、なんと通ってしまった。

`n=30000` と小さいのもあるが、枝刈りも聞いているのかもしれない。
