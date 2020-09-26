# ABC036過去問感想

- A問題は特になし。
  - 方程式はちゃんと紙に書きましょう。
- B問題は紙に3x3ぐらいのケースを書けば気づける。
- C問題は色々やりかたがありそうだが、ソートとマップを併用して解いた。
  - タイトルにもある「座圧」というのは「座標圧縮」のことらしい。
  - 解説にも詳しい説明はないが、「大小関係の情報だけを保持しつつ、値は小さくする操作」ぐらいの印象はこの時点で持っておく。
- D問題はわかりそうだがわからなかった。
  - tree DPというタイプのDP解法が適用できるらしい。
  - 無向木のDFSはあまり書く機会がなかった気がするが、今回書いたような **どこから来たか？すなわち親のノードIDを再帰関数の引数として渡しておく** と良さそう。
    - 今回の場合、1種類のフラグをもたせるやり方はできない（正しく再帰関数が呼び出されない）
