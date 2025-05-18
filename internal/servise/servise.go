package servise

import (
	"fmt"
	"math"
)

const returnMessage = "Ваш индекс массы тела: %.2f"

func GetPositiveInput(prompt string) float64 {
	var value float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&value)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			continue
		}
		if value <= 0 {
			fmt.Println("Число должно быть больше нуля. Попробуйте снова.")
			continue
		}
		break
	}
	return value
}

func CalculateIMT(height, weight float64) float64 {
	return weight / math.Pow(height, 2)
}

func ShowIMTMessage(imt float64) {
	switch {
	case imt < 16.0:
		fmt.Printf(returnMessage+", выраженный дефицит массы тела.\n", imt)
	case imt >= 16.0 && imt < 18.5:
		fmt.Printf(returnMessage+", недостаточная масса тела.\n", imt)
	case imt >= 18.5 && imt < 25.0:
		fmt.Printf(returnMessage+", норма. Молодец!\n", imt)
	case imt >= 25.0 && imt < 30.0:
		fmt.Printf(returnMessage+", избыточная масса тела.\n", imt)
	case imt >= 30.0 && imt < 35.0:
		fmt.Printf(returnMessage+", ожирение I степени.\n", imt)
	case imt >= 35.0 && imt < 40.0:
		fmt.Printf(returnMessage+", ожирение II степени.\n", imt)
	default:
		fmt.Printf(returnMessage+", ожирение III степени. Обратитесь к врачу.\n", imt)
	}
}
