package main

import (
	"fmt"
	"task3/queue"
	"task3/singleList"
	"task3/stack"
)

func main() {
	// --- Стек ---
	st := stack.New[string](5)

	st.Push("hello")
	st.Push("world")
	st.Push("!")

	fmt.Println("Вершина стека:", st.Peek())
	fmt.Println("Размер:", st.Size())

	for !st.IsEmpty() {
		fmt.Println("Вытолкнули:", st.Pop())
	}

	fmt.Println("Стек пуст?", st.IsEmpty())

	// --- Очередь ---
	q := queue.New[string](3)

	q.Push("первый")
	q.Push("второй")
	q.Push("третий")

	fmt.Println("Вытащили:", q.Pop())
	fmt.Println("Вытащили:", q.Pop())

	q.Push("новый")

	fmt.Println("Вытащили:", q.Pop())
	fmt.Println("Вытащили:", q.Pop())

	fmt.Println("Очередь пуста?", q.IsEmpty())

	// --- Список ---
	list := singleList.New[string]()

	list.Add("A")
	list.Add("B")
	list.Add("C")

	fmt.Println("Список:", list.Values())
	fmt.Println("Элемент 1:", list.Get(1))

	list.Remove(1)
	fmt.Println("После удаления 1 элемента:", list.Values())

	list.Remove(0)
	fmt.Println("После удаления 0 элемента:", list.Values())

	list.Remove(0)
	fmt.Println("Пустой список:", list.Values())
}
