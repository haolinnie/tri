/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tiega/tri/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d", "rm"},
	Short:   "Remove a todo",
	Long: `Done will remove a todo by its index.
	Can be invoked with aliases 'd' or 'rm'.`,

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a todo item")
		}
		return nil
	},

	Run: doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}

	// Get indicesslice from string slice
	indices := []int{}
	for _, s := range args {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Printf("%v", err)
		} else {
			indices = append(indices, i)
		}
	}

	// Sort indices in reverse
	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	// Remove items
	for _, idx := range indices {
		// Correct indexing to start at 0
		idx -= 1

		// Make sure index within range
		if idx >= len(items) {
			fmt.Printf("Index %d out of range, moving on.\n", idx+1)
			continue
		}
		fmt.Printf("Removing item \"%s\"\n", items[idx].Text)

		items = append(items[:idx], items[idx+1:]...)
	}

	// Save items
	todo.SaveItems(dataFile, items)
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
