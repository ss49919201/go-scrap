package main

type set[T comparable] struct {
	m map[T]bool
}

func newSet[T comparable]() *set[T] {
	return &set[T]{m: make(map[T]bool)}
}

func (s *set[T]) add(x T) {
	s.m[x] = true
}

func (s *set[T]) delete(x T) {
	delete(s.m, x)
}

func (s *set[T]) has(x T) bool {
	_, ok := s.m[x]
	return ok
}

func (s *set[T]) len() int {
	return len(s.m)
}

func main() {

}
