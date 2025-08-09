package gogenerics

import "cmp"

func QuickSortAscending[T cmp.Ordered](s []T, belowIndex int, upperIndex int) []T {
	if belowIndex < upperIndex {
		part := dividePartsAcsending(s, belowIndex, upperIndex)
		QuickSortAscending(s, belowIndex, part-1)
		QuickSortAscending(s, part+1, upperIndex)
	}

	return s
}

func dividePartsAcsending[T cmp.Ordered](s []T, belowIndex int, upperIndex int) int {
	center := s[upperIndex]
	i := belowIndex

	for j := belowIndex; j < upperIndex; j++ {
		if s[j] <= center {
			s[i], s[j] = s[j], s[i]
			i += 1
		}
	}

	s[i], s[upperIndex] = s[upperIndex], s[i]

	return i
}

func QuickSortDescending[T cmp.Ordered](s []T, belowIndex int, upperIndex int) []T {
	if belowIndex < upperIndex {
		part := dividePartsDescending(s, belowIndex, upperIndex)
		QuickSortDescending(s, belowIndex, part-1)
		QuickSortDescending(s, part+1, upperIndex)
	}

	return s
}

func dividePartsDescending[T cmp.Ordered](s []T, belowIndex int, upperIndex int) int {
	center := s[upperIndex]
	i := belowIndex

	for j := belowIndex; j < upperIndex; j++ {
		if s[j] >= center {
			s[i], s[j] = s[j], s[i]
			i += 1
		}
	}

	s[i], s[upperIndex] = s[upperIndex], s[i]

	return i
}
