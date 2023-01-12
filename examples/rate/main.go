package main

import (
	"context"
	"math"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()

	// 1秒間に0.5回(2秒間に1回)の処理を許可する
	l := rate.NewLimiter(0.5, 1)
	for i := 0; i < 10; i++ {
		if err := l.Wait(ctx); err != nil {
			panic(err)
		}
		println(i)
	}

	for i := 0; i < 10; i++ {
		if l.Allow() {
			println(i)
		} else {
			// 許可されなかったら少し待つ
			println("not allowed")
			time.Sleep(time.Second)
		}
	}
}

// 処理の最大頻度
// 回/秒
// 0だと実行なし
type limit float64

// 限界値
const infinity = limit(math.MaxFloat64)

// 1秒間の間の処理間隔を受け取ってlimitに変換する
// e.g every(1*time.Second) -> 1, every(2*time.Second) -> 0.5
func every(interval time.Duration) limit {
	if interval <= 0 {
		return infinity
	}
	return 1 / limit(interval.Seconds())
}

type limitter struct {
	limit     limit
	burst     int       // 瞬間的な処理数の増大,バケツのサイズ
	tokens    float64   // 現在のトークン数
	last      time.Time // tokensが更新された時間
	lastEvent time.Time //
}

type reservation struct {
	ok        bool
	limiter   *limitter
	tokens    int
	timeToAct time.Time
	lim       limit
}

func (l *limitter) allow() bool {
	// reserve
	var (
		now            = time.Now()
		n              = 1
		newReservation reservation
	)
	if l.limit == infinity {
		newReservation = reservation{
			ok:        true,
			limiter:   l,
			tokens:    0,
			timeToAct: time.Now(),
		}
	}

	// limitterの更新時刻が未来になっていたら更新
	last := l.last
	if now.Before(last) {
		last = now
	}

	// e.g tokens=1, limit=0.5, burst=2
	// 余っているバケットの容量を求める
	// e.g 2 - 1 = 1
	reminingBucketSize := float64(l.burst) - l.tokens
	// 余っているバケットの容量から追加できるトークン数を求める
	// rで割ってトークンを追加できる最大数を求める
	// トークンはr秒で1個増える為、rで割ると何個増えるかがわかる
	// e.g 1s / 0.5s = 2
	addableTokensAmountFromReminingBucketSize := reminingBucketSize / float64(l.limit)
	// ns単位にする
	// e.g 2s * 1e9ns = 2e9ns
	addableTokensAmountFromReminingBucketSizeNS := time.Duration(addableTokensAmountFromReminingBucketSize*1e9) * time.Nanosecond
	// 経過時刻から追加できるトークン数を求める
	// e.g 3e9ns
	addableTokensAmountFromElapseNS := now.Sub(last)
	// 追加したいトークン数が追加できるトークン数を超えていたら、追加できるトークン数にする
	// e.g 3e9ns > 2e9ns -> true
	if addableTokensAmountFromElapseNS > addableTokensAmountFromReminingBucketSizeNS {
		// e.g 2e9ns
		addableTokensAmountFromElapseNS = addableTokensAmountFromReminingBucketSizeNS
	}
	// 追加できるトークン数を秒に戻す
	// 少数も含める
	// e.g 2e9ns * 0.5s = 1
	addableTokensAmount := float64(addableTokensAmountFromElapseNS/1e9) * float64(l.limit)
	addableTokensAmount += float64(addableTokensAmountFromElapseNS%1e9) * float64(l.limit)
	// 超えていないかチェック
	// 秒より細かい単位で加算している為
	if burst := l.burst; addableTokensAmount > float64(burst) {
		addableTokensAmount = float64(burst)
	}

	// 追加できる数から追加したい数を減算
	addableTokensAmount -= float64(n)
	// 追加できる数が0未満になれば待ち時間が生じる
	var waitDuration time.Duration
	if addableTokensAmount < 0 {
		waitDuration = time.Duration(-addableTokensAmount*1e9) * time.Nanosecond
	}

	// 追加したい数がバケツの最大サイズ以下か
	withinMaxBucketSize := n <= l.burst
	// 待ち時間が発生していないか
	notWait := waitDuration <= 0

	return withinMaxBucketSize && notWait

	ok := withinMaxBucketSize && notWait

	newReservation = reservation{
		ok:      ok,
		limiter: l,
		lim:     l.limit,
	}
	if ok {
		newReservation.tokens = n
		newReservation.timeToAct = now.Add(waitDuration)

		l.last = now
		l.tokens = addableTokensAmount
		l.lastEvent = newReservation.timeToAct
	} else {
		l.last = last
	}

	return newReservation.ok
}

// time.Durationはns単位
