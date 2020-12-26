# ABC185 感想

33分49秒で全完してパフォ2400のカンストだった。  
狂ってる。

## A問題

`min` 関数。

## B問題

シミュレーションを丁寧にやる。

## C問題

パスカルの三角形をオーバーフローを気にせずにやればよい。  
多倍長整数であれば普通に組合せの公式で計算すればよいはず。

※解答ページを見ると、どうやら言語によっては、関係ない場所でのオーバーフローもプログラムの動作保証をしなくなる事があるらしい。  
Goはどうなんだろうか。。？

## D問題

両端に番兵を仕込んでから連続する白マスの数を配列として新たに作る。  
その中で最小の値が `k` として選ぶべき値になる。  
0除算しないように注意が必要。

テストするまで0除算の可能性や、もとの配列がソート済みでないことに気づかず、サンプルを通すのに時間がかかった。

今回はここで潜伏を解除して全て提出した。

## E問題

最初んーと考えたが、よく考えたらレーベンシュタイン距離そのものだった。  
そしてなぜか奇跡的に、3時間ほど前にけんちょん本の編集距離のDPをライブラリにまとめていたため、それを貼って型を `int` にするだけだった。

Fが終わってるほどに簡単だったので、これが命運を分けた。

## F問題

区間 `XOR` 、点 `XOR` 更新のセグメントツリーを使うだけ。  

Fにこんなの置いちゃダメだと思う。
