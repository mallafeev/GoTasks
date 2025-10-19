package queue

type Queue[T any] struct {
	s         []T // слайс в котором хранятся значения
	low, high int // индексы верхней и нижней границы очереди
	size      int // размер очереди
}

func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		s:    make([]T, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

// push - добавление в очередь значения
func (q *Queue[T]) Push(v T) {
	if (q.high+1)%q.size == q.low {
		panic("очередь переполнена")
	}
	if q.low == -1 {
		q.low = 0
		q.high = 0
	} else {
		q.high = (q.high + 1) % q.size
	}
	q.s[q.high] = v
}

// pop - получения значения из очереди и его удаление
// в вашей заготовке в pop передаётся ещё параметр v any, который не надо передавать. убрал.
func (q *Queue[T]) Pop() T {
	if q.low == -1 {
		panic("очередь пуста")
	}
	value := q.s[q.low]
	if q.low == q.high {
		q.low = -1
		q.high = -1
	} else {
		q.low = (q.low + 1) % q.size
	}
	return value
}

// isEmpty — проверка пустоты очереди
func (q *Queue[T]) IsEmpty() bool {
	return q.low == -1
}
