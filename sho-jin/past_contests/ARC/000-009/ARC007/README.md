# ARC007 過去問感想

## A問題（@2020-03-15）

rune型を扱えれば、やるだけ。

## B問題（@2020-01-30）

シミュレーションをするだけだが、行われる操作は正確に把握すること。

速解きに失敗しているので、集中すべき。。

## C問題（@2020-03-07）

実装が結構重かった。これも「パズルのお手伝い」みたいな感じだと思うが、あちらよりも難しいと思う。

解説スライド等がないので、正しい解法かどうかはわからない。

全探索をやる。
`x` 台用いるときは、最初の1台は固定し `x-1` をずらすことになる。
そのずらし方を `x` ごとに全列挙する。

OKかどうかの判定は、ずらした後に `2*n` 個分だけマス目を埋めてみて、
後半が `n` 個以上 `o` が続くように埋まっていればOKと判定できる。

。。判定部分の正しさの証明が正直良くわからない。なんとなく正しいだろう、ぐらいの感触。

**長いコピー先の配列の任意の区間に対して、短いコピー元の配列を写すという操作のコーディングは結構重要だと思う。**
慣れないと意外と難しいが、開区間を意識して関数化すると、割といい感じになると思う。

