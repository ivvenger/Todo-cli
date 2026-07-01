package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ivvenger/todo-cli/task"
)

func main() {
	storage := task.NewStorage("tasks.json")

	if len(os.Args) < 2 {
		fmt.Println("Укажите команду: add, list, done или rm")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Укажите текст задачи: todo add \"текст\"")
			os.Exit(1)
		}
		title := os.Args[2]
		t, err := storage.Add(title)
		if err != nil {
			fmt.Println("Ошибка при добавлении:", err)
			os.Exit(1)
		}
		fmt.Printf("Добавлена задача [%d]: %s\n", t.ID, t.Title)
	case "list":
		tasks, err := storage.All()
		if err != nil {
			fmt.Println("Ошибка при чтении списка:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("Список задач пуст")
			return
		}
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Укажите ID задачи: todo done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			os.Exit(1)
		}
		if err := storage.Complete(id); err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		fmt.Printf("Задача %d отмечена выполненной!\n", id)
	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Укажите ID задачи: todo rm <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			os.Exit(1)
		}
		if err := storage.Delete(id); err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		fmt.Printf("Задача %d удалена!\n", id)
	default:
		fmt.Printf("Неизвестная команда: %s\n", command)
		os.Exit(1)
	}
}