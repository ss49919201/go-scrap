package internal

type Query[T any] struct {
	Iterate func() func() (T, bool)
}

func From[T any](elms []T) Query[T] {
	return Query[T]{
		Iterate: func() func() (T, bool) {
			index := 0
			fn := func() (T, bool) {
				if index >= len(elms) {
					return empty[T](), false
				}
				defer func() { index++ }()
				return elms[index], true
			}
			return fn
		},
	}
}

func (q Query[T]) Where(predicate func(T) bool) Query[T] {
	return Query[T]{
		Iterate: func() func() (T, bool) {
			next := q.Iterate()
			fn := func() (T, bool) {
				for {
					elm, ok := next()
					if !ok {
						return empty[T](), false
					}

					if predicate(elm) {
						return elm, true
					}
				}
			}
			return fn
		},
	}
}

func empty[T any]() T {
	var v T
	return v
}
