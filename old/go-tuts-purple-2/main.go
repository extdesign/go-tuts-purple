package main

import "fmt"

var transactions []float64

func main() {
	arr := make([]int, 5, 10)
	fmt.Println(arr, len(arr), cap(arr))

	for {
		transaction := fetchUserTransaction()

		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	var balance float64 = calculateBalance()
	fmt.Printf("Сумма транзакций: %.2f\n", balance)
}

func fetchUserTransaction() float64 {
	fmt.Print("Введите число [0 для отмены]: ")
	var value float64
	fmt.Scan(&value)
	return value
}

func calculateBalance() float64 {
	var balance float64
	for _, value := range transactions {
		balance += value
	}

	return balance
}
