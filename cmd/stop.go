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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "End a zeit session",
	Long:  `The zeit session will be end and the duration of time logged will be recorded`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.NewConfigFromFile()

		if id := config.CurrentEntry; id != "" {
			entry := lib.NewEntryFromFile(id)
			entry.StopTracking(config)
			tags := " "
			for _, tag := range entry.Tags {
				tags += tag.Name + " "
			}
			fmt.Printf(
				"Stopping Project %s with tags [%s] at %s - Duration %s\n",
				entry.ProjectName,
				tags,
				au.Bold(au.Green(fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute()))),
				entry.Duration(),
			)
		} else {
			fmt.Print("You are not logging time\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
