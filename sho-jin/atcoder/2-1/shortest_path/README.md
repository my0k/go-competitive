# 幅優先探索

キューを使うだけで簡単そうだが、以下の点には注意。

- マーキングのタイミング
    - タイミングを誤ると、同じ場所を何回も踏んでしまい、TLEしてしまう。
- マーキングする値
    - よく考えずにマークする値を設定すると、誤答してしまう。
- 「最短路」を求める問題の場合、Dijkstra法のようによりよい経路の場合が見つかったら、値を更新するロジックも必要になる。
- 「最短路」だけでなく「最短手順」を求める問題にも応用できる可能性がある。
- 未踏の場所のフラグは `-1` よりも `infinity` で初期化したほうがロジックが書きやすくなる（気がする）。