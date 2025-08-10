package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/edmondandy/go_project/src/cli-task-manager/cmd"
	"github.com/edmondandy/go_project/src/cli-task-manager/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	// fmt.Println("db init worked")
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}