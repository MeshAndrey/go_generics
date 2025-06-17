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

func TestQueueRemove(t *testing.T) {
	q := Queue[int]{}

	q.Push(10)
	q.Push(20)
	q.Push(30)

	var (
		q1 *int
		q2 *int
		q3 *int
	)

	q.IterateFront(func(t *int) {
		if *t == 20 {
			q2 = t
		}
	})

	if q2 == nil {
		t.Errorf("q2 == nil")
	}

	q.Remove(q2)

	q.IterateFront(func(t *int) {
		if *t == 10 {
			q1 = t
		}
	})

	if q1 == nil {
		t.Errorf("q1 == nil")
	}

	q.Remove(q1)

	q.IterateFront(func(t *int) {
		if *t == 30 {
			q3 = t
		}
	})

	if q3 == nil {
		t.Errorf("q3 == nil")
	}

	q.Remove(q3)

	if q.Length() != 0 {
		t.Errorf("q.Length() != 0")
	}
}

func TestIterateFront(t *testing.T) {
	q := Queue[int]{}

	q.Push(10)
	q.Push(20)
	q.Push(30)

	q1, _ := q.Head()
	q3, _ := q.Tail()

	q1Found := false
	q.IterateFront(func(t *int) {
		if *t == q1 {
			q1Found = true
		}
	})

	if !q1Found {
		t.Errorf("q1Found == false")
	}

	q3Found := false
	q.IterateFront(func(t *int) {
		if *t == q3 {
			q3Found = true
		}
	})

	if !q3Found {
		t.Errorf("q3Found == false")
	}
}
