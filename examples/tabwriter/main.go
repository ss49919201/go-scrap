package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

func main() {
	buf := bytes.NewBuffer([]byte{})
	// 初期化
	w := tabwriter.NewWriter(buf, 2, 2, 0, ' ', tabwriter.Debug)
	// minwidthはtabとtabの間の最小幅
	// tabwidthはtabの幅
	// paddingはtabとtabの間に入れるpaddingの数
	// 書き込み
	fmt.Fprintln(w, "\tmike\t")
	fmt.Fprintln(w, "\tbob\tkazu\tken")
	fmt.Fprintln(w, "\tjon\troy")
	// 出力
	w.Flush()
	fmt.Println(buf.String())
	// Output:
	// 	 |mike|
	//   |bob |kazu|ken
	//   |jon |roy

	// 既にインデント済みの文字を入れたければ、文字列に変換してから書き込む
	// 初期化
	w2 := tabwriter.NewWriter(os.Stdout, 2, 2, 0, ' ', tabwriter.Debug)
	// minwidthはtabとtabの間の最小幅
	// tabwidthはtabの幅
	// paddingはtabとtabの間に入れるpaddingの数
	// 書き込み
	fmt.Fprintln(w2, "\t樋口一葉\t")
	fmt.Fprintln(w2, "\t野口英世\t夏目漱石\t聖徳太子")
	fmt.Fprintln(w2, "\t徳川家康\t")
	fmt.Fprintln(w2, buf.String())
	// 出力
	w2.Flush()
	// Output:
	// 	 |樋口一葉|
	//   |野口英世|夏目漱石|聖徳太子
	//   |徳川家康|
	//   |mike|
	//   |bob |kazu|ken
	//   |jon |roy

	// いい感じにテーブル作る
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	tw.Write([]byte("ID" + "\t"))
	tw.Write([]byte("FIELD_1" + "\t"))
	tw.Write([]byte("FIELD_2" + "\t"))
	tw.Write([]byte("FIELD_3" + "\n"))
	for _, v := range []struct {
		ID        string
		Count     int
		Title     string
		CreatedAt time.Time
	}{
		{"1", 1, "title1", time.Now()},
	} {
		tw.Write([]byte(v.ID + "\t"))
		tw.Write([]byte(strconv.Itoa(v.Count) + "\t"))
		tw.Write([]byte(v.Title + "\t"))
		tw.Write([]byte(v.CreatedAt.Format(time.RFC3339) + "\n"))
	}
	tw.Flush()
	// Output:
	// ID  FIELD_1  FIELD_2  FIELD_3
	// 1   1        title1   2022-10-22T15:15:52+09:00
}
