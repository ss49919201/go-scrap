package tw

import (
	"bytes"
	"text/tabwriter"
)

type DefaultTabWriter struct {
	buf    *bytes.Buffer
	writer *tabwriter.Writer
}

func Default() *DefaultTabWriter {
	buf := bytes.NewBuffer([]byte{})
	return &DefaultTabWriter{
		buf:    buf,
		writer: tabwriter.NewWriter(buf, 2, 2, 0, ' ', tabwriter.Debug),
	}
}

func (tw *DefaultTabWriter) WriteString(s string) {
	tw.writer.Write([]byte(s + "\t"))
}

func (tw *DefaultTabWriter) FulshAndString() (string, error) {
	if err := tw.writer.Flush(); err != nil {
		return "", err
	}

	return tw.buf.String(), nil
}
