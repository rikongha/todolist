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

// undoneCmd represents the undone command
var undoneCmd = &cobra.Command{
	Use:     "undone",
	Aliases: []string{"u"},
	Short:   "undone marks a task as not done",
	Long:    `undone marks the specified task as not done.`,
	Run:     undoneRun,
}

func init() {
	rootCmd.AddCommand(undoneCmd)
}

func undoneRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	todoList, err := todo.ReadTasks()
	if err != nil {
		log.Printf("%v\n", err)
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid todo index", err)
	}
	if index > 0 && index < len(todoList) {
		todoList[index-1].Done = false
		fmt.Printf("'%s' %v\n", todoList[index-1].Task, "marked as not done")
		err = todo.SaveTasks(todoList)
		if err != nil {
			log.Fatalf("Save items : %v\ns", err)
		}
	} else {
		log.Println(index, "index does not match match any entry")
	}
}
