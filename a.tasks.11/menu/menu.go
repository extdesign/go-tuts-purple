package menu

import (
	"a.tasks.11/output"
)

func ShowMenu() {
	menu := [4]string{
		"1. Создать аккаунт",
		"2. Найти аккаунт",
		"3. Удалить аккаунт",
		"4. Выход",
	}

	for _, item := range menu {
		output.Print(item + "\n")
	}
}
