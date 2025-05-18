package main

import (
	s "IMT/internal/servise"
)

func main() {
	userHeight := s.GetPositiveInput("Введите рост в метрах (пример: 1.75): \n")
	userKg := s.GetPositiveInput("Введите вес в килограммах (пример: 68.5): \n")

	imt := s.CalculateIMT(userHeight, userKg)
	s.ShowIMTMessage(imt)
}
