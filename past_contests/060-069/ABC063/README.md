# ABC063過去問感想

- A問題、C問題は特に今後振り返る必要もなさそう
    - CよりもBのほうが慣れてないと厳しそうな
- B問題は最長26文字のアルファベット小文字から構成される文字列について、すべての文字が異なるかどうかの判定
    - Goだと `rune` が `int32` のエイリアスであることを利用して、配列で情報を管理するのが一番シンプルにできそう
    - やってることは `map` だけど、ほかの問題でも当てはまるように、keyが整数で範囲も狭い場合は配列で良い
        - あと `map` はkeyの存在判定とかがちょっとめんどくさい（Goだとこのあたりも融通がきいたことを自動でやってくれてたような気もするが、詳細は要調査）