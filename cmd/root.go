package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd     = &cobra.Command{
		Use:   "todo",
		Short: "A simple todo CLI",
		Long:  `A longer description of your application.`,
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
		"/Users/ruphaa/Documents/Second-brain/Todo-go.md",
		"Absolute path to the markdown file for storing todos",
	)
}