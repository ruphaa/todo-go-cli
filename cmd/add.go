package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if err := addTask(task); err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Printf("✅ Added: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
