package main

import (
	"context"
	"time"
)

var _ context.Context = (*valutCtx)(nil)

type valutCtx struct {
	parentCtx  context.Context // 標準ライブラリの実装は埋め込み
	key, value any
}

func (v *valutCtx) Deadline() (deadline time.Time, ok bool) {
	return v.Deadline()
}

func (v *valutCtx) Done() <-chan struct{} {
	return v.Done()
}

func (v *valutCtx) Err() error {
	return v.Err()
}

func (v *valutCtx) Value(key any) any {
	// TODO: 未実装
	return nil
}
