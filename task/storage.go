package task

import (
	"encoding/json"
	"errors"
	"os"
)

// Storage читает и пишет список задач в JSON-файл по пути path.
type Storage struct {
	path string
}

// NewStorage создаёт хранилище, работающее с файлом по пути path.
func NewStorage(path string) *Storage {
	return &Storage{path: path}
}

// Load читает задачи из файла. Если файла ещё нет, возвращает пустой список
// без ошибки — это нормальная ситуация при первом запуске.
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

// Save сохраняет список задач в файл в формате JSON с отступами.
func (s *Storage) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.path, data, 0644)
}
