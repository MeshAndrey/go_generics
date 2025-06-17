package gogenerics

import "testing"

func TestSetAdd(t *testing.T) {
	s := Set[int]{}

	s.Add(0)
	s.Add(1)
	s.Add(2)

	if !s.Contains(0) || !s.Contains(1) || !s.Contains(2) {
		t.Errorf("!s.Contains(0) || !s.Contains(1) || !s.Contains(2)")
	}
}

func TestSetRemove(t *testing.T) {
	s := Set[int]{}

	s.Add(0)
	s.Add(1)
	s.Add(2)

	s.Remove(0)
	s.Remove(1)
	s.Remove(2)

	if s.Contains(0) || s.Contains(1) || s.Contains(2) {
		t.Errorf("s.Contains(0) || s.Contains(1) || s.Contains(2)")
	}
}

func TestSetContains(t *testing.T) {
	s := Set[int]{}

	s.Add(0)

	if !s.Contains(0) {
		t.Errorf("!s.Contains(0)")
	}
}

func TestSetIntersect(t *testing.T) {
	s := Set[int]{}

	s.Add(0)
	s.Add(1)
	s.Add(2)

	s1 := Set[int]{}

	s1.Add(0)
	s1.Add(1)

	i := s.Intersect(s1)

	if !i.Contains(0) || !i.Contains(1) {
		t.Errorf("!i.Contains(0) || !i.Contains(1)")
	}
}

func TestSetUnion(t *testing.T) {
	s := Set[int]{}

	s.Add(0)
	s.Add(1)
	s.Add(2)

	s1 := Set[int]{}

	s1.Add(3)
	s1.Add(4)

	u := s.Union(s1)

	if !u.Contains(0) || !u.Contains(1) || !u.Contains(2) || !u.Contains(3) || !u.Contains(4) {
		t.Errorf("!u.Contains(0) || !u.Contains(1) || !u.Contains(2) || !u.Contains(3) || !u.Contains(4)")
	}
}
