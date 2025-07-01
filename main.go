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
	result, openBracketCount, closedBracketCount, err := checkCorrectSequenceBracket(input)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	if result {
		fmt.Printf("Скобки расставлены Верно: %d открывающих и %d закрывающих.\n", openBracketCount, closedBracketCount)
	} else {
		fmt.Printf("Скобки расставлены Неверно: %d открывающих и %d закрывающих.\n", openBracketCount, closedBracketCount)
	}

}

func checkCorrectSequenceBracket(s string) (bool, int, int, error) {
	var err error
	flag := true
	stack := 0 // В первой версии использовал стэк, во второй - счётчик
	runes := []rune(s)
	openBracketCount, closedBracketCount := 0, 0

	for i := range runes {
		switch runes[i] {
		case '(':
			stack++
			openBracketCount++

		case ')':
			closedBracketCount++
			if stack == 0 {
				err = fmt.Errorf("Ставить ')' раньше, чем '(' - это ни в какие ворота..")
				flag = false
			}
			stack--
		}
	}
	// Обрабатываю случай неправильной комбинации скобок, при котором счётчик не изменит флаг:
	if stack != 0 && flag == true {
		flag = false
	}

	return flag, openBracketCount, closedBracketCount, err

}
