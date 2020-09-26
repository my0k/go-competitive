# AGC046-B. Extension

数え上げDP。重複なく数え上げる練習。

ダブった分を引く解法は適当に考えても通ってしまうが、そればかりに頼ると運ゲーっぽくなってしまう気がする。
なので、できるだけ基本方針は「重複なく数え上げる」方向をまずは検討したい。

Editorialでもそのような方法を提案しているが、ちょっと慣れてなくて難しかった。

## 重要なポイント

「何らかのパターンが与えられたとき、それを得るための操作列を復元する」という方針。
もちろん、具体的なパターンはわからないので、 `i*j` マスの状態でまとめてしまう。

※操作列というのは、例えば「行を増やして `j` 列目を黒にする」のような操作の列のことと思えば良い。

このように考えると何が嬉しいかというと、大体以下のような感覚になる。

- パターンを一意な操作列とみなす
- すると、次の操作によって得られる操作列の数を数え上げられる

数え上げDPでよくあるような、樹形図の数え上げのようなイメージになる（というのが自分の感覚、どう言語化するのが適切かわからない）。

そして、これは「一番上の行に複数の黒マスがある（もしくは一つもない（初期状態））」もしくは「一番上の行にただ1つだけ黒マスがある」
の2状態を考えることで、上手く遷移を考えることができる。
すなわち、前者であれば直前の操作は列を増やす操作であることが確定し、後者であれば行を増やす操作であることが確定する。

あとは、選べる黒マスの数を適切に乗算して足し上げていくことで、最終的な答えが求まる。

※初期状態は「A×Bマスすべてが黒マス」と考えたほうがわかりやすい。
初期状態は明らかに特殊な状態だが、このように考えると自然に全体の遷移を統一して考えることができる。

## 参考

[maspyさんのブログ解説](https://maspypy.com/atcoder-%E5%8F%82%E5%8A%A0%E6%84%9F%E6%83%B3-2020-06-20agc046)では、
別の手法による「重複なく数え上げる」方法を提案している。
しかも、こちらはトップダウンなので、再帰関数による実装になる（はず、多分）。