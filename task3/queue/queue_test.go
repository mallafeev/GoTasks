package queue

import (
	"testing"
)

func TestNewQueueIsEmpty(t *testing.T) {
	q := New[int](5)
	if !q.IsEmpty() {
		t.Error("очередь должна быть пустой")
	}
}

func TestQueuePushPopSingle(t *testing.T) {
	q := New[string](3)
	q.Push("s")
	if q.IsEmpty() {
		t.Error("очередь не должна быть пустой. добавление символа s")
	}
	value := q.Pop()
	if value != "s" {
		t.Errorf("ждали s, получили %q", value)
	}
	if !q.IsEmpty() {
		t.Error("очередь должна быть пустой")
	}
}

func TestQueueFIFO(t *testing.T) {
	q := New[int](5)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Pop() != 1 {
		t.Error("значение должно быть 1")
	}
	if q.Pop() != 2 {
		t.Error("значение должно быть 2")
	}
	if q.Pop() != 3 {
		t.Error("значение должно быть 3")
	}
	if !q.IsEmpty() {
		t.Error("очередь должна быть пустой")
	}
}

func TestQueueCircular(t *testing.T) {
	q := New[int](3)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Pop() != 1 {
		t.Fatal("значение должно быть 1")
	}
	q.Push(4)
	if q.Pop() != 2 {
		t.Error("значение должно быть 2")
	}
	if q.Pop() != 3 {
		t.Error("значение должно быть 3")
	}
	if q.Pop() != 4 {
		t.Error("значение должно быть 4")
	}
	if !q.IsEmpty() {
		t.Error("очередь должна быть пустой")
	}
}

func TestQueuePushOnFull(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен был случиться panic")
		}
	}()

	q := New[int](2)
	q.Push(1)
	q.Push(2)
	q.Push(3)
}

func TestQueuePopOnEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен случиться panic")
		}
	}()

	q := New[string](3)
	q.Pop()
}

func TestQueueWithFloat(t *testing.T) {
	q := New[float64](2)
	q.Push(3.14)
	q.Push(2.71)
	if q.Pop() != 3.14 {
		t.Error("значение должно быть 3.14")
	}
	if q.Pop() != 2.71 {
		t.Error("значение должно быть 2.71")
	}
}
