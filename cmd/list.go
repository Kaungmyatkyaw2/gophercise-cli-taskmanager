package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of our tasks.",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Smth went wrong: ", err.Error())
			return
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete. So take a rest.")
			return
		}

		fmt.Println("You have the following tasks: ")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
