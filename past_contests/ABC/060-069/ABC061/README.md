# ABC061過去問感想

- A問題、B問題は特になし
- C問題は考え方はあっていたが、新しくテンプレートに記述した標準入力の遅さがネックになってTLEしてしまった
    - 場合に応じてテンプレートのコメントアウトを切り替える（アホな）運用をしていくことにする
    - 自分は辞書を使ったが、**キーが小さい範囲の整数のときは、配列を確保**してしまったほうが速度的にもコードの読みやすさ的にもいいと思われる
    - `n <= 100000` における `O(nlogn)` は100msec程度に収まっているので大丈夫
- D問題
    - 一工夫してベルマンフォード法に帰着する問題
        - 一工夫する前から負のエッジが存在し、始点と終点が固定されていることからも、ベルマンフォード法を用いるという発想にたどり着くのはそれほど難しくなさそう
    - ACするためには、終点が負の閉路に含まれているかどうかまで導かないといけない
        - 単に負の経路を検出するだけでは不十分
            - 負の経路をたどった場合に終点にたどり着けるかどうかはわからないため
            - 例えば、有向グラフのため負の経路に閉じ込められて、終点に向かうことができるケースなどは想定できる
    - ともかくベルマンフォード法の要点を理解するのに良い問題