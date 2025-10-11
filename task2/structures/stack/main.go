package main

import "fmt"

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека

}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) {
	if s.head >= len(s.s)-1 {
		panic("стек переполнен")
	}
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func pop(s *stack) any {
	if s.head == -1 {
		panic("стек пуст")
	}
	value := s.s[s.head]
	s.head--
	return value
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head == -1 {
		panic("стек пуст")
	}
	return s.s[s.head]
}

// isEmpty - проверка пустоты стека
func isEmpty(s *stack) bool {
	return s.head == -1
}

// size - текущий размер стека
func size(s *stack) int {
	return s.head + 1
}

func main() {
	st := newStack(5)

	push(st, 10)
	push(st, "hello")
	push(st, 3.14)

	fmt.Println("Вершина стека:", peek(st))
	fmt.Println("Размер:", size(st))

	for !isEmpty(st) {
		fmt.Println("Вытолкнули:", pop(st))
	}

	fmt.Println("Стек пуст?", isEmpty(st))

}
