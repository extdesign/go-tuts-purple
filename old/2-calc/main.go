package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__ Calculator __")

	var operation string = fetchUserOperation()
	var numbers string = fetchUserNumbers()
	var preparedNumbers []float64 = prepareNumbers(numbers)
	var result float64

	switch operation {
	case "AVG":
		result = calcAVG(preparedNumbers)
	case "SUM":
		result = calcSUM(preparedNumbers)
	case "MED":
		result = calcMED(preparedNumbers)
	}

	fmt.Printf("%.2f", result)
}

func calcMED(numbers []float64) float64 {
	sort.Float64s(numbers)
	mid := len(numbers) / 2

	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}

	return numbers[mid]
}

func calcAVG(numbers []float64) float64 {
	var result float64

	for _, value := range numbers {
		result += value
	}

	return result / float64(len(numbers))
}

func calcSUM(numbers []float64) float64 {
	var result float64

	for _, value := range numbers {
		result += value
	}

	return result
}

func prepareNumbers(numbers string) []float64 {
	var nums []float64

	for _, v := range strings.Split(numbers, ",") {
		num, _ := strconv.ParseFloat(v, 64)
		nums = append(nums, num)
	}

	return nums
}

func fetchUserOperation() string {
	operation := ""

	for {
		fmt.Printf("Введите операцию %v: ", opertaions())
		fmt.Scan(&operation)

		if slices.Contains(opertaions(), operation) {
			break
		}
	}

	return operation
}

func fetchUserNumbers() string {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Print("Введите числа через запятую: ")

	if scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

func opertaions() []string {
	return []string{"AVG", "SUM", "MED"}
}
