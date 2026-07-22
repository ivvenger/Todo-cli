package task

import (
	"fmt"
	"strings"
	"time"
)

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

func (s *Storage) All() ([]Task, error) {
	return s.Load()
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func findIndex(tasks []Task, id int) int {
	for i := range tasks {
		if tasks[i].ID == id {
			return i
		}
	}
	return -1
}

func FilterByStatus(tasks []Task, done bool) []Task {
	result := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		if t.Done == done {
			result = append(result, t)
		}
	}
	return result
}
