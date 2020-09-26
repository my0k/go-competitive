# ABC042 D. いろはちゃんとマス目

Last Change: 2020-05-16 00:25:27.

## @2020-05-16

一発AC。

以前解いたときは数え上げのパートで排反に考えられておらず、重複なく数えることができなかったが、
今回は正しく整理できた。

移動不可のマス目の右上から1つずつ斜め上に関節点となるマス目を移していき、
それぞれについて組み合わせを計算してやれば良い。
マス目の番号には注意（ちゃんと紙に書いて整理しよう）。

Editorialとは違う事象の分割だが、排反であればちゃんと計算できる。
