# ABC105過去問感想

- A問題は特になし。
- B問題はなぜか全探索の解法が思い浮かず、動的計画法しか頭に出てこなかった。
  - 普通にそれぞれの個数を `n, m` と置いたときにそれぞれ取りうる範囲が2桁程度に限定されているのだから、それらについてfor文回して等値判定するだけだった。。
    - 個人的にかなり反省すべき問題。
    - **方程式ちょいと書けば自然とそう解くはずなので、やはり紙とペンは大事。**
- C問題は `-2` 進数変換という問題。
  - 単に2進数変換するノリで計算しても正しくならない（負数の剰余は定義がブレる？）
  - PDF通りに解いた。
  - 本番参加していたら諦めずにロジカルに考えて頑張りたい。
  - **記数法について振り返られるからいい問題だと思う。**
    - `-2` の商をとったあとの世界では、再帰的に同じことを繰り返せば良い
  - **一言まとめ: 偶奇性に着目するのは呼吸と同じ。**
- D問題はほとんど考察があっていたのに仕留めきれなかったので猛省。
  - 累積和にあまりとか計算したりするうちに、どこからどこまでの和なのか？とかがごちゃごちゃになってしまった。
    - 図示するのに加えて、日本語も必要に応じて書いて、おつむの整理をしなきゃいけない。
  - あまりが0のときはmapのインクリメントとanswerの加算の順番を入れ替えないとうまくいかないケースがある気がするが、解説放送では一緒くたに扱っていたが。。？
    - 今となっては確かめるのがちょっと難しいのでスルーしておく。
      - 試しに放送と同じようにしたコードを書いてみたがWAした。。なぜ。。？
      - 微妙にやってることが違うようにも見える（累積和計算する部分）が、今日は仕事で疲れたのでおしまい。
  - **一言まとめ: あまりが等しいもの同士を引くと割り切れる数になる。**