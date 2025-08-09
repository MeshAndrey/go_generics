package gogenerics

import (
	"fmt"
	"slices"
	"testing"
)

func TestQuickSortAscending(t *testing.T) {
	values := []int{3, 2, 1, 0}

	orderedValues := QuickSortAscending(values, 0, len(values)-1)

	fmt.Println(orderedValues)

	if slices.Compare(orderedValues, []int{0, 1, 2, 3}) != 0 {
		t.Errorf("slices.Compare(orderedValues, []int{0, 1, 2, 3}) != 0")
		t.FailNow()
	}
}

func TestQuickSortDescending(t *testing.T) {
	values := []int{0, 1, 2, 3}

	orderedValues := QuickSortDescending(values, 0, len(values)-1)

	fmt.Println(orderedValues)

	if slices.Compare(orderedValues, []int{3, 2, 1, 0}) != 0 {
		t.Errorf("slices.Compare(orderedValues, []int{3, 2, 1, 0}) != 0")
		t.FailNow()
	}
}
