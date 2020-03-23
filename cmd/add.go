/*
Copyright © 2020 Tiger Nie <nhl0819@gmail.com>

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
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tiega/tri/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Long:  `Add will create a new todo item to the list`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a todo item.")
		}
		return nil
	},

	Run: addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {
	// Get items
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	// Add items
	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)

		fmt.Printf("Adding todo \"%s\"\n", x)
		items = append(items, item)
	}

	// Save items
	err = todo.SaveItems(dataFile, items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
