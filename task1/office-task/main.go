package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

/*
Scanf работает некорректно при swicth case, потому что когда мы вводим пункт в меню "1" и жмём enter,
программа передаёт /n дальше, тем самым первое поле заполняется пустотой.
Я долго пытался пофиксить это, не трогая оригинальный код, но не вышло.
В реальной работе я бы пошел к разработчику этого элемента кода и попросил фикса, но тут я сам по себе и сам пофикшу.
Scanln тоже не работает у меня, как только не пробовал, поэтому сделал через библиотеку
*/
func main() {
	const size = 512
	empls := [size]*Employee{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(commands)
		if !scanner.Scan() {
			break
		}
		cmdStr := scanner.Text()
		cmd, err := strconv.Atoi(cmdStr)
		if err != nil {
			fmt.Println("Неверная команда")
			continue
		}

		switch cmd {
		case 1:
			empl := new(Employee)
			fmt.Println("\nИмя:")
			scanner.Scan()
			empl.Name = scanner.Text()

			fmt.Println("Возраст:")
			scanner.Scan()
			// добавим проверку на число ещё по-хорошему
			if age, err := strconv.Atoi(scanner.Text()); err == nil {
				empl.Age = age
			} else {
				fmt.Println("Неверный возраст, введите число")
				empl.Age = 0
			}

			fmt.Println("Позиция:")
			scanner.Scan()
			empl.Position = scanner.Text()

			fmt.Println("Зарплата:")
			scanner.Scan()
			if salary, err := strconv.Atoi(scanner.Text()); err == nil {
				empl.Salary = salary
			} else {
				fmt.Println("Неверная зарплата, нужно ввести число")
				empl.Salary = 0
			}
			added := false
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					added = true
					fmt.Println("Сотрудник добавлен")
					break
				}
			}
			if !added {
				fmt.Println("Лимит 512 сотрудников достигнут")
			}
		case 2:
			// думаю самый корректный вариант это удалять по индексу в массиве, ведь имена или другие поля могут повторяться
			fmt.Println("Список сотрудников:")
			count := 0
			for i := 0; i < len(empls); i++ {
				if empls[i] != nil {
					count++
					fmt.Printf("#%d. %s\n", count, empls[i].Name)
				}
			}
			if count == 0 {
				fmt.Println("Некого удалять")
				continue
			}

			fmt.Print("Введите номер сотрудника для удаления: ")
			if !scanner.Scan() {
				fmt.Println("Ошибка ввода")
				continue
			}
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num > count {
				fmt.Println("Неверный номер")
				continue
			}
			actualIndex := -1
			currentCount := 0
			for i := 0; i < len(empls); i++ {
				if empls[i] != nil {
					currentCount++
					if currentCount == num {
						actualIndex = i
						break
					}
				}
			}

			if actualIndex != -1 {
				empls[actualIndex] = nil
				fmt.Println("Сотрудник удалён")
			} else {
				fmt.Println("Ошибка: сотрудник не найден")
			}
		case 3:
			fmt.Println("Вывод сотрудников")
			count := 0
			for i := 0; i < len(empls); i++ {
				if empls[i] != nil {
					count++
					fmt.Printf("Сотрудник #%d. Имя: %s, Возраст: %d, Позиция: %s, Зарплата: %d\n",
						count, empls[i].Name, empls[i].Age, empls[i].Position, empls[i].Salary)
				}
			}
			if count == 0 {
				fmt.Println("никого нет")
			}
		case 4:
			fmt.Println("Выходим с программы")
			return
		}
	}
}
