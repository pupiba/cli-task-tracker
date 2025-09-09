package task

import (
	"cli-task-tracker/pkg/utils"
	"encoding/json"
	"fmt"
	"os"
)

type StorageJSON struct {
	data        []Task
	entries_num int
	last_id     int
}

func NewStorageJSON() *StorageJSON {
	return &StorageJSON{}
}

func (storage *StorageJSON) ReadData() error {
	bytes, err := os.ReadFile(utils.FILENAME)

	if err != nil {
		return fmt.Errorf("file opening error %w", err)
	}

	err = json.Unmarshal(bytes, &storage.data)

	if err != nil {
		return fmt.Errorf("json decoding error %w", err)
	}

	storage.GetLastID()
	storage.GetLen()

	return nil
}

// func (storage *StorageJSON) WriteData() {

// }

func (storage *StorageJSON) GetLastID() int {
	storage.last_id = storage.data[storage.entries_num].ID
	return storage.last_id
}

func (storage *StorageJSON) GetLen() int {
	storage.entries_num = len(storage.data)
	return storage.entries_num
}
