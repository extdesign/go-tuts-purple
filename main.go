package main

import "fmt"

const USD_TO_EURO float64 = 1.20
const USD_TO_RUB float64 = 64.57

func main() {
	fmt.Println(convertEuroToRub())
}

func convertEuroToRub() float64 {
	return USD_TO_EURO * USD_TO_RUB
}
