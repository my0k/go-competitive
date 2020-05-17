# ABC150 D. Semi Common Multiple

Last Change: 2020-05-15 21:39:52.

## @2020-05-15

難しい。考察を誤ってしまい、Editorialを観ることになってしまった。

Editorialとは少し違った方法をとっているが、おそらく本質的には同じことをしているはず。

まず「`A[k]/2` は奇数である」というとんでもない誤った考察をしてしまい、そこから抜け出せなかった。
やばい。
偶奇性は整数の基本なので、慎重に正しく考えること！

以降も偶奇性について注意深く考察すると、 `x` が存在しないようなケースがたくさん出てくる。
それらを漏れなく処理することが、ACするためには大切（テストケースも90個ぐらいあり、強力）。

例外を排除できたあとの求め方は、最小公倍数の奇数倍となるため、ここは切り上げ処理を入れれば簡単。
