package gogenerics

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	data []T
}

func (s *Queue[T]) Push(value T) {
	if s == nil {
		return
	}

	s.data = append(s.data, value)
}

func (s *Queue[T]) Pop() {
	if s == nil {
		return
	}

	if s.IsEmpty() {
		return
	}

	l := s.Length()

	if l == 1 {
		s.data = []T{}
		return
	}

	s.data = s.data[1:]
}

func (s Queue[T]) Length() int {
	return len(s.data)
}

func (s Queue[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s Queue[T]) Head() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Queue is empty")
	}

	return s.data[0], nil
}

func (s Queue[T]) Tail() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("Queue is empty")
	}

	return s.data[s.Length()-1], nil
}

func (s Queue[T]) Print() {
	for _, v := range s.data {
		fmt.Println(v)
	}
}
