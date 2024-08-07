package gogenerics

import "testing"

func TestPush(t *testing.T) {
	is := Stack[int]{}

	is.Push(10)
	is.Push(20)
	is.Push(30)

	if is.Length() != 3 {
		t.Errorf("error to push ints to stack")
	}

	ss := Stack[string]{}

	ss.Push("a")
	ss.Push("b")
	ss.Push("c")

	if ss.Length() != 3 {
		t.Errorf("error to push strings to stack")
	}

	type testType struct{ a string }

	ts := Stack[testType]{}

	ts.Push(testType{a: "a"})
	ts.Push(testType{a: "b"})
	ts.Push(testType{a: "c"})

	if ts.Length() != 3 {
		t.Errorf("error to push custom struct objects to stack")
	}
}

func TestPop(t *testing.T) {
	is := Stack[int]{}

	is.Push(10)
	is.Push(20)
	is.Push(30)

	if is.Length() != 3 {
		t.Errorf("error to push ints to stack")
	}

	is.Pop()
	is.Pop()
	is.Pop()

	if is.Length() != 0 {
		t.Errorf("error to pop ints from stack")
	}

	ss := Stack[string]{}

	ss.Push("a")
	ss.Push("b")
	ss.Push("c")

	if ss.Length() != 3 {
		t.Errorf("error to push strings to stack")
	}

	ss.Pop()
	ss.Pop()
	ss.Pop()

	if ss.Length() != 0 {
		t.Errorf("error to pop strings from stack")
	}

	type testType struct{ a string }

	ts := Stack[testType]{}

	ts.Push(testType{a: "a"})
	ts.Push(testType{a: "b"})
	ts.Push(testType{a: "c"})

	if ts.Length() != 3 {
		t.Errorf("error to push custom struct objects to stack")
	}

	ts.Pop()
	ts.Pop()
	ts.Pop()

	if ts.Length() != 0 {
		t.Errorf("error to pop custom struct objects from stack")
	}
}

func TestTop(t *testing.T) {
	is := Stack[int]{}

	is.Push(10)
	is.Push(20)
	is.Push(30)

	if is.Length() != 3 {
		t.Errorf("error to push ints to stack")
	}

	it, _ := is.Top()

	if it != 30 {
		t.Errorf("wrong top value")
	}

	ss := Stack[string]{}

	ss.Push("a")
	ss.Push("b")
	ss.Push("c")

	if ss.Length() != 3 {
		t.Errorf("error to push strings to stack")
	}

	st, _ := ss.Top()

	if st != "c" {
		t.Errorf("wrong top value")
	}
}

func TestPrint(t *testing.T) {
	is := Stack[int]{}

	is.Push(10)
	is.Push(20)
	is.Push(30)

	is.Print()

	ss := Stack[string]{}

	ss.Push("a")
	ss.Push("b")
	ss.Push("c")

	ss.Print()
}
