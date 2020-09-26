# ARC103 過去問感想

Last Change: 2020-03-17 00:32:21.

## E問題（@2020-03-16）

木グラフの構築問題。
わかりそうで全然わからなかったので、解説PDFを読んでACした。

スターグラフを起点にして考えるとかなり厳しいことになる。
構築では極端なグラフから変形していくというのはありがちな気がするが、
当然ながらそれだけでは解けないので注意したい。

解説PDFにある、以下の一文が全てかもしれない（構築アルゴリズムの証明はちょっとむずかしいので飛ばしても良い、と思う）。

> 直感的には、存在してはいけないサイズの部分木が存在しないように、うまく **根付き木** を構成します。

「根付き木」の部分は単に「部分木」と読み替えたほうが良いような気もする。

**少し思ったこと: 木は再帰的な構造故に、構築方法も再帰的に考えられないか？と考えてみるのはありかもしれない。**
