// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"time"

	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/andela/zeit/lib"

)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows how much time you have logged since you started logging the current time",
	Long:  `A very log description`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.NewConfigFromFile()

		if id := config.CurrentEntry; id != "" {
			entry := lib.NewEntryFromFile(id)
			currentTime := au.Green(fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute()))
			fmt.Printf(
				"As at %s, you have logged %s for Project %s\n",
				au.Bold(currentTime),
				entry.Duration(),
				entry.ProjectName,
			)
		} else {
			fmt.Print("You are not logging time\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
