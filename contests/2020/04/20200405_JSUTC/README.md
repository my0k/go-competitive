# Judge System Update Test Contest 202004 感想

言語アップデートもあったのでunratedだけど参加してみた。

- AはなんかC++なら新機能が使えるらしいが、普通にやるだけ。
- Bはボールの種類を分けてそれぞれソートするのが楽だけど、無名関数渡すソートを使いたかったのでソート関数を定義してやった。
- Cは制約から全探索。順列をすべて当てはめて試す。
  - [フック長の公式](https://ja.wikipedia.org/wiki/フック長の公式)とかいうのがあるらしい。知らない。
- Dはちょっと考えたが、累積GCDを前計算しておくと、求めたい箇所が二分探索で求められる。

