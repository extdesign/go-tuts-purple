package main

import "fmt"

const USD_TO_EURO float64 = 1.20
const USD_TO_RUB float64 = 64.57

func main() {
	fmt.Printf("Вы можете купить 1.00 euro за %.2f рублей\n", convertEuroToRub())

	userValue := getUserValue()
	fmt.Printf("За %.2f USD вы получите XXX RUB\n", convert(userValue, "USD", "RUB"))
}

func convertEuroToRub() float64 {
	return USD_TO_EURO * USD_TO_RUB
}

func getUserValue() float64 {
	fmt.Println("Введите сумму:")

	var userValue float64 = 0.00
	fmt.Scan(&userValue)

	return userValue
}

func convert(userValue float64, currencyFrom string, currencyTo string) float64 {
	return userValue
}
