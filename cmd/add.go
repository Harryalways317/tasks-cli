package cmd

import (
    "github.com/spf13/cobra"
    "github.com/Harryalways317/tasks/internal/taskmanager"
)

var addCmd = &cobra.Command{
    Use:   "add [description]",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        description := args[0]
        taskmanager.AddTask(description)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
