package singleList

import (
	"testing"
)

func TestNewListIsEmpty(t *testing.T) {
	list := New[int]()
	if list.size != 0 {
		t.Error("список должен быть пустой")
	}
	if list.first != nil || list.last != nil {
		t.Error("первый(first) и последний(last) элементы должы быть пустые")
		// далее в местах, где проверка first и last, будут использоваться слова "первый" и "последний" соответственно
	}
}

func TestListAddSingle(t *testing.T) {
	list := New[string]()
	list.Add("s")

	if list.size != 1 {
		t.Error("размер должен быть 1")
	}
	if list.first == nil || list.last == nil {
		t.Fatal("первый(first) и последний(last) элементы не должы быть пустые")
	}
	if list.first.v != "s" {
		t.Errorf("ждали первый s, получили %q", list.first.v)
	}
	if list.last.v != "s" {
		t.Errorf("ждали послединй s, получили %q", list.last.v)
	}
	if list.first.next != nil {
		t.Error("next у элемента s есть. не должен быть, элемент один")
	}
}

func TestListAddMultiple(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.size != 3 {
		t.Error("размер должен быть 3")
	}

	if list.first.v != 1 {
		t.Error("значение должно быть 1")
	}
	if list.first.next.v != 2 {
		t.Error("значение должно быть 2")
	}
	if list.first.next.next.v != 3 {
		t.Error("значение должно быть 3")
	}
	if list.last.v != 3 {
		t.Error("последнее значение должно быть 3")
	}
	if list.last.next != nil {
		t.Error("next у последнего элемента есть. не должен быть, элемент последний")
	}
}

func TestListGet(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	if list.Get(0) != "a" {
		t.Error("значение должно быть a")
	}
	if list.Get(1) != "b" {
		t.Error("значение должно быть b")
	}
	if list.Get(2) != "c" {
		t.Error("значнение должно быть c")
	}
}

func TestListGetInvalidIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен был случиться panic")
		}
	}()

	list := New[int]()
	list.Add(100)
	list.Get(1)
}

func TestListRemoveOnlyElement(t *testing.T) {
	list := New[string]()
	list.Add("s")

	list.Remove(0)

	if list.size != 0 {
		t.Error("размер должен быть 0")
	}
	if list.first != nil || list.last != nil {
		t.Error("первый и последний элементы должы быть пустые")
	}
}

func TestListRemoveFirst(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.Remove(0)

	if list.size != 2 {
		t.Error("размер должен быть 2")
	}
	if list.Get(0) != 2 {
		t.Error("первый элемент должен быть 2")
	}
	if list.Get(1) != 3 {
		t.Error("второй элемент должен быть 3")
	}
}

func TestListRemoveLast(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	list.Remove(2)

	if list.size != 2 {
		t.Error("размер должен быть 2")
	}
	if list.last.v != "b" {
		t.Error("последний должен быть b")
	}
	if list.last.next != nil {
		t.Error("next у последнего элемента есть. не должен быть, элемент последний")
	}
}

func TestListRemoveMiddle(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.Remove(1)

	if list.size != 2 {
		t.Error("размер должен быть 2")
	}
	if list.Get(0) != 1 {
		t.Error("значение должно быть 1")
	}
	if list.Get(1) != 3 {
		t.Error("значение должно быть 3")
	}
}

func TestListRemoveInvalidIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("должен был случиться panic")
		}
	}()

	list := New[int]()
	list.Remove(0)
}

func TestListValues(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	values := list.Values()
	expected := []string{"a", "b", "c"}

	if len(values) != len(expected) {
		t.Fatalf("разная длина. ждали %d, а по итогу %d", len(expected), len(values))
	}

	for i := range expected {
		if values[i] != expected[i] {
			t.Errorf(" элемент values[%d] = %q, ждали %q", i, values[i], expected[i])
		}
	}
}

func TestListValuesEmpty(t *testing.T) {
	list := New[int]()
	values := list.Values()
	if len(values) != 0 {
		t.Error("размер должен быть 0")
	}
}

func TestListWithFloat(t *testing.T) {
	list := New[float64]()
	list.Add(1.1)
	list.Add(2.2)
	if list.Get(0) != 1.1 {
		t.Error("значение должно быть 1.1")
	}
	list.Remove(0)
	if list.Get(0) != 2.2 {
		t.Error("значение должно быть 2.2")
	}
}
