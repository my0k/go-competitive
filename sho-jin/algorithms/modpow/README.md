# 剰余、逆元

[けんちょん氏の記事](https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#%E3%81%AA%E3%81%9C-1000000007-%E3%81%AA%E3%81%AE%E3%81%8B)の自分用メモ。
すべてを頭に入れられたわけではないので、現時点で必要と感じた部分だけまとめている。

## 足し算、掛け算

**都度、MOD（法）であまりをとる。**

## 引き算

> **「$a$ を $b$ で割ったあまり」というのは「$a$ に $b$ の倍数を足し引きして得られる数のうち、$0$ 以上 $b$ 未満のもの」。**

検証したところ、GoでもC++と同じように素直に負の数の剰余を計算すると負の数値が返ってくるため、法の数値を足す必要がある。

## 割り算

### `mod p` における「逆元（Inverse Element）」

結論から書いてしまう。

$$a \div b \equiv a \times b^{-1} ({\rm mod}~~p)$$

$b^{-1}$は法 $p$ のもとでの逆元。

割りたい数の逆元がわかれば、

```go
// 誤ったコード
ans /= b
ans %= MOD
```

とやりたいところが、

```go
// 逆元を使って正しい剰余を計算する
binv := modinv(b, MOD)
ans *= binv
ans %= MOD
```

と書くことができる。

### 法 $p$ のもとでの逆元

#### 1. フェルマーの小定理から計算

##### フェルマーの小定理

$p$ を素数、$a$ を $p$ の倍数でない整数として

$$a^{p-1} \equiv 1 ({\rm mod}~~p)$$

が成立する。

##### 変形

フェルマーの小定理の式を変形すると、

$$a \times a^{p-2} \equiv 1 ({\rm mod}~~p)$$

となることから、**$a^{p-2}$ が $a$ の逆元になっていることを意味している。**

よって、**二分累乗法を用いることで $O(\log(p))$ で求めることができる。**

##### 実装

実装については、ABC110のD問題の解説放送で示されていた、[chokudaiさんのC++コード](https://atcoder.jp/contests/abc110/submissions/3260417)を参考にさせていただいた。

#### 2. 拡張ユークリッドの互除法から計算

work in progress...

### 逆元が存在する条件

0除算が許されない、行列式が0の際に逆行列が存在しないのと同じく、逆元も常に存在するとは限らない。

$\mod p$ での $a$ の逆元が存在する条件は、 **$p$ と $a$ とが互いに素であること。**

競技プログラミングで逆元を求めるときは、例えば $p = 1000000000+7$ などの素数に対して行うため、それほど深刻には考えなくてよいが、
$a = 0, p, 2p, 3p, ...$ で問題になってくるため、巨大な組み合わせの数などを求める際には注意が必要になる場合があるとのこと。