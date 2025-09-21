package main

import (
	"errors"
	"fmt"
	"slices"
)

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

func currencies() []string {
	return []string{"RUB", "USD", "EURO"}
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

	if len(USER_CURRENCY_FROM) != 0 {
		for _, currency := range currencies() {
			if currency != USER_CURRENCY_FROM {
				currenciesCurrent = append(currenciesCurrent, currency)
			}
		}
	} else {
		currenciesCurrent = currencies()
	}

	fmt.Printf("Введите валюту %v ", currenciesCurrent)
	var currency string = ""
	fmt.Scan(&currency)

	if !checkCurrencyFromUser(currenciesCurrent, currency) {
		return "", errors.New("wrong currency")
	}

	return currency, nil
}

func checkCurrencyFromUser(currencies []string, currency string) bool {
	return slices.Contains(currencies, currency)
}

func convertEuroToRub() float64 {
	return USD_TO_EURO * USD_TO_RUB
}

func convert() float64 {
	if USER_CURRENCY_FROM == USER_CURRENCY_TO {
		return USER_VALUE
	}

	switch USER_CURRENCY_FROM {
	case "RUB":

		switch USER_CURRENCY_TO {
		case "USD":
			return USER_VALUE / USD_TO_RUB
		case "EURO":
			return USER_VALUE / convertEuroToRub()
		}

	case "USD":
		switch USER_CURRENCY_TO {
		case "RUB":
			return USER_VALUE * USD_TO_RUB
		case "EURO":
			return USER_VALUE * USD_TO_EURO
		}

	case "EURO":
		switch USER_CURRENCY_TO {
		case "USD":
			return USER_VALUE / USD_TO_EURO
		case "RUB":
			return USER_VALUE / convertEuroToRub()
		}
	}

	return 0
}
