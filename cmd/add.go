package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to list ",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Smth went wrong: ", err.Error())
			return
		}

		fmt.Printf("Added \"%s\" to task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
