package storage

import (
	"a.tasks.11/files"
)

type Storage struct {
	Data []byte
	db   files.JsonDb
}

func NewStorage(db *files.JsonDb) (*Storage, error) {
	return &Storage{
		Data: []byte{},
		db:   *db,
	}, nil
}

func (storage *Storage) Save() {
	if storage.Data == nil {
		storage.Data = []byte{}
	}

	storage.db.Write(storage.Data)
}

func (storage *Storage) Read() error {
	data, err := storage.db.Read()

	if err != nil {
		return err
	}

	storage.Data = data
	return nil
}
