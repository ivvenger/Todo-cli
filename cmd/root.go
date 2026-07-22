package cmd

import (
	"fmt"
	"os"

	"github.com/ivvenger/todo-cli/task"
	"github.com/spf13/cobra"
)

var storage *task.Storage

var dbPath string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Простой менеджер задач в командной строке",
	Long:  "todo — CLI-утилита для управления списком дел: добавление, просмотр, отметка и удаление задач",
	// Ошибки и usage печатаем сами в Execute, чтобы не дублировать вывод.
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		storage = task.NewStorage(dbPath)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbPath, "file", "tasks.json", "путь к файлу с задачами")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}
