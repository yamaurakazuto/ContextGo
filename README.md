context深掘り
[Background]  ← トップルート（不滅）
      │
      └── ctxA (WithCancel) ← cancel() で止まる
             │
             └── 子goroutineたちも止まる

      └── ctxB (WithTimeout 10s) ← 10秒後に止まる
             │
             └── その子達も止まる


zenn
https://zenn.dev/yamakazu/articles/681fc727645ef4