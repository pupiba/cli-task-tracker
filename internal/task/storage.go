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

func NewStorageJSON() (*StorageJSON, error) {
	storage := &StorageJSON{}
	if err := storage.ReadData(); err != nil {
		return &StorageJSON{}, err
	}
	return storage, nil
}

func (storage *StorageJSON) ReadData() error {
	bytes, err := os.ReadFile(utils.FILENAME)

	if err != nil {
		return fmt.Errorf("file opening error: %w", err)
	}

	err = json.Unmarshal(bytes, &storage.data)

	if err != nil {
		return fmt.Errorf("json decoding error: %w", err)
	}

	storage.GetLen()
	storage.GetLastID()

	return nil
}

func (storage *StorageJSON) WriteData() error {
	bytes, err := json.MarshalIndent(storage.data, "", "\t")

	if err != nil {
		return fmt.Errorf("data encoding error: %w", err)
	}

	err = os.WriteFile(utils.FILENAME, bytes, 0664)

	if err != nil {
		return fmt.Errorf("error writing to a file: %w", err)
	}

	return nil
}

func (storage *StorageJSON) GetLastID() int {
	storage.last_id = storage.data[storage.entries_num-1].ID
	return storage.last_id
}

func (storage *StorageJSON) GetLen() int {
	storage.entries_num = len(storage.data)
	return storage.entries_num
}

func (storage *StorageJSON) AddToJSON(task *Task) {
	storage.data = append(storage.data, *task)
}
