package pattern

import "errors"

type Stack struct {
	sp   int
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{
		sp:   0,
		data: []interface{}{},
	}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.sp == 0 {
		return nil, errors.New("stack is empty")
	}
	s.sp--
	return s.data[s.sp], nil
}

func (s *Stack) Push(v interface{}) error {
	if v == nil {
		return errors.New("value is invalid")
	}
	s.data = append(s.data, v)
	s.sp++
	return nil
}
