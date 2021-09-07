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
	"strconv"

	"github.com/rikongha/todo/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d"},
	Short:   "Done marks a task as completed",
	Long: `Done marks a task as completed.

You can input the index of the todo to mark it as done and 
use undone to undo tasks you are yet to complete.`,
	Run: doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)

	//flags
	// doneCmd.Flags().BoolVar(&markAll, "all", false, "Mark all todos as done")
}

func doneRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	todoList, err := todo.ReadTasks()
	if err != nil {
		log.Printf("%v\n", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid todo index", err)
	}
	if i > 0 && i < len(todoList) {
		todoList[i-1].Done = true
		fmt.Printf("'%s' %v\n", todoList[i-1].Task, "marked done")
		err = todo.SaveTasks(todoList)
		if err != nil {
			log.Fatalf("Save items : %v\ns", err)
		}
	} else {
		log.Println(i, "index does not match match any entry")
	}
}
