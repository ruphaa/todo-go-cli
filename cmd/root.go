package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:              "todo",
		Short:            "A simple todo CLI",
		Long:             `A longer description of your application.`,
		PersistentPreRun: resolveTodoPath,
	}

	todoPath string
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&todoPath,
		"path",
		"p",
		"",
		"Absolute path to the markdown file for storing todos",
	)
}

func resolveTodoPath(cmd *cobra.Command, args []string) {
	if todoPath != "" {
		return
	}

	if envPath := os.Getenv("TODO_PATH"); envPath != "" {
		todoPath = envPath
		return
	}

	home, err := os.UserHomeDir()
	if err == nil {
		todoPath = filepath.Join(home, "Todo.md")
	} else {
		todoPath = "./Todo.md"
	}
}