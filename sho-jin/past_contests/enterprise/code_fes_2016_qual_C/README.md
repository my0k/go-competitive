# CODE FESTIVAL 2016 qual C 過去問感想

※A, Bは2019-12-10に解いた。

- A問題は愚直に `O(n)` を考えた。
  - 解答PDFでは先に二重ループでの全探索が紹介されていたが、どちらもバグ発生率は変わらなそう。
- B問題は優先度付きキューで考えたが、こどふぉっぽいなと思った。
  - 200点の中では最強クラスでは？
  - 模範解答が転載パズルっぽい（言い過ぎ？）
    - 答えは、最大の `A[i]` を `m` とすると `Max(m-1-(K-m), 0)` と求められる。
    - 考え方は、以下のような感じ。
      - `m` 個のある種類のケーキを並べる。
      - この間にそれ以外のケーキを並べることができれば、連続するものはなくなる。
        - そうでない場合は、最も多い種類のケーキが連続してしまう。
        - 一度埋まってしまえば、別の種類のケーキは適当に連続せずに並べてしまえるので、このようにして簡単に計算できてしまう。
- C問題はコーナーケースに1時間位苦しめられた。。
  - こういった矛盾を検知するような問題は、計算量に余裕があるのなら、誤りを侵さない範囲でキツめに判定を加えておいたほうが良さそう。
  - 今回は、標高を確定させた段階でも、元の情報との整合性をチェックすべきだった。
  - 簡単だと思うが、その分本番であたったら地獄だったと思う。。