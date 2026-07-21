// Package task содержит модель задачи, хранилище и операции над списком дел.
package task

import "errors"

// Sentinel-ошибки пакета. Их удобно проверять через errors.Is,
// в том числе в тестах и в командах CLI.
var (
	// ErrEmptyTitle возвращается, когда пытаются добавить задачу с пустым заголовком.
	ErrEmptyTitle = errors.New("заголовок задачи не может быть пустым")

	// ErrTaskNotFound возвращается, когда задача с указанным ID не найдена.
	ErrTaskNotFound = errors.New("задача не найдена")
)
