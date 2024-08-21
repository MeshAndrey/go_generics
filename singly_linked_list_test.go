package gogenerics

import (
	"slices"
	"testing"
)

func TestSinglyLinkedListAppend(t *testing.T) {
	ll := SinglyLinkedList[int]{
		Head: &SinglyLinkedListNode[int]{
			Value: 0,
			Next:  nil,
		},
	}

	ll.Append(&SinglyLinkedListNode[int]{
		Value: 1,
		Next:  nil,
	})

	ll.Append(&SinglyLinkedListNode[int]{
		Value: 2,
		Next:  nil,
	})

	values := []int{}

	ll.IterateHandle(func(node *SinglyLinkedListNode[int]) {
		if node == nil {
			return
		}

		values = append(values, node.Value)
	})

	if slices.Compare(values, []int{0, 1, 2}) != 0 {
		t.Errorf("slices.Compare(values, []int{0, 1, 2}) != 0")
	}
}

func TestSinglyLinkedListRemoveNode(t *testing.T) {
	ll := SinglyLinkedList[int]{
		Head: &SinglyLinkedListNode[int]{
			Value: 0,
			Next:  nil,
		},
	}

	ll.Append(&SinglyLinkedListNode[int]{
		Value: 1,
		Next:  nil,
	})

	tail := &SinglyLinkedListNode[int]{
		Value: 2,
		Next:  nil,
	}

	ll.Append(tail)

	ll.RemoveNode(tail)

	values := []int{}

	ll.IterateHandle(func(node *SinglyLinkedListNode[int]) {
		if node == nil {
			return
		}

		values = append(values, node.Value)
	})

	if slices.Compare(values, []int{0, 1}) != 0 {
		t.Errorf("slices.Compare(values, []int{0, 1}) != 0")
	}
}
