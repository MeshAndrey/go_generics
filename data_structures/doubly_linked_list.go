package gogenerics

import "fmt"

type DoublyLinkedListNode[T any] struct {
	Value    T
	Previous *DoublyLinkedListNode[T]
	Next     *DoublyLinkedListNode[T]
}

// simple doubly-linked list. does not contain length counter
type DoublyLinkedList[T any] struct {
	Head *DoublyLinkedListNode[T]
	Tail *DoublyLinkedListNode[T]
}

func (s *DoublyLinkedList[T]) PushBack(newNode *DoublyLinkedListNode[T]) {
	if newNode == nil {
		return
	}

	if s.Tail == nil && s.Head == nil {
		s.Head = newNode
		s.Tail = newNode
		return
	}

	s.Tail.Next = newNode
	s.Tail.Next.Previous = s.Tail
	s.Tail = newNode
}

func (s *DoublyLinkedList[T]) PushFront(newNode *DoublyLinkedListNode[T]) {
	if newNode == nil {
		return
	}

	if s.Head == nil && s.Tail == nil {
		s.Head = newNode
		s.Tail = newNode
		return
	}

	newNode.Next = s.Head
	s.Head = newNode
}

// internal iterator pattern.
func (s *DoublyLinkedList[T]) IterateFront(handler func(node *DoublyLinkedListNode[T])) {
	p := s.Head // temporary pointer

	for p != nil {
		handler(p)
		p = p.Next
	}
}

// internal iterator pattern
func (s *DoublyLinkedList[T]) IterateForward(handler func(node *DoublyLinkedListNode[T])) {
	p := s.Tail // temporary pointer

	for p != nil {
		handler(p)
		p = p.Previous
	}
}

func (s *DoublyLinkedList[T]) PrintFront() {
	s.IterateFront(func(node *DoublyLinkedListNode[T]) {
		fmt.Print(node.Value, " -> ")
	})
}

func (s *DoublyLinkedList[T]) PrintForward() {
	s.IterateForward(func(node *DoublyLinkedListNode[T]) {
		fmt.Print(node.Value, " -> ")
	})
}

func (s *DoublyLinkedList[T]) RemoveNode(deleteNode *DoublyLinkedListNode[T]) {
	if deleteNode == nil {
		return
	}

	p := s.Head

	for p.Next != nil {
		if p.Next == deleteNode { // 1, 2, 4
			if p.Next.Next == nil {
				p.Next = nil
			} else {
				p.Next = p.Next.Next
				p.Next.Previous = p
			}

			return
		}

		p = p.Next
	}
}
