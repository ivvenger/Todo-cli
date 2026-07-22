package task

import "errors"

var (
	ErrEmptyTitle = errors.New("Заголовок задачи не может быть пустым!")
	ErrTaskNotFound = errors.New("Задача не найдена!")
)
