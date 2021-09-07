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
	"text/tabwriter"

	"github.com/rikongha/todo/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list displays tasks in todo list",
	Long:  `list displays all tasks in the todo list.`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listRun(cmd *cobra.Command, args []string) {
	todoList, err := todo.ReadTasks()
	if err != nil {
		log.Printf("%v\n", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 4, 1, 2, ' ', 0)
	for x, todos := range todoList {
		fmt.Fprintln(w, strconv.Itoa(x+1)+". "+todos.Task+"\t"+todos.BeautifyDone())
	}
	w.Flush()
}
