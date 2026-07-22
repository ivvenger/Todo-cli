package task

import (
	"encoding/json"
	"errors"
	"os"
)

type Storage struct {
	path string
}

func NewStorage(path string) *Storage {
	return &Storage{path: path}
}

func (s *Storage) Load() ([]Task, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Storage) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}
