# ダイクストラの実装方針

抽象化できるところは抽象化して、本番でバグらせる可能性を減らしたい。

## 抽象化のポイント

- 順序が定義可能な構造体 `V`
  - priority queueやDP配列の値の型に該当
  - priority queueの実装の都合上、ポインタを用いる必要がある
- `V` の無限大の定義
  - DP配列の初期化に用いる
- `V` の順序定義用の関数 `Less()`
  - strictly lessが表現できれば良い
- 次のノード遷移を経たあとの `V` の計算関数
  - `Progress()` とか？

## 続き（@2020-08-09）

- エッジコスト、あるいは次のノードの値について、、
  - 0に該当する値の定義
    - 無限大の対になる？
  - 非負であることの定義
    - デバッグ目的で、この概念を含めておくと良いかもしれない
- 初期化処理なども関数で渡してしまう？
- `Less` 関数をpriority queueのメソッドとして共通化している部分をどうするか？
  - プログラム実行時に動的に上書きしたい
    - Goのメソッドはたしか第一引数がレシーバだった気がするので、不可能ではないようにも思える
- 必要なパーツを関数に引数として与える形だと、無限大・ゼロに相当する値は自然な構造体インスタンスにできる
  - 渡すべき関数オブジェクトはせいぜい2,3程度になるので、そのほうがきれいになるかもしれない

## 続き（@2020-08-12）

- `vinit` は始点に紐づくものなので、始点のリスト `S` に一緒くたにして渡したい
  - 零に相当するものではなかった（多分）
- 次のノード遷移で `V` が減少していないかどうかのチェックコードは `GenNextV` に書くことになりそう
  - これはこれで柔軟で良い（panicさせたり、その時々で選べる）

## 参考

http://kutimoti.hatenablog.com/entry/2019/04/27/225727
