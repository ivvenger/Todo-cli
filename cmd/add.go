// Package cmd содержит команды CLI todo: add, list, done, rm.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [текст задачи]",
	Short: "Добавить новую задачу",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		t, err := storage.Add(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Добавлена задача [%d]: %s\n", t.ID, t.Title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
