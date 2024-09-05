package cmd

import (
    "fmt"
	"time"
    "github.com/Harryalways317/tasks/internal/taskmanager"
    "github.com/spf13/cobra"
)

var allTasks bool

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks := taskmanager.ListTasks(allTasks)
        for _, task := range tasks {
            fmt.Printf("ID: %d, Description: %s, Created At: %s, Done: %t\n",
                task.ID, task.Description, task.CreatedAt.Format(time.RFC1123), task.IsComplete)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
    listCmd.Flags().BoolVarP(&allTasks, "all", "a", false, "Show all tasks, including completed")
}
