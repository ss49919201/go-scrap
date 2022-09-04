package pattern

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

// return 文に比較演算子
func RetBool() bool {
	t := true
	f := false
	return t == f
}

func Switch(parm int) {
	switch parm {
	// defaultはどこに書いてもいい
	default:
		fmt.Println("default")
	case 1:
		fmt.Println("true")
	// OR条件で書ける
	case 2, 3:
		fmt.Println("false")
	}
}

func CreateFile(name string) {
	// ./ にファイルを作成
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f.Name())

	// 掃除
	defer os.Remove(f.Name())
}

func CreateTempFile(name string) {
	// 第1引数を空文字列にすることで 環境変数TMPDIR にファイルを作成
	f, err := os.CreateTemp("", name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f.Name())

	// 掃除
	defer os.Remove(f.Name())
}

func AppendSlice() {
	// FIXME: capを確保する
	l := []int{1, 2, 3, 4}
	r := []int{1, 2, 3, 4}
	l = append(l, r...)
	fmt.Println(l)
}

func TypeAssertion() {
	var i interface{} = "hoge"
	// interface{}型の値の、underlying な値を利用するためのもの
	fmt.Println(i.(string))
}

type T struct {
	i int
}

func NewT() *T {
	return &T{
		i: 1,
	}
}

func (t *T) Set(int) {
	t.i = 1
}

func (t *T) Get() int {
	return t.i
}

// Channel

type sendCh chan<- string

type receiveCh <-chan string

func ChanIo() {
	c := make(chan string)
	var s sendCh
	var r receiveCh
	s, r = c, c
	//goroutineでないと fatal error: all goroutines are asleep - deadlock!
	go func() { s <- "Send" }()
	defer close(c)
	time.Sleep(time.Second)
	fmt.Println(<-r)
}

// reflect

func IsInt(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Int
}

func IsValid(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.IsValid()
}

func IsNil(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.IsNil()
}
