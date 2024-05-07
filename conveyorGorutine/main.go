package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите числа. Для завершения программы введите \"стоп\".")

	numChan := make(chan int)
	resultChan := make(chan int)

	go func() {
		for {
			num := <-numChan
			if num == -1 {
				break
			}
			square := num * num
			resultChan <- square
		}
		close(resultChan)
	}()

	go func() {
		for {
			square, ok := <-resultChan
			if !ok {
				break
			}
			product := square * 2
			fmt.Println("Результат:", product)
		}
	}()

	for {
		var input string
		fmt.Scan(&input)

		if input == "стоп" {
			close(numChan)
			fmt.Println("Программа завершена.")
			return
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введено не число")
			continue
		}

		numChan <- num
	}
}
