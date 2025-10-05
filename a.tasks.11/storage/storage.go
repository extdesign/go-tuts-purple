package storage

import (
	"errors"

	"a.tasks.11/files"
)

type Storage struct {
	Filepath string
	Data     []byte
}

func NewStorage(filepath string) (*Storage, error) {
	if filepath == "" {
		return nil, errors.New("filepath cannot be empty")
	}

	return &Storage{
		Filepath: filepath,
		Data:     []byte{},
	}, nil
}

func (storage *Storage) Save() {
	if storage.Data == nil {
		storage.Data = []byte{}
	}

	files.WriteFile(storage.Data, storage.Filepath)
}

func (storage *Storage) Read() error {
	data, err := files.ReadFile(storage.Filepath)

	if err != nil {
		return err
	}

	storage.Data = data
	return nil
}
