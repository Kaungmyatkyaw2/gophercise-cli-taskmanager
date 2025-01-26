package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete ",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)

			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)

			}
		}

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}
			tasks := tasks[id-1]
			err := db.DeleteTask(tasks.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err.Error())
			} else {
				fmt.Printf("Mark \"%d\" as completed.\n", id)
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
