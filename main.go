package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SkylerBair/task/cmd"
	"github.com/SkylerBair/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
