package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	must(db.Init(dbPath))

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Errorf("Error: %s\n", err.Error())

		os.Exit(1)
	}
}
