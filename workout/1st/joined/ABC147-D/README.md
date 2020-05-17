# ABC147 D. Xor Sum 4

Last Change: 2020-05-15 15:47:16.

## @2020-05-15

1回TLEしてしまった。

コンテスト中はオーバーフローに気づくのに莫大なペナルティと時間がかかってしまったが、
それについては警戒できたためWAはなかった。

今回は、 `2^j` を求めるときに `modpow` を使ってしまったのがTLEの原因。

**`modpow` は定数倍が非常に大きいことに留意する必要がある！（ある数値についての `1000000007-2` の累乗数を求める必要があるため）**

定期的に思い出しておきたい。
