/*
Copyright © 2021 R. Ikongha <rikongha@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/rikongha/todo/todo"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:     "cleanup",
	Aliases: []string{"c"},
	Short:   "Cleanup removes completed tasks",
	Long:    `Cleanup removes all tasks that have been marked as completed.`,
	Run:     cleanUpRun,
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}

func cleanUpRun(cmd *cobra.Command, args []string) {
	todoList, err := todo.ReadTasks()
	if err != nil {
		log.Printf("%v\n", err)
	}
	count := 0
	for i, todos := range todoList {
		//remove todos marked as done
		if todos.Done {
			copy(todoList[i:], todoList[i+1:])    // Shift a[i+1:] left one index.
			todoList = todoList[:len(todoList)-1] // Truncate slice.
			count++
		}
	}
	err = todo.SaveTasks(todoList)
	if err != nil {
		log.Fatalf("Save items : %v\ns", err)
	}
	if count < 1 {
		fmt.Println("No task has been completed!")
	} else {
		fmt.Println("Todolist successfully cleaned.")
	}
}
