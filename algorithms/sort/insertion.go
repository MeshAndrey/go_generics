package gogenerics

import "cmp"

func InsertionSortAscending[T cmp.Ordered](s []T) []T {
	for j := 1; j < len(s); j++ {
		key := s[j]

		// insertion
		i := j - 1

		for i >= 0 && s[i] > key {
			s[i+1] = s[i]
			i = i - 1
		}

		s[i+1] = key
	}

	return s
}

func InsertionSortDescending[T cmp.Ordered](s []T) []T {
	for j := 1; j < len(s); j++ {
		key := s[j]

		// insertion
		i := j - 1

		for i >= 0 && s[i] < key {
			s[i+1] = s[i]
			i = i - 1
		}

		s[i+1] = key
	}

	return s
}
