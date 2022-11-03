package internal

func (q Query[T]) ToSlice() []T {
	var slice []T
	next := q.Iterate()
	for {
		elm, ok := next()
		if ok {
			slice = append(slice, elm)
		} else {
			break
		}
	}
	return slice
}
