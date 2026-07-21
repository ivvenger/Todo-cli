package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [id]",
	Short: "Удалить задачу",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("id должен быть числом, получено %q", args[0])
		}
		if err := storage.Delete(id); err != nil {
			return err
		}
		fmt.Printf("Задача [%d] удалена!\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
