package singleList

type List[T any] struct {
	first *item[T]
	last  *item[T]
	size  int
}

type item[T any] struct {
	v    T
	next *item[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		first: nil,
		last:  nil,
		size:  0,
	}
}

// add - добавление значения в связный список
func (l *List[T]) Add(v T) {
	newItem := &item[T]{
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
func (l *List[T]) Get(idx int) T {
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
func (l *List[T]) Remove(idx int) {
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
func (l *List[T]) Values() []T {
	result := make([]T, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}
