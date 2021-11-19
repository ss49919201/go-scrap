package pattern

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"io"
)

func GzipCSV(w io.Writer, records [][]string) error {
	buf := bytes.NewBuffer(nil)
	if err := writeCSV(buf, records); err != nil {
		return err
	}

	if err := compressGzip(w, buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func writeCSV(w io.Writer, records [][]string) error {
	cw := csv.NewWriter(w)

	// 内部でFlushが呼ばれる
	if err := cw.WriteAll(records); err != nil {
		return err
	}
	return nil
}

func compressGzip(w io.Writer, payload []byte) error {
	zw := gzip.NewWriter(w)

	if _, err := zw.Write(payload); err != nil {
		return err
	}

	if err := zw.Flush(); err != nil {
		return err
	}

	if err := zw.Close(); err != nil {
		return err
	}
	return nil
}
