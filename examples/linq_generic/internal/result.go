package internal

func (q Query[T]) ToSlice() []T {
	var slice []T
	next := q.Iterate()
	for {
		elm, ok := next()
		if !ok {
			break
		}

		slice = append(slice, elm)
	}
	return slice
}
