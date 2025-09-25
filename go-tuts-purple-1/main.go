package main

import (
	"errors"
	"fmt"
	"strconv"
)

type currencyMap = map[string]float64

func main() {
	var err error

	var userData = map[string]string{
		"from":  "",
		"to":    "",
		"value": "",
	}

	err = fetchCurrencyFromUser(userData, "from")
	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}

	err = fetchValueFromUser(userData)
	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}

	err = fetchCurrencyFromUser(userData, "to")
	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}

	result, err := convert(userData)

	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}

	fmt.Printf("%.2f", result)
}

func currencies() currencyMap {
	return currencyMap{
		"RUB":  1,
		"USD":  64.57,
		"EURO": 77.48,
	}
}

func fetchValueFromUser(userData map[string]string) error {
	fmt.Print("Введите число: ")

	var value float64 = 0
	fmt.Scan(&value)

	if value <= 0 {
		return errors.New("value must be grather than 0")
	}

	userData["value"] = fmt.Sprintf("%.2f", value)
	return nil
}

func fetchCurrencyFromUser(userData map[string]string, code string) error {
	var currenciesCurrent []string

	for key := range currencies() {
		if key == code {
			continue
		}

		currenciesCurrent = append(currenciesCurrent, key)
	}

	fmt.Printf("Введите валюту %v ", currenciesCurrent)
	var currency string = ""
	fmt.Scan(&currency)

	if !checkCurrencyFromUser(currency) {
		return errors.New("wrong currency")
	}

	userData[code] = currency

	return nil
}

func checkCurrencyFromUser(currency string) bool {
	return currencies()[currency] != 0
}

func convert(userData map[string]string) (float64, error) {
	value64, err := strconv.ParseFloat(userData["value"], 64)

	if err != nil {
		return 0.00, nil
	}

	if userData["from"] == userData["to"] {
		return value64, nil
	}

	return value64 * (currencies()[userData["from"]] / currencies()[userData["to"]]), nil
}
