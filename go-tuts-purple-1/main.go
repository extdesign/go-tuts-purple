package main

import (
	"errors"
	"fmt"
	"strconv"
)

type currencyMap = map[string]float64

func main() {
	var err error

	var currencyMap = currencyMap{
		"RUB":  1,
		"USD":  64.57,
		"EURO": 77.48,
	}

	var userData = map[string]string{
		"from":  "",
		"to":    "",
		"value": "",
	}

	for {
		err = fetchCurrencyFromUser(userData, "from", &currencyMap)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err.Error())
			continue
		}

		break
	}

	for {
		err = fetchValueFromUser(userData)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err.Error())
			continue
		}

		break
	}

	for {
		err = fetchCurrencyFromUser(userData, "to", &currencyMap)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err.Error())
			continue
		}

		break
	}

	result, err := convert(userData, &currencyMap)

	if err != nil {
		fmt.Printf("Ошибка: %s\n", err.Error())
	}

	fmt.Printf("%.2f", result)
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

func fetchCurrencyFromUser(userData map[string]string, code string, currencyMap *map[string]float64) error {
	var currenciesCurrent []string

	for key := range *currencyMap {
		if key == code {
			continue
		}

		currenciesCurrent = append(currenciesCurrent, key)
	}

	fmt.Printf("Введите валюту %v ", currenciesCurrent)
	var currency string = ""
	fmt.Scan(&currency)

	if _, ok := (*currencyMap)[currency]; !ok {
		return errors.New("wrong currency")
	}

	userData[code] = currency

	return nil
}

func convert(userData map[string]string, currencyMap *map[string]float64) (float64, error) {
	value64, err := strconv.ParseFloat(userData["value"], 64)

	if err != nil {
		return 0.00, nil
	}

	if userData["from"] == userData["to"] {
		return value64, nil
	}

	return value64 * ((*currencyMap)[userData["from"]] / (*currencyMap)[userData["to"]]), nil
}
