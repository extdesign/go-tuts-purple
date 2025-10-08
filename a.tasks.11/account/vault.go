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

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func NewVault(db *files.JsonDb) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountStruct{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nОшибка чтения json данных: %s\n\n", err.Error()))

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []AccountStruct{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    *db,
	}
}

func (vault *VaultWithDb) AddAccount(acc AccountStruct) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []AccountStruct {
	var result []AccountStruct

	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			result = append(result, acc)
		}
	}

	return result
}

func (vault *VaultWithDb) RemoveAccountByUrl(url string) error {
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
	jsonBytes, err := json.MarshalIndent(vault, "", "  ")

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось преобразовать: %s\n\n", err.Error()))
		return
	}

	err = vault.db.Write(data)

	if err != nil {
		output.PrintRed(fmt.Sprintf("Ошибка: %s\n\n", err.Error()))
	}
}
