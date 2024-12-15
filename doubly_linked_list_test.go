package gogenerics

import (
	"slices"
	"testing"
)

func TestDoublyLinkedListPushBack(t *testing.T) {
	ll := DoublyLinkedList[int]{}

	ll.PushBack(&DoublyLinkedListNode[int]{
		Value: 0,
	})

	ll.PushBack(&DoublyLinkedListNode[int]{
		Value: 1,
	})

	ll.PushBack(&DoublyLinkedListNode[int]{
		Value: 2,
	})

	values := []int{}

	ll.IterateFront(func(node *DoublyLinkedListNode[int]) {
		if node == nil {
			return
		}

		values = append(values, node.Value)
	})

	if slices.Compare(values, []int{0, 1, 2}) != 0 {
		t.Errorf("slices.Compare(values, []int{0, 1, 2}) != 0")
	}
}

func TestDoublyLinkedListPushFront(t *testing.T) {
	ll := DoublyLinkedList[int]{}

	ll.PushFront(&DoublyLinkedListNode[int]{
		Value: 0,
	})

	ll.PushFront(&DoublyLinkedListNode[int]{
		Value: 1,
	})

	ll.PushFront(&DoublyLinkedListNode[int]{
		Value: 2,
	})

	values := []int{}

	ll.IterateFront(func(node *DoublyLinkedListNode[int]) {
		if node == nil {
			return
		}

		values = append(values, node.Value)
	})

	if slices.Compare(values, []int{2, 1, 0}) != 0 {
		t.Errorf("slices.Compare(values, []int{2, 1, 0}) != 0")
	}
}

func TestDoublyLinkedListRemoveNode(t *testing.T) {
	ll := DoublyLinkedList[int]{}

	ll.PushBack(&DoublyLinkedListNode[int]{Value: 0})

	ll.PushBack(&DoublyLinkedListNode[int]{Value: 1})

	tail := DoublyLinkedListNode[int]{Value: 2}

	ll.PushBack(&tail)

	ll.RemoveNode(&tail)

	values := []int{}

	ll.IterateFront(func(node *DoublyLinkedListNode[int]) {
		if node == nil {
			return
		}

		values = append(values, node.Value)
	})

	if slices.Compare(values, []int{0, 1}) != 0 {
		t.Errorf("slices.Compare(values, []int{0, 1}) != 0")
	}
}
