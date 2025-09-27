package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type accountStruct struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	accountStruct
}

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите url: ")

	account, err := newAccountWithTimeStamp(&login, &password, &url)

	if err != nil {
		fmt.Printf("Ошибки при инициализации: %s\n", err.Error())
		return
	}

	account.generatePassword(10)
	fmt.Printf("Ваши данные: %v", account)

	// account2 := accountStruct{
	// 	login,
	// 	password,
	// 	url,
	// }

	// account3 := accountStruct{
	// 	login:    login,
	// 	password: password,
	// }

	// account4 := accountStruct{}

	// outputPassword(account)
	// outputPassword(account2)
	// outputPassword(account3)
	// outputPassword(account4)

	// outputPassword2(&account)

	// presentOfRune()
	// randd()
}

func promptData(prompt string) string {
	fmt.Print(prompt)

	var res string
	fmt.Scanln(&res)
	return res
}

func outputPassword(account accountStruct) {
	fmt.Printf("Ваш логин: %s | ", account.login)
	fmt.Printf("Ваш пароль: %s | ", account.password)
	fmt.Printf("URL: %s\n", account.url)
}

func outputPassword2(account *accountStruct) {
	fmt.Printf("Ваш логин: %s | ", account.login)
	fmt.Printf("Ваш пароль: %s | ", account.password)
	fmt.Printf("URL: %s\n", account.url)
	fmt.Println(account)
}

func presentOfRune() {
	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}
}

func randd() {
	r := rand.IntN(10)
	fmt.Println(r)
}

func (acc *accountStruct) generatePassword(n int) {
	lettersRunes := []rune("abcdefjABCEFJ!_*")
	res := make([]rune, n)

	for i := range n {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	acc.password = string(res)
}

func (acc *accountStruct) outputPassword3() {
	fmt.Printf("Ваши данные: %s\n", acc)
}

func newAccountWithTimeStamp(login, password, urlString *string) (*accountWithTimeStamp, error) {
	acc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		accountStruct: accountStruct{
			login:    *login,
			password: *password,
			url:      *urlString,
		},
	}

	err := acc.validateArguments(login, password, urlString)

	if err != nil {
		return nil, err
	}

	acc.login = *login
	acc.password = *password
	acc.url = *urlString

	return acc, nil
}

// func newAccount(login, password, urlString *string) (*accountStruct, error) {
// 	acc := &accountStruct{}

// 	err := acc.validateArguments(login, password, urlString)

// 	if err != nil {
// 		return nil, err
// 	}

// 	acc.login = *login
// 	acc.password = *password
// 	acc.url = *urlString

// 	return acc, nil
// }

func (acc *accountStruct) validateArguments(login, password, urlString *string) error {
	err := acc.checkUrl(urlString)

	if err != nil {
		return err
	}

	err = acc.checkLogin(login)

	if err != nil {
		return err
	}

	err = acc.checkPassword(password)

	if err != nil {
		acc.generatePassword(10)
	}

	return nil
}

func (acc *accountStruct) checkUrl(urlString *string) error {
	_, err := url.ParseRequestURI(*urlString)

	if err != nil {
		return errors.New("url is invalid")
	}

	return nil
}

func (acc *accountStruct) checkLogin(login *string) error {
	if len(*login) < 3 {
		return errors.New("login is too short")
	}

	return nil
}

func (acc *accountStruct) checkPassword(password *string) error {
	if len(*password) < 3 {
		return errors.New("password is too short")
	}

	return nil
}
