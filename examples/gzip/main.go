package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

// decompress command: `gzip -d tmp.gz`
func main() {
	pauloadBuf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(pauloadBuf).Encode(map[string]any{
		"s": "string",
		"m": map[string]any{
			"nested": 1e2,
		},
	}); err != nil {
		panic(err)
	}

	compressedBuf := bytes.NewBuffer(nil)
	if err := compressGzip(compressedBuf, pauloadBuf.Bytes()); err != nil {
		panic(err)
	}

	f, err := os.Create("tmp.gz")
	if err != nil {
		panic(err)
	}

	if _, err := f.Write(compressedBuf.Bytes()); err != nil {
		panic(err)
	}
}

func compressGzip(w io.Writer, payload []byte) error {
	zw := gzip.NewWriter(w)

	if _, err := zw.Write(payload); err != nil {
		return err
	}

	// flush後に閉じられる
	// (w *huffmanBitWriter) flush() が内部で呼ばれる
	return zw.Close()
}
