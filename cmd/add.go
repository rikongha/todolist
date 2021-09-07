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
	"os"
	"strings"

	"github.com/rikongha/todo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to the todo list",
	Long: `Add will create a new pending todo task.
	
	todo add [some note]`,
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	todoList, err := todo.ReadTasks()
	if err != nil {
		log.Printf("%v\n", err)
	}

	desc := strings.Join(args, " ")
	task := todo.Todo{Task: desc}
	todoList = append(todoList, task)

	err = todo.SaveTasks(todoList)
	if err != nil {
		panic(err)
	}
	fmt.Println("Task successfully added!")
}
