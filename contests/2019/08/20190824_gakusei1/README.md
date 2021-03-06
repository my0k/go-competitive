# 第一回日本最強プログラマー学生選手権予選 感想

30分2完でパフォは1300ちょい。
500, 600はちょっと自分にはきつかったが、
2問早解きさえできれば十分にパフォ1500は取れたコンテストだったので、
反省するところはしっかりしたい。

- Aは素直に全探索。
- Bは転倒数に関する問題、反省が必要。
  - **「なんかこれ見たことある！」で勇み足でコピペしてもほぼいいことはないので、これはやめること！**
    - 最終的にコピペコードを使うことになったとしても、それはちゃんと考察を終えてからやるべき。
    - 今回はBITを使おうとして時間を無駄に食ってしまった。
- Cはflip系の500点問題。解きたかった。
  - 色々と良い性質に気づいて履いたと思うが、そもそもDPしようとしてしまったのがダメだった。。
    - DPする方法もあるのかもしれないが、自分には経験値がないと捨てるべきだった。
  - 重要な問題と思うので、解説放送で今後のための知見を得る。
- Dは二部グラフの知識に帰着できるらしい、、がグラフ理論をほぼほぼ忘れてしまっているのでピンと来ない。。
  - これも二部グラフに対する理解が深められると思うので、しっかり解説放送で理解していく。

---

## C問題

- **典型: 反転操作の捉え方のパターン**
  - `[l, r] -> [0, r], [0, l-1]`
  - これは覚えておいたほうが良さそう（発想に繋げられる）。
    - **左を固定できるのは整理するときに強い、という印象がある。**
- 今回は反転操作と、各マスを右端・左端としたとき、各マスについて左端・右端のどちらにすべきかを明らかにできる。
  - 難しく見えるが典型らしいので、とりあえず理解だけはしておくこと。
- 最後の数え上げのやり方も何気に重要っぽい。
  - ただただ掛け算をしていくだけでよいが、あまり自分はやったことのない形なので、必要があれば次につなげたい。
  - 左端・右端の位置の妥当性についても、簡単な例外処理をしないと嘘を出力してしまうので、ちょっとだけ注意が必要。

## D問題

- 二部グラフに関するセンスが必要（？）
  - 改めて定義: 頂点集合を2つの部分集合にわけたときに、集合内に辺が存在しないもの。
    - 集合間ですべての点対に辺がある場合は「完全二部グラフ」となる。
  - いままであまり馴染みがなかったせいか、なかなか二部グラフと単語を効いてもグラフの形がイメージできない。これはまずい。
    - 関連する性質の前に、まず **集合内で辺がそんざいしないこと** を強烈に念頭に置きたい。
      - 集合間の点対の辺の有無についてはとりあえずは気にしなくて良い。
      - 二部グラフの定義には抵触しないので。
- 解説放送を聞いても難しいと感じるが、要点を整理してみる。
  - **あるレベルに着目したときに、その部分集合のグラフは二部グラフとなっている必要がある。**
    - 距離が奇数となる閉路を含んではいけない場合は二部グラフとなる、というのは知識としてないとちょっと厳しい。。
    - [wikipedia](https://ja.wikipedia.org/wiki/%E9%96%89%E8%B7%AF%E3%82%B0%E3%83%A9%E3%83%95)より、奇閉路・偶閉路というワードはちゃんとあるらしい。
      - 奇閉路: 頂点が奇数の閉路のこと。
- 各レベルにおいて二部グラフになっている、かつ、そのような状態で完全グラフのすべての辺を網羅する、ことが目的となる。
  - これは解説放送や解答PDFでも述べられている通り、ビットのテクニックを使うと巧い。
    1. 各ノードを0-basedな数値が採番されるとする。
    2. ノード番号をビット表記したときに、第 `i` ビット目に注目する（この `i` をレベルと考える。）
      - レベルは正の整数である必要があるため、実際には `i+1` がレベル。
    3. このとき、`i` ビット目の数値が異なるノード同士を結びつけると、そのレベルでは二部グラフが作られる。
    4. さらに、ノード番号はビット列がすべて異なるため、必ずどこかで必要な辺があるレベルによって結ばれることになる。
      - 一瞬、これは理解しづらいが、逆に「ビット列がすべて同じなものは、確実に考えている範囲のレベルで辺が引かれない」と考えるとイメージしやすいかも。
      - 実装時は、ある辺があるレベルによって上書きされるようなものでも構わない。

### 解説PDFより抜粋

この問題に関しては解説PDFの内容が勉強になる。

> 二部グラフについて言えることとして、
> 「頂点を適切に2色に塗り分けることで、同色の2頂点の間に辺が無いようにできる」
> というものがありますが〜

二部グラフの定義を忘れがちだが、2色彩色と併せてイメージを持っておくと、忘れにくいかもしれない。

