package task

import (
	"fmt"
	"time"
)

func (s *Storage) Add(title string) (Task, error) {
	tasks, err := s.Load()
	if err != nil {
		return Task{}, err
	} 
	
	newTask := Task{
		ID:			nextID(tasks),
		Title:		title,
		Done:		false,
		CreatedAt:	time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := s.Save(tasks); err != nil {
		return Task{}, err
	}

	return newTask, nil
}


func nextID(tasks []Task) int{
	maxID := 0
	for _, t := range(tasks) {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}


func (s *Storage) Complete(id int) error {
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true;
			return s.Save(tasks)
		}
	}

	return fmt.Errorf("Задача с ID %d не найдена", id)
}


func (s *Storage) Delete(id int) error{
	tasks, err := s.Load()
	if err != nil {
		return err
	}

	newTasks := make([]Task, 0, len(tasks))
	found := false 

	for _, t := range tasks {
		if t.ID == id {
			found = true 
			continue 
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return fmt.Errorf("Задача с ID %d не найдена", id)
	}

	return s.Save(newTasks)
}


func (s *Storage) All() ([]Task, error) {
	return s.Load()
}

