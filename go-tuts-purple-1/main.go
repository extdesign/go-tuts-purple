package main

import (
	"errors"
	"fmt"
)

type currencyMap = map[string]float64

const USD_TO_EURO float64 = 1.20
const USD_TO_RUB float64 = 64.57

var USER_VALUE float64 = 0
var USER_CURRENCY_FROM string = ""
var USER_CURRENCY_TO string = ""

func main() {
	var err error

	for {
		USER_CURRENCY_FROM, err = fetchCurrencyFromUser()

		if err == nil {
			break
		}
	}

	for {
		USER_VALUE, err = fetchValueFromUser()

		if err == nil {
			break
		}
	}

	for {
		USER_CURRENCY_TO, err = fetchCurrencyFromUser()

		if err == nil {
			break
		}
	}

	var result float64 = convert()
	fmt.Printf("%.2f", result)
}

func currencies() currencyMap {
	return currencyMap{
		"RUB":  1,
		"USD":  64.57,
		"EURO": 77.48,
	}
}

func fetchValueFromUser() (float64, error) {
	fmt.Print("Введите число: ")

	var value float64 = 0
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("value must be grather than 0")
	}

	return value, nil
}

func fetchCurrencyFromUser() (string, error) {
	var currenciesCurrent []string

	for key := range currencies() {
		if key == USER_CURRENCY_FROM {
			continue
		}

		currenciesCurrent = append(currenciesCurrent, key)
	}

	fmt.Printf("Введите валюту %v ", currenciesCurrent)
	var currency string = ""
	fmt.Scan(&currency)

	if !checkCurrencyFromUser(currency) {
		return "", errors.New("wrong currency")
	}

	return currency, nil
}

func checkCurrencyFromUser(currency string) bool {
	return currencies()[currency] != 0
}

func convert() float64 {
	if USER_CURRENCY_FROM == USER_CURRENCY_TO {
		return USER_VALUE
	}

	return USER_VALUE * (currencies()[USER_CURRENCY_FROM] / currencies()[USER_CURRENCY_TO])
}
