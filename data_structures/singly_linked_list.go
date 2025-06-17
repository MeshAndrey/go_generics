package gogenerics

import "fmt"

type SinglyLinkedListNode[T any] struct {
	Value T
	Next  *SinglyLinkedListNode[T]
}

// simple singly-linked list. does not contain tail pointer and length counter
type SinglyLinkedList[T any] struct {
	Head *SinglyLinkedListNode[T]
}

func (s *SinglyLinkedList[T]) Append(newNode *SinglyLinkedListNode[T]) {
	if newNode == nil {
		return
	}

	p := s.Head // temporary pointer

	for p.Next != nil {
		p = p.Next
	}

	p.Next = newNode
}

// internal iterator pattern
func (s *SinglyLinkedList[T]) IterateHandle(handler func(node *SinglyLinkedListNode[T])) {
	p := s.Head // temporary pointer

	for p != nil {
		handler(p)
		p = p.Next
	}
}

func (s *SinglyLinkedList[T]) Print() {
	s.IterateHandle(func(node *SinglyLinkedListNode[T]) {
		fmt.Print(node.Value, " -> ")
	})
}

func (s *SinglyLinkedList[T]) RemoveNode(deleteNode *SinglyLinkedListNode[T]) {
	if deleteNode == nil {
		return
	}

	p := s.Head

	for p.Next != nil {
		if p.Next == deleteNode {
			if p.Next.Next == nil {
				p.Next = nil
			} else {
				p.Next = p.Next.Next
			}

			return
		}

		p = p.Next
	}
}
