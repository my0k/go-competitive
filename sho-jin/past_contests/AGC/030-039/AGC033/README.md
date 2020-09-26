# AGC033 過去問感想

Last Change: 2020-02-14 01:24:42.

※コンテストも参加していたが、こちらにメモしていく。

## B問題（@2020-02-14）

実はこれはDPだった。。

- 自明なDPから発想をスタートすると良い。
  - `O(HWN)`: `dp[i][j][k]` を `k番目に後手が (i, j) マスにいた時に勝てるかどうかをbool値でもつ` というふうに、まずは愚直に考える。
    - 確かにこの起点は重要。

**典型: 後ろから考える**

やはりゲーム系は後ろから考えるのが特に重要。。

逆算部分はちゃんと真面目に先手・後手・先手・・・でそれぞれの手番を最適に考える必要がある。
最初は、先手しか最適に動かすようなコードを書いてしまったため、WAになってしまう。

このあたりは[けんちょんさんの解説](http://drken1215.hatenablog.com/entry/2019/05/07/000200)が素晴らしくわかりやすい。
半開区間っぽく考えている点は注意（自分は閉区間とみなしてコードを書いた）。
