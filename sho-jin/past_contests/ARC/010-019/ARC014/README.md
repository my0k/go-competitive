# ARC014 過去問感想

Last Change: 2020-02-27 01:12:58.

## A問題（@2020-02-27）

ただの偶奇性の判定。

## B問題（@2020-01-27）

制約が小さいので、しりとりを素直にシミュレーションする。
ルール違反の際の勝敗判断は2の剰余を考えた（最下位ビットを見るほうが競技プログラマっぽいかも）。

一度使った単語を辞書に詰め忘れたため、1WAしてしまった。やばい。

## C問題（@2020-02-20）

制約が際どいし、特に部分点なんかは明らかにビット全探索してシミュレーションしてください、という感じだったので迷った。
なにかしらのDPを考えるのかと思ったが、翌々考えたら大抵の同じ色は消せることに気づく。

結局、それぞれのパリティの和が答えになる。ギャグっぽい。
