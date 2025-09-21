package main

import "fmt"

var transactions []float64

func main() {
	for {
		transaction := fetchUserTransaction()

		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	printResult()
}

func fetchUserTransaction() float64 {
	fmt.Print("Введите число [0 для отмены]: ")
	var value float64
	fmt.Scan(&value)
	return value
}

func printResult() {
	fmt.Println(transactions)
}
