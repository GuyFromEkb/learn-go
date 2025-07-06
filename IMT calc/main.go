package main

import (
	"fmt"
	"math"
)

type userData struct {
	Height float64
	Weight float64
}

func main() {

	for {

		fmt.Println("__ Расчёт индекса массы тела __")

		userData := readUserData()
		result := getImt(userData)
		printResult(result)

		fmt.Print("Рассчитать ещё раз? (y/n): ")
		var goAgain string
		fmt.Scanln(&goAgain)

		if goAgain == "n" {
			break
		}
	}

}

func readUserData() userData {
	var userData userData

	for {
		fmt.Print("введите ваш рост в сантиметрах: ")
		_, err := fmt.Scanln(&userData.Height)
		if err != nil || userData.Height <= 0 {
			fmt.Println("Ошибка: введите положительное число для роста.")
			clearInputBuffer()
			continue
		}
		break
	}

	for {
		fmt.Print("введите ваш вес в кг: ")
		_, err := fmt.Scanln(&userData.Weight)
		if err != nil || userData.Weight <= 0 {
			fmt.Println("Ошибка: введите положительное число для веса.")
			clearInputBuffer()
			continue
		}
		break
	}

	return userData
}

func clearInputBuffer() {
	var discard string
	fmt.Scanln(&discard)
}

func getImt(userData userData) float64 {
	const IMTPower = 2

	IMT := userData.Weight / math.Pow(userData.Height/100, IMTPower)

	return IMT
}

func printResult(imt float64) {

	var description string

	switch {
	case imt <= 16:
		description = "Выраженный дефицит массы тела"
	case imt > 16 && imt <= 18.5:
		description = "Недостаточная (дефицит) масса тела"
	case imt > 18.5 && imt <= 25:
		description = "Норма"
	case imt > 25 && imt <= 30:
		description = "Избыточная масса тела (предожирение)"
	case imt > 30 && imt <= 35:
		description = "Ожирение первой степени"
	case imt > 35 && imt <= 40:
		description = "Ожирение второй степени"
	default:
		description = "Ожирение третьей степени (морбидное)"
	}

	fmt.Printf("Ваш индекс массы тела: %.2f\n", imt)
	fmt.Println("Состояние:", description)
}
