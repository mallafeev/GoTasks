package stack

type Stack[T any] struct {
	s    []T // слайс в котором хранятся значения в стеке
	head int // индекс головы стека
}

func New[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, size),
		head: -1,
	}
}

// push - добавление в стек значения
func (s *Stack[T]) Push(v T) {
	if s.head >= len(s.s)-1 {
		panic("стек переполнен")
	}
	s.head++
	s.s[s.head] = v
}

// pop - получения значения из стека и его удаление из вершины
func (s *Stack[T]) Pop() T {
	if s.head == -1 {
		panic("стек пуст")
	}
	value := s.s[s.head]
	s.head--
	return value
}

// peek - просмотр значения на вершине стека
func (s *Stack[T]) Peek() T {
	if s.head == -1 {
		panic("стек пуст")
	}
	return s.s[s.head]
}

// isEmpty - проверяет, пуст ли стек (опционально, но полезно)
func (s *Stack[T]) IsEmpty() bool {
	return s.head == -1
}

// size - текущий размер стека (опционально)
func (s *Stack[T]) Size() int {
	return s.head + 1
}
