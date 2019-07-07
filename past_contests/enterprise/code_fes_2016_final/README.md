# CODE FESTIVAL 2016 final 過去問感想

- B問題は部分和の初めて見る問題。
  - 細部までは見抜けなかったが、一応模範解答と同じコンセプトで解いていると思われる。
  - 自分の解法では、$(1/2)n(n+1) - A = n$ という方程式を解き、選ばない問題のスコアの和を考えた。
    - こちらはスコアが大きい方から取れるだけ取っていく、ということができるので楽。
    - 証明してないが、なんとなく大きいものを貪欲に選んでいけば、最終的には `A` が作れそうな感じがしたので。。
    - 必ずしも出力が昇順でなくてもよい、というのも微妙にヒントになっている気がする。
  - **模範解答の考え方も重要なので、別に復習しておく。**
  - 部分点は動的計画法で解ける制約。
    - 部分点でDPを積極的にやれるほどまだ強くはないのが残念。。
- C問題は一見難しそうだったが、言語の数の和の制約からUFを見抜いて解けた。
  - このたぐいの制約は積極的に使っていきたい（以前は使わなくても解けるものがあったが）。
  - 解答スライドの説明とは若干違う見方をしている気がするが、自分が実装した手法のほうが簡単にみえる。
    - ざっくりというと、言語ごとに直接会話ができる（連結な）人同士を結びつけていく、というのを全言語に横断してやっていく。
    - 最後にどの人でもよいので、属する集合のサイズを確認して終わり。
