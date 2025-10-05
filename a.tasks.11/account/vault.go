package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"a.tasks.11/files"
	"a.tasks.11/output"
)

type Vault struct {
	Accounts  []AccountStruct `json:"account"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

var valutFilename string = "data.json"

func NewVault() *Vault {
	file, err := files.ReadFile(valutFilename)

	if err != nil {
		return &Vault{
			Accounts:  []AccountStruct{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nОшибка чтения json данных: %s\n\n", err.Error()))
		return nil
	}

	return &vault
}

func (vault *Vault) AddAccount(acc AccountStruct) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []AccountStruct {
	var result []AccountStruct

	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			result = append(result, acc)
		}
	}

	return result
}

func (vault *Vault) RemoveAccountByUrl(url string) error {
	var accounts []AccountStruct
	isDeleted := false

	for _, acc := range vault.Accounts {
		if !strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
			continue
		}

		isDeleted = true
	}

	if isDeleted {
		vault.Accounts = accounts
		vault.save()
	}

	return errors.New("нет аккаунтов с таким URL")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	json, err := json.MarshalIndent(vault, "", "  ")

	if err != nil {
		return nil, err
	}

	return json, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось преобразовать: %s\n\n", err.Error()))
		return
	}

	files.WriteFile(data, valutFilename)
}
