# DDCC 2019 本戦 過去問感想

- A問題は300点だが、若干実装が大変。というか考察も少し大変な気がする。
  - 制約がいろいろと優しくて、面倒な部分が考えなくて良いようになっている。
    - どの雪のマス目も2マス以上連続している: 氷に変えても、前後の氷とつながることはない。
    - 最初と最後は必ず雪のマス: 氷に変えるべきところを調べる際に、余分な例外処理が不要になる。
  - サンプルからも、「できるだけ氷が最初から連続しているところを探し、それを更に延長するようにマス目を変更すればよい」というのは何となく分かる。
  - 実装は、氷の最大と変えるべきマス目のチェックを一度にやるのはバグの原因になると思ったため、別々の処理に分けたが、TLEしてしまった。
    - `[]rune -> string` の変換は軽くはないらしい（それはそうだった）。
  - 最大の連続氷マスが更新されるたびに、その次のマス目を保持しておくだけでよかった。

