package gogenerics

import "testing"

func TestQueuePush(t *testing.T) {
	q := Queue[int]{}

	q.Push(10)
	q.Push(20)
	q.Push(30)

	if q.Length() != 3 {
		t.Errorf("q.Length() != 3")
	}
}

func TestQueuePop(t *testing.T) {
	q := Queue[int]{}

	q.Push(10)
	q.Push(20)
	q.Push(30)

	if q.Length() != 3 {
		t.Errorf("q.Length() != 3")
	}

	q.Pop()
	q.Pop()
	q.Pop()

	if q.Length() != 0 {
		t.Errorf("q.Length() != 0")
	}
}

func TestQueueHead(t *testing.T) {
	q := Queue[int]{}

	q.Push(10)
	head, _ := q.Head()
	if head != 10 {
		t.Errorf("head != 10")
	}

	q.Push(20)
	head, _ = q.Head()
	if head != 10 {
		t.Errorf("head != 10")
	}

	q.Push(30)
	head, _ = q.Head()
	if head != 10 {
		t.Errorf("head != 10")
	}

	q.Pop()
	head, _ = q.Head()
	if head != 20 {
		t.Errorf("head != 20")
	}
}

func TestQueueTail(t *testing.T) {
	q := Queue[int]{}

	q.Push(10)
	tail, _ := q.Tail()
	if tail != 10 {
		t.Errorf("tail != 10")
	}

	q.Push(20)
	tail, _ = q.Tail()
	if tail != 20 {
		t.Errorf("tail != 20")
	}

	q.Push(30)
	tail, _ = q.Tail()
	if tail != 30 {
		t.Errorf("tail != 30")
	}

	q.Pop()
	tail, _ = q.Tail()
	if tail != 30 {
		t.Errorf("tail != 30")
	}
}
