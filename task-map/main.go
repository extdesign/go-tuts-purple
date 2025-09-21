package main

import (
	"fmt"
)

var bookmarks map[string]string

type bookmarkMap = map[string]string

func main() {
	var operation string
	bookmarks = make(bookmarkMap, 5)

menuLoop:
	for {
		showMenu()

		fmt.Print("Ваш выбор: ")
		fmt.Scan(&operation)

		switch operation {
		case "1":
			showBookmarks()
		case "2":
			addBookmark()
		case "3":
			deleteBookmark()
		case "4":
			break menuLoop
		}
	}
}

func showBookmarks() {
	fmt.Println("\nВаши закладки:")
	fmt.Println("----------------")

	if len(bookmarks) == 0 {
		fmt.Println("Закладок нет")
	}

	for key, value := range bookmarks {
		fmt.Printf("[%s] %s\n", key, value)
	}

	fmt.Println("----------------")
}

func addBookmark() {
	var key string
	var value string

	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)

	fmt.Print("Введите значение: ")
	fmt.Scan(&value)

	bookmarks[key] = value
}

func deleteBookmark() {
	var key string

	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)

	delete(bookmarks, key)
}

func showMenu() {
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выйти")
}
