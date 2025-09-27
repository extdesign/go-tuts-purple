package main

import (
	"fmt"
	"project-4-password/account"
	"project-4-password/files"

	"github.com/fatih/color"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	myAccount, err := account.NewAccountWithTimeStamp(&login, &password, &url)

	if err != nil {
		red := color.New(color.FgRed).PrintfFunc()
		red("Ошибки при инициализации: %s", err.Error())
		//fmt.Printf("Ошибки при инициализации: %s\n", err.Error())
		return
	}

	fmt.Printf("Ваши данные: %v", myAccount)

	files.WriteFile()
}

func promptData(prompt string) string {
	c := color.New(color.FgCyan)
	c.Print(prompt)

	var res string
	fmt.Scanln(&res)
	return res
}
