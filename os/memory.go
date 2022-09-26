package main

import "time"

type Memory interface {
	MainMemory | SecondoryMemory
}

type MainMemory struct{}

type SecondoryMemory struct{}

type job struct{}

// 優先度の高いジョブを実行する為に、現在のジョブを中断する
func preemptJob(current, target job) {
	stopIO(current)
	// スワップアウト
	// 二次記憶に移動
	moveToMemory(current, SecondoryMemory{})
	// スワップイン
	// メインメモリにロード
	load(target, MainMemory{})
}

// ジョブの入出力動作を停止する
func stopIO(j job)

func moveToMemory[T Memory](j job, m T) {}

func load(j job, m MainMemory) {}

// 時分割システムを実行する
// 一つのジョブの一回の実行に対して一定時間割り当てる
// 高速ランダムアクセス可能な二次記憶を用意する
func ExecTimeDivision(quantum time.Duration) func() {
	cancelCh := make(chan struct{})
	cancel := func() {
		cancelCh <- struct{}{}
	}

	go func() {
		for {
			select {
			case <-time.Tick(quantum):
				switchJob()
			case <-cancelCh:
			}
		}
	}()

	return cancel
}

func switchJob() {}
