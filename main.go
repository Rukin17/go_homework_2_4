// Напишите программу, которая запрашивает у пользователя ввод строки-формулы, а выводит сообщение о правильности написания круглых скобок, например:

// Пример 1)
// строка на вход: (1+1)*(2+2)
// вывод: Скобки расставлены верно, 2 открывающиеся, 2 закрывающиеся

// Пример 2)
// Строка на вход: ((1+1) + (2+2) ))
// вывод: Скобки расставлены неправильно, 3 открывающиеся, 4 закрывающиеся

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку: ")
	input, _ := reader.ReadString('\n')
	result, openBracketCount, closedBracketCount := checkCorrectSequenceBracket(input)

	if result {
		fmt.Printf("Скобки расставлены Верно: %d открывающих и %d закрывающих.\n", openBracketCount, closedBracketCount)
	} else {
		fmt.Printf("Скобки расставлены Неверно: %d открывающих и %d закрывающих.\n", openBracketCount, closedBracketCount)
	}

}

func pop(stack []string) []string {
	lastElem := len(stack) - 1
	stack = stack[:lastElem]
	return stack
}

func checkCorrectSequenceBracket(s string) (bool, int, int) {
	var stack []string
	runes := []rune(s)
	openBracketCount, closedBracketCount := 0, 0

	for i := range runes {
		if runes[i] == '(' {
			stack = append(stack, string(runes[i]))
			openBracketCount++

		} else if runes[i] == ')' {
			closedBracketCount++
			if len(stack) == 0 { // В этом месте бага
				return false, openBracketCount, closedBracketCount
			} else {
				stack = pop(stack)
			}

		} else {
			continue
		}
	}
	return len(stack) == 0, openBracketCount, closedBracketCount
}
