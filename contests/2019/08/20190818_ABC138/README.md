# ABC138感想

21分4完でパフォ1200ちょい、レート微減。

Eは実装が複雑だったが、正しく考察できていれば複雑な部分は回避できていたので、
思考の流れについて反省したい。

- A問題は条件分岐を素直にするだけ。
- B問題は正直苦手なやつだけど、サンプルで素直に類推する。
- C問題はハフマン符号っぽさのある貪欲法。
  - 以外と難しいので解説をちゃんと読みたい。
- D問題は遅延伝搬をうまく使う根付き木の問題。
- E問題は割と素直にやればよいが、シミュレーション部分にミスが有り、時間を書けて嘘解法を作ってしまった。
  - 出現位置のカウント配列がソート済みであることをもっと有効活用すべきだった。
    - 二分探索をスムーズに用いれば、自分の解法の嘘に気づけたかもしれない。。

## 根付き木の定義

Wikipediaより。また忘れてしまっていたので。

> あるノードを選んで、それを一番「上」にあると考えると、
> そのノードを基準として2つのノードに上下の関係を考えることが出来る
> （すべてのノードの組み合わせについて定義されるとは限らない）。
> このとき、その一番上のノードを根（ね、英: root）という。根を持つ木を単なる木と区別して根付き木という。

