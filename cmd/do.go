package cmd

import (
	"fmt"

	"github.com/SkylerBair/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids int
		err := db.CompleteTask(ids)
		if err != nil {
			fmt.Println("Error marking task as complete.")
		}
		fmt.Println("marked as complete.")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
