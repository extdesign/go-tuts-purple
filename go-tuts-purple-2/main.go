package main

import "fmt"

var transactions []float64

func main() {
	for {
		transaction := fetchUserTransaction()
		transactions = append(transactions, transaction)

		fmt.Print("Добавить еще транзакцию? [y,n]: ")
		var answer string = ""
		fmt.Scan(&answer)

		if answer != "y" {
			break
		}
	}

	printResult()
}

func fetchUserTransaction() float64 {
	fmt.Print("Введите число: ")
	var value float64 = 0
	fmt.Scan(&value)
	return value
}

func printResult() {
	fmt.Println(transactions)
}
