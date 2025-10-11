package main

import "fmt"

type queue struct {
	s         []any // слайс в котором хранятся значения
	low, high int   // индексы верхней и нижней границы очереди
	size      int   // размер очереди

}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

// push - добавление в очередь значения
func push(q *queue, v any) {
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
func pop(q *queue) any {
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
func isEmpty(q *queue) bool {
	return q.low == -1
}

func main() {
	q := newQueue(3)

	push(q, "первый")
	push(q, "второй")
	push(q, "третий")

	fmt.Println("Вытащили:", pop(q))
	fmt.Println("Вытащили:", pop(q))

	push(q, "новый")

	fmt.Println("Вытащили:", pop(q))
	fmt.Println("Вытащили:", pop(q))

	fmt.Println("Очередь пуста?", isEmpty(q))
}
