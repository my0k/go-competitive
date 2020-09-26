# ABC073過去問感想

- A, B問題は特になし
- C問題もマップを使ったシミュレーションを行えば簡単だった
    - N=100000のため、ソートを使っても良い（解いてるときも考えたが、こっちのほうが実装の時間が長くなりそうだったのでやめた）
- D問題はワーシャルフロイドと順列を駆使する問題
    - 解法は制限からあからさまな気がする
    - 順列は以前の自作関数を使ったが、こういう場合取り出した値がなにか脳内だけで管理し続けるのは困難なので、贅沢にわかり易い名前をつけた一時変数を積極的に使っていく
        - 必要な場面でスライスを使い忘れるミスをしてしまい、テストを通すまでに時間がかかった
        - コードゴルフとかは考えない
    - 400点問題はそろそろスピードを意識しだしてもいいかもしれない

## D問題の復習（2019/07/20）

問題自体はワーシャルフロイド＋アルファという感じでとても易しいが、ワーシャルフロイド法自体ですこし気になったので復習した。

収穫として、**「DP配列を3次元にして過去の更新を残す方法では、1-basedな頂点IDをを使わないと厳しいかもしれない」** ということがわかった。

（特に必要がなければ、配列再利用版を普通に使いましょう、という気持ちになった。）