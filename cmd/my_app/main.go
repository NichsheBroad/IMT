package main

import (
	"IMT/internal/servise"
)

func main() {
	userHeight := servise.GetPositiveInput("Введите рост в метрах (пример: 1.75): ")
	userKg := servise.GetPositiveInput("Введите вес в килограммах (пример: 68.5): ")

	imt := servise.CalculateIMT(userHeight, userKg)
	servise.ShowIMTMessage(imt)
}
