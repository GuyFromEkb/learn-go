package scan

import (
	"bufio"
	"fmt"
	"os"
)

func ScanPrompt(placeholder string, canEmpty bool) string {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(placeholder)
		scanner.Scan()
		input := scanner.Text()

		if input == "" && !canEmpty {
			fmt.Println("Некорректный ввод. Значение не может быть пустым, повторите попытку.")
			continue
		}

		return input
	}

}
