package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := loadTasks()
		if err != nil {
			fmt.Printf("Error loading tasks: %v\n", err)
			return
		}

		fmt.Printf("📝 Tasks from: %s\n\n", todoPath)

		if len(tasks) == 0 {
			fmt.Println("   (none yet — try `todo new` or `todo add <task>`)")
			return
		}

		for i, task := range tasks {
			fmt.Printf("   %d. %s\n", i+1, task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
