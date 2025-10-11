package main

import "fmt"

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

// add - добавление значения в связный список
func add(l *singlyLinkedList, v any) {
	newItem := &item{
		v:    v,
		next: nil,
	}
	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		panic("индекс вне диапазона")
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		panic("индекс вне диапазона")
	}

	if idx == 0 {
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
	} else {
		current := l.first
		for i := 0; i < idx-1; i++ {
			current = current.next
		}
		toRemove := current.next
		current.next = toRemove.next
		if toRemove == l.last {
			l.last = current
		}
	}

	l.size--
}

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []any {
	result := make([]any, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}

func main() {
	list := newSinglyLinkedList()

	add(list, "A")
	add(list, "B")
	add(list, "C")

	fmt.Println("Список:", values(list))
	fmt.Println("Элемент 1:", get(list, 1))

	remove(list, 1)
	fmt.Println("После удаления 1 элемента:", values(list))

	remove(list, 0)
	fmt.Println("После удаления 0 элемента:", values(list))

	remove(list, 0)
	fmt.Println("Пустой список:", values(list))
	fmt.Println("Размер:", list.size)
}
