package main

import (
	"fmt"
)

func main() {
	// a := 5
	// res := double(a)

	// fmt.Printf("a: %d\n", a)
	// fmt.Printf("doubleA: %d\n", res)

	// double2(&a)
	// fmt.Printf("a after double: %d\n", a)
	// fmt.Printf("doubleA: %d\n", res)

	// var b *int
	// unnil(&b)

	// main2()

	a := []int{1, 2, 3}
	changeArrayA(a)
	fmt.Println("a", a)

	sliceA := a[1:3]
	changeArrayA(sliceA)

	fmt.Println("a", a, "sliceA", sliceA)
}

func changeArrayA(a []int) {
	a[0] = 10
	fmt.Println("a in function:", a)
}

// func double(n int) int {
// 	return n * 2
// }

// func double2(n *int) {
// 	*n *= 2
// }

func main2() {
	a := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	reverse(&a)
}

func reverse(a *[11]int) {
	half := len(*a) / 2

	for i := len(*a) - 1; i > half; i-- {
		a[i], a[len(*a)-i-1] = a[len(*a)-i-1], a[i]
	}
	fmt.Println(a)

	for index, value := range *a {
		(*a)[len(a)-1-index] = value
	}
	fmt.Println(a)
}
