package main

import (
	"fmt"
	"strings"
)

func isValidRoman(s string) bool {
	valid := "IVXLCDM"
	for i := 0; i < len(s); i++ {
		c := s[i]
		found := false
		for j := 0; j < len(valid); j++ {
			if c == valid[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func romanToArabic(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]]
		if value < prev {
			total -= value
		} else {
			total += value
		}
		prev = value
	}

	return total
}

func main() {
	const validSymbols = "I, V, X, L, C, D, M"

	for {
		fmt.Print("Введите римское число: ")
		var input string
		fmt.Scanln(&input)

		input = strings.ToUpper(strings.TrimSpace(input))

		if input == "" {
			fmt.Println("Ввод не может быть пустым.")
			continue
		}

		if !isValidRoman(input) {
			fmt.Printf("Некорректный ввод. Допустимые символы: %s\n", validSymbols)
			continue
		}

		result := romanToArabic(input)
		fmt.Printf("Арабское число: %d\n", result)
		break
	}
}
