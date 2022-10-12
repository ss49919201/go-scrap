package main

// https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
func filterWithoutAllocating[T any](s []T, keep func(v T) bool) []T {
	result := s[:0] // []
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

// https://github.com/golang/go/wiki/SliceTricks#filter-in-place
func filterInPlace[T any](s *[]T, keep func(v T) bool) {
	n := 0
	for _, v := range *s {
		if keep(v) {
			(*s)[n] = v
			n++
		}
	}
	*s = (*s)[:n]
}
