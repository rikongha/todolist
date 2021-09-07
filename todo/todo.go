package todo

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

type Todo struct {
	Task string `csv:"Task"`
	Done bool   `csv:"Done"`
}

const (
	fileName = "todolist.csv"
)

// get filepath of csv
func getFilepath() (string, error) {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, fileName)
		_, err = os.Stat(filename)
		if err == nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		filename = filepath.Join(dirname, fileName)
	}
	return filename, nil
}

// Write to CSV file
func SaveTasks(todoList []Todo) error {
	bytes, err := gocsv.MarshalBytes(todoList)
	if err != nil {
		return err
	}

	filename, err := getFilepath()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Read from CSV file
func ReadTasks() ([]Todo, error) {
	filename, err := getFilepath()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = gocsv.UnmarshalBytes(bytes, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Beautify mark as done
func (t *Todo) BeautifyDone() string {
	if t.Done {
		return "[X]"
	} else {
		return "[ ]"
	}
}
