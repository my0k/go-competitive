# DISCO presents ディスカバリーチャンネルコードコンテスト 2016 過去問感想

- C問題は一見頑張ればできそうに見えて、色々と知識がないととても解けるとは思えない問題だった（当社比）。
  - まず高度合成数に関する知識がないと辛いと思う。
    - [高度合成数](https://ja.wikipedia.org/wiki/%E9%AB%98%E5%BA%A6%E5%90%88%E6%88%90%E6%95%B0)
      - 端的にいうと **自然数で、それ未満のどの自然数よりも約数の個数が多い自然数のこと**
        - 簡単なところで言うと、`1, 2, 4, 6` などは高度合成数。
        - ラマヌジャンによって考察されたらしい（意外と最近）。
    - `10^9` の約数はたかだか1344個に抑えられるとのこと。意外と少なくて驚いた。
  - さらに、 **「`x, y` の積が `K` の倍数であることと、`GCD(x,K), GCD(y,K)` の積が `K` の倍数であることは同値」** という定理が必要。
    - 上の定理において逆の理解は直感的であるため、覚えるとしたらこちらだけを覚えておきたい。
      - 「最大公約数同士の積が `K` の倍数じゃないのに、元の数同士の積が `K` の倍数であることはありえない」というのは背理法で証明できるが、かなり難しい。
  - これらの知見を組み合わせると、各 `Ai` の値の `K` との最大公約数の数も1344個しかなく、これらの分布を調べて、丁寧に組み合わせを計算していけば良い。
    - 実装も若干気をつけなければいけない点がある。
    - 制約の都合上、連想配列を使うことになるが、数え上げの際は小さい値から2重に、重複なく走査する必要があり、そのために別に連想配列のキー配列を用意する必要がある。
      - こちらは単なるint型のスライスとしておいて、最後にソートしておけば良い。
      - 難しくはないが、初めて使ったテクニックだったので、覚えておきたい。
  - しばしば悩む最大公約数部分の計算量について。
    - `GCD(x, K)` の部分が `O(log(K))` というように見積もられているあたり、大きい方の数の対数とおぼえておけば良さそう。

---

## C問題の定理について

`GCD(x, K) = x', GCD(y, K) = y'` とする。

`x', y'` は見方を変えると、「**`x, y` がもつ `K` のエッセンス（素因数）**」とも言えるような気がする（直感的解釈）。

そのような見方をすると、 `K` のエッセンス同士をかけ合わせて `K` の倍数になれないのであれば、もとの数同士の積も `K` の倍数とはなりえない、と考えることができる。

逆に、`x, y` の積が `K` の倍数である、というのはもう少し細分化して考えると、先に述べた `x', y'` の効用が全てである、と考えられる。
（`x/x', y/y'` はそれぞれ `K` に関して言えば、エッセンスを剥ぎ取られた「残り滓」と言えるもので、`x/x', y/y'` の積は `K` に寄与することはない。）

