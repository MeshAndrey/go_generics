package gogenerics

type Set[T comparable] struct {
	valueMap map[T]bool
}

func (s *Set[T]) Add(newValue T) {
	if s.valueMap == nil {
		s.valueMap = map[T]bool{}
	}

	s.valueMap[newValue] = true
}

func (s *Set[T]) Remove(removeValue T) {
	if s.valueMap == nil {
		s.valueMap = map[T]bool{}
	}

	delete(s.valueMap, removeValue)
}

func (s *Set[T]) Contains(value T) bool {
	if s.valueMap == nil {
		return false
	}

	_, contains := s.valueMap[value]
	return contains
}

func (s *Set[T]) Intersect(compareSet Set[T]) Set[T] {
	if s.valueMap == nil || compareSet.valueMap == nil {
		return Set[T]{}
	}

	intersect := Set[T]{}

	for value := range compareSet.valueMap {
		if s.Contains(value) {
			intersect.Add(value)
		}
	}

	return intersect
}

func (s *Set[T]) Union(anotherSet Set[T]) Set[T] {
	if s.valueMap == nil || anotherSet.valueMap == nil {
		return Set[T]{}
	}

	union := Set[T]{}

	for value := range s.valueMap {
		union.Add(value)
	}

	for value := range anotherSet.valueMap {
		union.Add(value)
	}

	return union
}
