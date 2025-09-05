package gogenerics

import (
	"errors"
	"fmt"
	"sync"
)

type Queue[T any] struct {
	m    sync.Mutex
	data []T
}

func (s *Queue[T]) Push(value T) {
	if s == nil {
		return
	}
	s.m.Lock()
	defer s.m.Unlock()

	s.data = append(s.data, value)
}

func (s *Queue[T]) Pop() {
	if s == nil {
		return
	}
	s.m.Lock()
	defer s.m.Unlock()

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

func (s *Queue[T]) Head() (T, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.IsEmpty() {
		return *new(T), errors.New("Queue is empty")
	}

	return s.data[0], nil
}

func (s *Queue[T]) Tail() (T, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.IsEmpty() {
		return *new(T), errors.New("Queue is empty")
	}

	return s.data[s.Length()-1], nil
}

func (s *Queue[T]) Print() {
	s.m.Lock()
	defer s.m.Unlock()

	for _, v := range s.data {
		fmt.Println(v)
	}
}

func (s *Queue[T]) Remove(t *T) {
	s.m.Lock()
	defer s.m.Unlock()

	if t == nil {
		return
	}

	if s.Length() == 0 {
		return
	}

	newSlice := []T{}
	found := false
	for i := 0; i < len(s.data); i++ {
		if &s.data[i] != t {
			continue
		}

		found = true

		if s.Length() == 1 { // if only one at queue
			newSlice = []T{}
		} else if i == 0 { // if first
			newSlice = append(newSlice, s.data[i+1:]...)
		} else if i == (s.Length() - 1) { // if last
			newSlice = append(newSlice, s.data[:i]...)
		} else { // at the middle ?
			newSlice = append(newSlice, s.data[:i]...)
			newSlice = append(newSlice, s.data[i+1:]...)
		}

		break
	}

	if found {
		s.data = newSlice
	}
}

func (s *Queue[T]) IterateFront(handler func(t *T)) {
	s.m.Lock()
	defer s.m.Unlock()

	for i := 0; i < len(s.data); i++ {
		handler(&s.data[i])
	}
}
