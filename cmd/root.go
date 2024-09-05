package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "tasks",
    Short: "Tasks is a CLI task manager",
}

// Execute executes the root command.
func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
    // Here you can define global flags
}
