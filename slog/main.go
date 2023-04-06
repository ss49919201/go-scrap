package main

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	_Info(context.Background(), "hoge", attr{
		"string": "str",
		"int":    1,
		"struct": struct{ I int }{1}, // only public field
	})
}

var defaultLogger = slog.New(
	slog.NewJSONHandler(os.Stdout),
)

type attr map[string]any

func _Info(ctx context.Context, msg string, attr attr) {
	attrs := make([]slog.Attr, 0, len(attr))
	for k, v := range attr {
		attrs = append(attrs, slog.Any(k, v))
	}
	defaultLogger.LogAttrs(ctx, slog.LevelInfo, msg, attrs...)
}
