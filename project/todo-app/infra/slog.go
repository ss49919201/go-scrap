package infra

import (
	"context"
	"os"
	"sync"

	"golang.org/x/exp/slog"
)

func main() {}

var defaultLogger = slog.New(
	slog.NewJSONHandler(os.Stdout),
)

type logger struct {
	level slog.Level
	attrs []slog.Attr
	sync.Mutex
}

func Info() *logger {
	return &logger{
		level: slog.LevelInfo,
		attrs: make([]slog.Attr, 0),
	}
}

func (l *logger) Str(k string, v string) *logger {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.attrs = append(l.attrs, slog.String(k, v))
	return l
}
func (l *logger) Int(k string, v int) *logger {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.attrs = append(l.attrs, slog.Int(k, v))
	return l
}

func (l *logger) Any(k string, v any) *logger {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	l.attrs = append(l.attrs, slog.Any(k, v))
	return l
}

func (l *logger) Msg(ctx context.Context, msg string) {
	defaultLogger.LogAttrs(ctx, l.level, msg, l.attrs...)
}

func (l *logger) Send(ctx context.Context) {
	defaultLogger.LogAttrs(ctx, l.level, "", l.attrs...)
}
