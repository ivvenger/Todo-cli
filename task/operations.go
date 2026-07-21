package task

import (
	"fmt"
	"strings"
	"time"
)

// Add добавляет новую задачу в хранилище и возвращает её.
// Пустой (или состоящий из пробелов) заголовок отклоняется.
func (s *Storage) Add(title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	tasks, err := s.Load()
	if err != nil {
		return Task{}, err
	}

	newTask := Task{
		ID:        nextID(tasks),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := s.Save(tasks); err != nil {
		return Task{}, err
	}

	return newTask, nil
}

// Complete помечает задачу с указанным ID выполненной.
func (s *Storage) Complete(id int) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	i := findIndex(tasks, id)
	if i < 0 {
		return fmt.Errorf("id %d: %w", id, ErrTaskNotFound)
	}

	tasks[i].Done = true
	return s.Save(tasks)
}

// Delete удаляет задачу с указанным ID.
func (s *Storage) Delete(id int) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	i := findIndex(tasks, id)
	if i < 0 {
		return fmt.Errorf("id %d: %w", id, ErrTaskNotFound)
	}

	tasks = append(tasks[:i], tasks[i+1:]...)
	return s.Save(tasks)
}

// All возвращает все задачи из хранилища.
func (s *Storage) All() ([]Task, error) {
	return s.Load()
}

// --- Чистые функции: без обращения к файлу, легко тестируются ---

// nextID возвращает следующий свободный ID (максимальный + 1).
func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// findIndex возвращает индекс задачи с данным ID или -1, если не найдена.
func findIndex(tasks []Task, id int) int {
	for i := range tasks {
		if tasks[i].ID == id {
			return i
		}
	}
	return -1
}

// FilterByStatus возвращает задачи, у которых поле Done равно done.
func FilterByStatus(tasks []Task, done bool) []Task {
	result := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		if t.Done == done {
			result = append(result, t)
		}
	}
	return result
}
