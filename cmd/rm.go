package cmd

import (
	"fmt"
	"strconv"

	"github.com/SkylerBair/task/db"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a task from your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong;", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to delete \"%d\". Error: %s\n", id, err)
			} else {
				fmt.Printf("Deleted list item #\"%d\".\n", id)
			}
		}

		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}
