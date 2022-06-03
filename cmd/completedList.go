package cmd

import (
	"fmt"
	"os"

	"github.com/SkylerBair/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var completedListCmd = &cobra.Command{
	Use:   "completedList",
	Short: "List all of your completed tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllCompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no completed tasks time to get to work! ")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d, %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedListCmd)
}
