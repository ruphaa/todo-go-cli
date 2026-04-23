package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Interactively add new tasks",
	Long:  `Start an interactive session to add tasks one by one. Press 'q' to quit.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("👋 Hey there! What do you want to add today?")
		fmt.Println("   (type 'q' to quit)")
		fmt.Println()

		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				return
			}

			task := strings.TrimSpace(input)

			if task == "" {
				continue
			}

			if strings.ToLower(task) == "q" {
				fmt.Println("\n✌️ See you later!")
				return
			}

			if err := addTask(task); err != nil {
				fmt.Printf("❌ Oops, couldn't add that: %v\n\n", err)
				continue
			}

			fmt.Printf("✅ Added \"%s\" to your list!\n\n", task)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
