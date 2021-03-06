# 最小全域木

UnionFindがあればクラスカル法のほうがプリム法よりも簡単なため、
基本的にはクラスカル法一択に絞ったほうが良い（と思う）。

とはいえ、問題によっては `O(|V|^2)` のプリム法なども必要になる場合がある。

~~また、クラスカル法自体が「辺のソート→UFを使って追加辺を選択」で終わるため、~~
~~ダイクストラ法とは違いスニペットに頼っても良さそう。~~
SPFAと同じく、問題ごとに要求されるアレンジの幅が広いので、
個々のパーツ（UnionFindやpriority queue）を除いて、やはりゼロから書けるほうが望ましい。

## 問題まとめ

### 練習問題

とりあえずやるだけのもの。応用問題を解くための基礎、前提。

- [AOJ 最小全域木](https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/12/ALDS1_12_A)
  - `O(|V|^2)` のプリム法の練習用。
    - ダイクストラ法と違い、最小全域木を繰り返し求めるケースは少なそうなので、あまり必要ではないかも。
- [AOJ 最小全域木](https://onlinejudge.u-aizu.ac.jp/courses/library/5/GRL/2/GRL_2_A)
  - priority queueを使ったプリム法、もしくはクラスカル法の練習用。

### 応用問題

少なからず工夫が必要なもの。

典型的なパターンは把握しておきたい。

クラスカル法じゃ面倒でプリム法なら簡単、とかはあるんだろうか（あったらやだなぁ）？

- [yukicoder No.748 yuki国のお財布事情](https://yukicoder.me/problems/no/748)
  - すぐに思いつくレベルの簡単な工夫が必要。星の数の割には易しめだと思う。
  - クラスカル法がお誂え向きな問題。
- [第一回アルゴリズム実技検定 L-グラデーション](https://atcoder.jp/contests/past201912-open/tasks/past201912_l)
  - 小さい塔の扱い方がポイントになる。これをすべて考慮してMSTを作ると、嘘解法にしかならない（多分）。
    - 小さい塔が必須であるかのようにMSTが組まれてしまうため、修正ができない。
  - SPFAと同じく、全探索的な思考は常に必要となる。
- [Codeforces Rounde No.597 D.Shichikuji and Power Grid](https://codeforces.com/problemset/problem/1245/D)
  - 貪欲的な発想とダミーノード・ダミーエッジの追加という考え方が必要。
  - 鮮やかだし難しいと思うけど、多分ド典型の類なのか、点数は1900点と感じたよりも低め。
  - **制約が厳しいため、少なくともGolangでは `O(|V|^2)` の愚直なプリム法じゃないと通せない。**
- [ABC065 D.Built?](https://atcoder.jp/contests/abc065/tasks/arc076_b)
  - 不要な辺を減らす考察が最重要で、それが終わればあとは素直にMSTを求めれば良い。
    - 「不要な辺をへらす」というのは高難易度のグラフ問題で考えなければならない点かもしれない。
- [CODE FESTIVAL 2016 qual B C.Gr-idian MST](https://atcoder.jp/contests/code-festival-2016-qualb/tasks/codefestival_2016_qualB_c)
  - Built?とセットで抑えておきたい問題。
    - ちなみにどちらもDEGwerさんがwriterの問題。

### MSTと見抜くために

- とにかく辺の数に注目。
  - 少なければ勝ちやすい。

---

### アルゴリズムピックアップ

#### クラスカル法

端点のノードIDとコストを保持したエッジをソートし、
コストが小さいものから併合する必要のあるものだけを足していけば良い。

併合のチェックもUnion Findの `uf.Same(e.from, e.to), uf.Unite(e.to)` しか使わない。
そのため、ノードIDが0-basedか1-basedかや、ノードの個数上限なども適当でも良い
（連結成分の個数や連結成分のサイズを確認することがないため）。

#### クラスカル法の実装手順

1. ソート可能なエッジリストを構成する。
  - エッジ構造体は `from, to, cost` を最低限備えていれば良い。
2. エッジリストをコストで昇順ソートする。
3. 閉路判定用のUF木を確保する。
  - 確保するノード数は最大値を適当に超えさせておけば良い。
4. エッジリストを先頭からチェックし `!uf.Same(e.from, e.to)` を満たすもののコストだけ答えに加算する。

#### プリム法

ダイクストラ法と似ているが、こちらのほうがアレンジの幅は狭いと思われる（エッジのコストはいじりづらい？）。

ダイクストラ法と同様に、各ノードの状態を色で考えるのが混乱しなくて良い。

`WHITE -> GRAY -> BLACK` の順に変化していく。
ただし、始点（※ダイクストラ法と異なり、任意に選択して良いため普通は `0`）
に関してはループ外ですぐさま更新が必要となる。

- 初期状態は `WHITE`
- キューに入れられるか、一度でもチェックされたら `GRAY` に変化
- キューから取り出されるなどして、木への追加が確定したら `BLACK` に変化

#### プリム法の実装手順

利用ケースが想定される `O(|V|^2)` のバージョンについてのtips。
とはいえ、ゼロから書くのはしんどいので、基本はスニペットを呼び出して適宜書き換えを行う方針で行く。

- エッジコストは隣接行列で管理する。
- ノードIDは0-basedとしたほうが良い。
- スライス `d, p, color` をノード数分確保。
  - それぞれ当該ノードの、親からノードへのエッジのコスト、親ノードID、ノードの状態を意味する。
- ループを抜けたら全ノードがMSTに組み込まれているので、組み込まれた辺の総和を計算する。

