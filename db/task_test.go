package db_test

import (
	"path/filepath"
	"testing"

	"github.com/SkylerBair/task/db"
)

func TestDB(t *testing.T) {
	path := filepath.Join(t.TempDir(), "taskdb.test")
	err := db.Init(path)
	if err != nil {
		t.Errorf("failed to get test db: %v", err)
	}
	id, err := db.CreateTask("test task")
	if err != nil {
		t.Errorf("failed to create task: %v", err)
	}
	if id == -1 {
		t.Errorf("failed to create task")
	}
	tasks, err := db.AllTasks()
	if err != nil {
		t.Errorf("failed to get tasks")
	}
	if len(tasks) != 1 {
		t.Errorf("failed to get correct number of tasks")
	}
	err = db.CompleteTask(id)
	if err != nil {
		t.Errorf("failed to mark task as completed.")
	}
	tasks, err = db.AllTasks()
	if err != nil {
		t.Errorf("failed to open task list")
	} else if len(tasks) != 0 {
		t.Errorf("Failed to mark task as completed.")
	}

}
