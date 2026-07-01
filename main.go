package main

import (
	"fmt"

	"github.com/ivvenger/todo-cli/task"
)

func main() {
	storage := task.NewStorage("tasks.json")

	t1, err := storage.Add("убрать дом")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Printf("Добавлена задача %d: %s\n", t1.ID, t1.Title)

	t2, err := storage.Add("сходить на почту")
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Printf("Добавлена задача %d: %s\n", t2.ID, t2.Title)

	if err := storage.Complete(t1.ID); err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Printf("Задача %d выполнена!\n", t1.ID)

	tasks, err := storage.All()
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Println("Список задач:")
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
	}

	if err := storage.Delete(t2.ID); err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	fmt.Printf("Запись %d - удалена!\n", t2.ID)
}