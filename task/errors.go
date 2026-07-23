package task

import "errors"

var (
	ErrEmptyTitle   = errors.New("заголовок задачи не может быть пустым")
	ErrTaskNotFound = errors.New("задача не найдена")
)
