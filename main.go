package main

import "fmt"

const returnMessage = "Ваш индекс массы тела: %.2f"

func main() {

	var userHeight, userKg, IMT float64

	fmt.Printf("Оптимальный индекс массы тела варьируется от 18.5 до 25.0!\n")

	fmt.Print("\nВведите Рост в метрах: ")
	fmt.Scanln(&userHeight)
	fmt.Println("Вы ввели:", userHeight, "метра.\n")

	fmt.Print("Введите вес: ")
	fmt.Scanln(&userKg)
	fmt.Println("Вы ввели:", userKg, "килограмм.\n")

	IMT = userKg / (userHeight * userHeight)

	if IMT > 25 {
		fmt.Printf(returnMessage+", вам следует обратиться за помощью.\n", IMT)
	} else if IMT < 18.5 {
		fmt.Printf(returnMessage+", вам нужно больше кушац!\n", IMT)
	} else {
		fmt.Printf(returnMessage+", ЭЭЭ ЩЦ да ты красручег!\n", IMT)
	}
}
