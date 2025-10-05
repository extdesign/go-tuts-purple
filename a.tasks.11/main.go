package main

import (
	"fmt"

	"a.tasks.11/account"
	"a.tasks.11/menu"
	"a.tasks.11/output"
)

func main() {

	output.PrintBlue("\n__ Менеджер паролей ___\n\n")
	vault := account.NewVault()

menuVariant:
	for {
		menu.ShowMenu()
		menuItem := output.Prompt("\nВыберете пункт меню: ")

		switch menuItem {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			removeAccount(vault)
		case "4":
			break menuVariant
		}
	}
}

func createAccount(vault *account.Vault) {
	login := output.Prompt("Введите логин: ")
	password := output.Prompt("Введите пароль: ")
	url := output.Prompt("Введите URL: ")

	myAccount, err := account.NewAccount(&login, &password, &url)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nОшибки при инициализации: %s\n\n", err.Error()))
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := output.Prompt("Введите URL: ")

	foundedAccounts := vault.FindAccountsByUrl(url)

	if len(foundedAccounts) > 0 {
		for _, acc := range foundedAccounts {
			acc.Output()
		}
	} else {
		output.PrintRed("\nНет аккаунтов с таким URL\n\n")
	}
}

func removeAccount(vault *account.Vault) {
	url := output.Prompt("Введите URL: ")
	err := vault.RemoveAccountByUrl(url)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось удалить аккаунт. Ошибки: %s\n\n", err))
	} else {
		output.PrintYellow("Аккаунт удален\n\n")
	}
}
