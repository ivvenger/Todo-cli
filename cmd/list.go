package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать все задачи",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		tasks, err := storage.All()
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("Список задач пуст!")
			return nil
		}
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
