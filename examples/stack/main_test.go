package pattern

import (
	"reflect"
	"testing"
)

type testTbl[T any] struct {
	name string
	want *Stack[T]
	err  error
}

func TestStackInt(t *testing.T) {
	test := testTbl[int]{
		"int stack",
		&Stack[int]{
			sp:   0,
			data: []int{},
		},
		nil,
	}

	t.Run(test.name, func(t *testing.T) {
		if got := NewStack[int](); !reflect.DeepEqual(got, test.want) {
			t.Errorf("NewStack() = %v, want %v", got, test.want)
		}
	})
}

func TestStackIntSlice(t *testing.T) {
	test := testTbl[[]int]{
		"int slice stack",
		&Stack[[]int]{
			sp:   0,
			data: [][]int{},
		},
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := NewStack[[]int](); !reflect.DeepEqual(got, test.want) {
			t.Errorf("NewStack() = %v, want %v", got, test.want)
		}
	})
}

func TestStackFunc(t *testing.T) {
	type fn func()
	test := testTbl[fn]{
		"func stack",
		&Stack[fn]{
			sp:   0,
			data: []fn{},
		},
		nil,
	}
	t.Run(test.name, func(t *testing.T) {
		if got := NewStack[fn](); !reflect.DeepEqual(got, test.want) {
			t.Errorf("NewStack() = %v, want %v", got, test.want)
		}
	})

	funcElm := func() {}

	test2 := testTbl[fn]{
		"push func stack",
		&Stack[fn]{
			sp: 1,
			data: []fn{
				funcElm,
			},
		},
		nil,
	}

	t.Run(test2.name, func(t *testing.T) {
		s := NewStack[fn]()
		s.Push(funcElm)
		if s.sp != test2.want.sp {
			t.Errorf("actual %v, want %v", s.sp, test2.want.sp)
		}
		// 関数の等値比較はポインタ値で見るしかない
		if reflect.ValueOf(s.data[0]).Pointer() != reflect.ValueOf(test2.want.data[0]).Pointer() {
			t.Errorf("actual %v, want %v", reflect.ValueOf(s.data[0]).Pointer(), reflect.ValueOf(test2.want.data[0]).Pointer())
		}
	})

	test3 := testTbl[fn]{
		"failed push func stack",
		&Stack[fn]{
			sp:   0,
			data: []fn{},
		},
		nil,
	}

	t.Run(test3.name, func(t *testing.T) {
		s := NewStack[fn]()
		if err := s.Push(nil); err == nil {
			t.Error("want error")
		}
		if !reflect.DeepEqual(s, test3.want) {
			t.Errorf("actual %v, want %v", s.sp, test3.want.sp)
		}
	})
}
