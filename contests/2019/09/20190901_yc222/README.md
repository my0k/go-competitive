# yukicoder222（セグ木コン） 感想

セグメントツリーに関する問題のみのコンテスト。
かなり有用な問題集なので、時間を見つけて全部解いてみたい。

- A問題。これだけなんとか解けた。
  - 盛大に誤読してしまった。
  - 最終的にかなり汚い実装になったが、どうするのが良かっただろうか？
    - ノードを構造体にするほうが良かったかも？
- B問題は区間更新から遅延セグメントツリーに走ってしまった。
  - これでも解けるっぽいが、なぜかサンプルが合わない。
    - ノードの境界部分の処理をきっちりやらないとおかしくなってしまうため、そこらへんでミスってるのかも。。
  - スマートな解法は階差数列を取ることだが、これでも境界部分で頭を使う必要がありそうなので、こちらでもちゃんと解いてみること。

