# ACLC 1st 感想

A問題もっと簡単にしてください。2回目の太陽。

## A問題

平面走査的な発想はすぐに出てきたが、中途半端に研究でスカイラインクエリの話とかをかじってたせいで迷走してしまった。

**まず連結成分サイズを求めることからぶれてはいけなかった。**

そしてエッジを貼る数を減らさないといけないわけだが、自分はこれを区間で管理しようと頑張ってしまった。
実際は、連結成分のつなぎ先を1つだけでも管理しておけばいいわけで、この問題ではなんとスタックだけで効率的に済ませてしまえる。
