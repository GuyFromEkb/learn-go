package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	transaction := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for {
		fmt.Print("Введите транзакцию :")
		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		if input == "s" || input == "stop" {

			sum := 0.0

			for _, v := range transaction {
				sum += v
			}

			fmt.Print("Остановка, если сложить все транзакции, то получится:", sum)
			break
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("некорректный ввод...")
			continue
		}

		transaction = append(transaction, value)
		fmt.Println("Текущие транзакции: ", transaction)
		fmt.Println("Что бы посчитать транзакции, введите: 'stop' / 's'")

	}

}
