package main

import (
	"fmt"
	"github.com/Badara-Senpai/cli_task_manager/task/cmd"
	"github.com/Badara-Senpai/cli_task_manager/task/db"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")

	detect(db.Init(dbPath))
	detect(cmd.RootCmd.Execute())
}

func detect(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
