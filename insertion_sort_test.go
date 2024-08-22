package gogenerics

import (
	"slices"
	"testing"
)

func TestInsertionSortAscending(t *testing.T) {
	values := []int{3, 2, 1, 0}

	orderedValues := InsertionSortAscending(values)

	if slices.Compare(orderedValues, []int{0, 1, 2, 3}) != 0 {
		t.Errorf("slices.Compare(orderedValues, []int{0, 1, 2, 3}) != 0")
		t.FailNow()
	}
}

func TestInsertionSortDescending(t *testing.T) {
	values := []int{0, 1, 2, 3, 4}

	orderedValues := InsertionSortDescending(values)

	if slices.Compare(orderedValues, []int{4, 3, 2, 1, 0}) != 0 {
		t.Errorf("slices.Compare(orderedValues, []int{4, 3, 2, 1, 0}) != 0")
		t.FailNow()
	}
}
