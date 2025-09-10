package task

import (
	"cli-task-tracker/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type StorageJSON struct {
	Data        []Task
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

	err = json.Unmarshal(bytes, &storage.Data)

	if err != nil {
		return fmt.Errorf("json decoding error: %w", err)
	}

	storage.GetLen()
	storage.GetLastID()

	return nil
}

func (storage *StorageJSON) WriteData() error {
	bytes, err := json.MarshalIndent(storage.Data, "", "\t")

	if err != nil {
		return fmt.Errorf("data encoding error: %w", err)
	}

	err = os.WriteFile(utils.FILENAME, bytes, 0664)

	if err != nil {
		return fmt.Errorf("error writing to a file: %w", err)
	}

	storage.ReadData()

	return nil
}

func (storage *StorageJSON) GetLastID() int {
	len_slice := storage.GetLen()
	if len_slice == 0 {
		storage.last_id = 0
	} else {
		storage.last_id = storage.Data[storage.entries_num-1].ID
	}
	return storage.last_id
}

func (storage *StorageJSON) GetLen() int {
	storage.entries_num = len(storage.Data)
	return storage.entries_num
}

func (storage *StorageJSON) AddToJSON(task *Task) {
	storage.Data = append(storage.Data, *task)
}

func (storage *StorageJSON) GetTaskByID(id int) (int, error) {
	for i := 0; i < storage.entries_num; i++ {
		if storage.Data[i].ID == id {
			return i, nil
		}
	}
	return -1, errors.New("error: this index does not exist")
}
