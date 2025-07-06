package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Создать приложение, которое сначала выдаёт меню:
– 1. Посмотреть закладки
– 2. Добавить закладку
– 3. Удалить закладку
– 4. Выход
При 1 – Выводит закладки
При 2 – 2 поля ввода названия и адреса и после добавление
При 3 – Ввод названия и удаление по нему
При 4 – Завершение
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	bookmarks := map[string]string{}

	fmt.Print(
		`Команды для ввода:
– 1. Посмотреть закладки
– 2. Добавить закладку
– 3. Удалить закладку
– 4. Выход
`,
	)
Menu:
	for {
		fmt.Print("введите нужную команду: ")
		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			length := len(bookmarks)
			if length == 0 {
				fmt.Println("Список закладок пуст")
			} else {
				for key, value := range bookmarks {
					fmt.Println(key, "-", value)
				}
			}

			continue
		case "2":
			fmt.Println("введите ключ: ")
			scanner.Scan()
			key := strings.TrimSpace(scanner.Text())
			if key == "" {
				fmt.Println("Ошибка! ожидается, что будет введено корректное значение (строка не может быть пустой)!")
				continue
			}

			fmt.Println("введите значение: ")
			scanner.Scan()
			value := strings.TrimSpace(scanner.Text())

			if value == "" {
				fmt.Println("Ошибка! ожидается, что будет введено корректное значение (строка не может быть пустой)!")
				continue
			}

			bookmarks[key] = value
			println("Закладка с ключом " + key + " успешно добавлена!")
			continue
		case "3":
			fmt.Println("введите ключ: ")
			scanner.Scan()
			key := strings.TrimSpace(scanner.Text())

			_, isExist := bookmarks[key]

			if isExist {
				delete(bookmarks, key)
				println("Закладка с ключом " + key + " успешно удаленна!")
			} else {
				println("Закладки с ключом " + key + " не была найдена в хранилище...")
			}

			continue
		case "4":
			break Menu
		}
	}

}
