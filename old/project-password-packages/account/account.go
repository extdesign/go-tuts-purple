package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
)

type AccountStruct struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	AccountStruct
}

func (acc *AccountStruct) GeneratePassword(n int) {
	lettersRunes := []rune("abcdefjABCEFJ!_*")
	res := make([]rune, n)

	for i := range n {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString *string) (*AccountWithTimeStamp, error) {
	acc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		AccountStruct: AccountStruct{
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
