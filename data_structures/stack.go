package gogenerics

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(value T) {
	if s == nil {
		return
	}

	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() {
	if s == nil {
		return
	}

	if s.IsEmpty() {
		return
	}

	l := s.Length()

	s.data = s.data[:l-1]
}

func (s Stack[T]) Length() int {
	return len(s.data)
}

func (s Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s Stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}

	return s.data[s.Length()-1], nil
}

func (s Stack[T]) Print() {
	for _, v := range s.data {
		fmt.Println(v)
	}
}
