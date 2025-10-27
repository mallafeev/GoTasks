package stack

import (
	"testing"
)

func TestNewStackIsEmpty(t *testing.T) {
	st := New[int](5)
	if !st.IsEmpty() {
		t.Error("стек должен быть пустой")
	}
	if st.Size() != 0 {
		t.Error("размер должен быть 0")
	}
}

func TestStackPushSingle(t *testing.T) {
	st := New[string](3)
	st.Push("x")

	if st.IsEmpty() {
		t.Error("стек не должен быть пустой")
	}
	if st.Size() != 1 {
		t.Error("размер должен быть 1")
	}
	if st.Peek() != "x" {
		t.Errorf("ждали x, получили %q", st.Peek())
	}
}

func TestStackPushMultiple(t *testing.T) {
	st := New[int](5)
	st.Push(1)
	st.Push(2)
	st.Push(3)

	if st.Size() != 3 {
		t.Error("размер должен быть 3")
	}
	if st.Peek() != 3 {
		t.Errorf("ждали 3, получили %d", st.Peek())
	}
}

func TestStackPop(t *testing.T) {
	st := New[string](3)
	st.Push("a")
	st.Push("b")
	st.Push("c")

	if st.Pop() != "c" {
		t.Error("ждали c")
	}
	if st.Pop() != "b" {
		t.Error("ждали b")
	}
	if st.Pop() != "a" {
		t.Error("ждали a")
	}
	if !st.IsEmpty() {
		t.Error("стек должен быть пустой")
	}
	if st.Size() != 0 {
		t.Error("размер должен быть 0")
	}
}

func TestStackPeek(t *testing.T) {
	st := New[int](3)
	st.Push(1)
	st.Push(2)

	if st.Peek() != 2 {
		t.Error("ждали 2")
	}
	if st.Size() != 2 {
		t.Error("размер должен быть 2")
	}
}

func TestStackPushOnFull(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен был случиться panic")
		}
	}()

	st := New[int](2)
	st.Push(1)
	st.Push(2)
	st.Push(3)
}

func TestStackPopOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен был случиться panic")
		}
	}()

	st := New[string](3)
	st.Pop()
}

func TestStackPeekOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен был случиться panic")
		}
	}()

	st := New[int](3)
	st.Peek()
}

func TestStackSize(t *testing.T) {
	st := New[float64](4)
	if st.Size() != 0 {
		t.Error("размер должен быть 0")
	}

	st.Push(1.1)
	if st.Size() != 1 {
		t.Error("размер должен быть 1")
	}

	st.Push(2.2)
	if st.Size() != 2 {
		t.Error("размер должен быть 2")
	}

	st.Pop()
	if st.Size() != 1 {
		t.Error("размер должен быть 1")
	}
}

func TestStackWithBool(t *testing.T) {
	st := New[bool](2)
	st.Push(true)
	st.Push(false)

	if st.Pop() != false {
		t.Error("ждали false")
	}
	if st.Pop() != true {
		t.Error("ждали true")
	}
}
