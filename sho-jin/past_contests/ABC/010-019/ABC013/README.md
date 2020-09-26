# ABC013 過去問感想

Last Change: 2020-04-09 00:34:05.

## A問題（@2020-03-05）

runeをintに変換する方法を学ぼう問題。

## B問題（@2020-02-06）

増やし続けるか減らし続けるかの2パターンのうちのどちらかなので、両方試す。

こういうの結構苦手。
自分は `b+10` して `a` から増やす、というのでやろうとした。

。。いやなんかWAになってた。

こういうのはうまい性質を見つけてやろうとするとドツボにはまるので、
素直な全探索をすべき。。

## C問題（@2020-04-09）

全探索なりDPなりしたいが、DPをするには満腹度のMAXが大きすぎる。
またシンプルな全探索は `O(n^2)` かかってしまうように見える。

詰めると実は質素な食事を何日するか？だけを全探索すれば良いことがわかる（1つ決まるともう片方の最適な戦略が定まる）。
よって `O(n)` でも止まる。

。。が、自分の詰め方だと結構場合分けしなければならず、2WAしてしまった。
