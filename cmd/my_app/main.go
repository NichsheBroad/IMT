package main

import (
	s "IMT/internal/servise"
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover: ", r)
		}
	}()

	fmt.Println("__Ку вы начали работу с IMT(Калькулятором индекса массы тела!__")

	for {

		userHeight := s.GetPositiveInput("Введите рост в метрах (пример: 1.75): \n")
		userKg := s.GetPositiveInput("Введите вес в килограммах (пример: 68.5): \n")

		imt := s.CalculateIMT(userHeight, userKg)
		s.ShowIMTMessage(imt)
		repeateCalculation := s.CheckRepeateCalcukation()
		if !repeateCalculation {
			fmt.Println("Спасибо что пользовались мной)..\nбб")
			break
		}
	}
}
