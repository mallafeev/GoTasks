package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func readIntInRange(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		if num, err := strconv.Atoi(input); err == nil {
			if num >= min && num <= max {
				return num
			}
			fmt.Printf("Число должно быть от %d до %d.\n", min, max)
		} else {
			fmt.Println("Ошибка: введите целое число.")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rows := readIntInRange("Введите количество строк (1–10): ", 1, 10)
	cols := readIntInRange("Введите количество столбцов (1–10): ", 1, 10)

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	used := make(map[int]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var num int
			for {
				num = rand.Intn(1000) + 1
				if !used[num] {
					used[num] = true
					break
				}
			}
			matrix[i][j] = num
		}
	}

	fmt.Println("\nСгенерированная матрица:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%4d ", matrix[i][j])
		}
		fmt.Println()
	}
}
