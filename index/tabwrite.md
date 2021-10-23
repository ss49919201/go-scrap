```
// 初期化
2, 2, 0, ' ', tabwriter.Debug)
// minwidthはtabとtabの間の最小幅
// tabwidthはtabの幅
paddingの数
// 書き込み
fmt.Fprint(w, "\t\n")
fmt.Fprint(w, "\tfuga")
fmt.Fprint(w, "\tpiyo")
// 出力
w.Flush()
```