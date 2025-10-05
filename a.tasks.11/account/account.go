package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"a.tasks.11/output"
)

type AccountStruct struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (acc *AccountStruct) GeneratePassword(n int) {
	lettersRunes := []rune("abcdefjABCEFJ!_*")
	res := make([]rune, n)

	for i := range n {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	acc.Password = string(res)
}

func NewAccount(login, password, urlString *string) (*AccountStruct, error) {

	acc := &AccountStruct{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     *login,
		Password:  *password,
		Url:       *urlString,
	}

	err := acc.validateArguments(login, password, urlString)

	if err != nil {
		return nil, err
	}

	acc.Login = *login
	acc.Password = *password
	acc.Url = *urlString

	return acc, nil
}

func (acc *AccountStruct) validateArguments(login, password, urlString *string) error {
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
		acc.GeneratePassword(10)
	}

	return nil
}

func (acc *AccountStruct) checkUrl(urlString *string) error {
	_, err := url.ParseRequestURI(*urlString)

	if err != nil {
		return errors.New("url is invalid")
	}

	return nil
}

func (acc *AccountStruct) checkLogin(login *string) error {
	if len(*login) < 3 {
		return errors.New("login is too short")
	}

	return nil
}

func (acc *AccountStruct) checkPassword(password *string) error {
	if len(*password) < 3 {
		return errors.New("password is too short")
	}

	return nil
}

func (acc *AccountStruct) Output() {
	s := fmt.Sprintf("\nLogin: %s\nPassword: %s\nUrl: %s\n\n", acc.Login, acc.Password, acc.Url)
	output.PrintYellow(s)
}
