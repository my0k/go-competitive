# ABC131感想

Eで3ペナしてしまい、実質70分5完でパフォーマンスは1600弱。
Eのバグらせ方は競技プログラミングのみならず、業務プログラミングでもやりかねないものだったので、対策を真剣に考える。

- A問題は素直に文字列として考え、全パターンをif文に直接書く。
- B問題は賢いやり方もあるだろうが、確実に行きたかったため制約を利用して全探索した。
  - 賢いやり方に比べて大幅に実装量が増えてしまい、実際に手間取ったので、ちょっとは感が他方が良かったかも。。
  - 賢いやり方は、美味しさの絶対値が最も小さいものを選ぶ、というもの。
- C問題はベン図を思い浮かべながら解くやつ。
  - `[a, b]` の閉区間の範囲で条件を満たすものの個数は `C([0, b]) - C[0, a-1]` として求めるのは超典型。
- D問題は区間スケジューリングっぽいが、別に区間でもない。
  - 証明せずに直感的に締め切りでソートして解いた。
  - 証明は練習したいので、復習をちゃんとすべし。
  - [いつもお世話になっているあのお方の記事](http://drken1215.hatenablog.com/entry/2019/06/22/224800)
- E問題はグラフ系の構築問題。考察が早くできて嬉しいからと言って雑な実装をしないこと！
  - まず、サンプルで数字だけを見て、上限である最短距離2のノード対の数を誤って数えないこと。
    - **典型: 最大・最小が見積もれるときは具体的にどういうケースなのかを考える。**
      - 最近この典型意識がうまく働いているので、継続したい。
  - 最もやらかしてしまったと感じたのが、forループの `break` タイミング。
    - ループの後ろ側に書いてしまうことで、ループを一度も回さなくてよいケースをカバーできていなかった。
    - **実装テク: ループの `break` 条件が前でも後ろでも意味が等しい場合は、前に書く。**
    - **実装テク: ループを一度も実行しないケースは典型的なコーナーケースと心得る。**
      - どちらかというとこちらのほうが大事かも。。
  - **実装テク: バグを生んだときは、バグの箇所だけを見るのではなく、上から一度ソースコードを全スキャンする！**
    - syunsukeさんがつぶやいていた内容だが、とても共感できる。性格的にも自分にあっていそうなメソッドであるため、真似してみることにする。
  - **実装テク: それなりに計算が要求される、再度タイピングするには負荷がかかる式は、変数に入れてしまう。**
    - この問題だと、スターグラフにおける最短距離2のノード対の個数である `(n-1)*(n-2)/2` など。
    - これも以降のバグ削減に効いてくるはず。
  - **実装テク（？）: ラベル付きforループを使う。**
    - ネストされたforループを `break` するときにはバグを減らせるのではないかと思う。
    - これについては試験運用が必要なので、ちょっと試し切りしてみる。
- F問題はグラフ問題に帰着するらしい。考え方に汎用性がありそうなので、要復習。

## 単純（かつ連結）グラフについて

**多重辺・自己ループがないグラフのこと。**

### 多重グラフ

**単純グラフでないもの。**

（ひょっとしたら多重グラフは単純グラフを含んでいるかもしれない。）
