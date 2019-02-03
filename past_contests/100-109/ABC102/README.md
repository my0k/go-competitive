# ABC102過去問感想

- A, B問題は特になし。
- C問題は一応は解けたし、模範解答どおりだったが。。
  - グリーディに分類されると思うが、証明はできていない。
    - 実験をしながら推測できたので試しに通してみたら通っただけ。全然自信はなかった。
  - 解説PDFもグリーディ部分の証明がなかったので、放送で勉強。
  - **動く点をイメージすると良い、データが1次元なのでシンプルな数直線上で距離をイメージするべき。**
    - **点が $\epsilon$ 動くと、総和がどれくらい動くかは、数直線で考えると分かる。**
    - **具体的には、右に動く場合を考えると、左に $a$ 点、右に $b$ 点ある場合、距離の総和は $a*\epsilon + b*(-\epsilon)$ 変動することとなる。**
    - これがイメージできると、中央値のところまで動く過程では、常に距離の総和は減少することが分かる！
      - 偶数の場合は、（統計学的な定義の意味での、浮動小数点数を含む）中央値に近い `A[i]` の間であれば、どこでも最小値は小さくなることも理解できる。
        - 自分は保険を打った無意味な `Min` 関数を挟んでいたが、まあこれはこれでよし。
  - **まずはありのままを図示することを心がけたほうが良いかもしれない。**
- D問題は600点なのでまた強くなってから。。