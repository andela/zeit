// Copyright © 2016 IKEM OKONKWO <ikem.okonkwo@andela.com>
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
	"strings"
	"time"

	"github.com/andela/zeit/lib"
	"github.com/kjk/betterguid"
	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking time",
	Long:  `Full tracking description`,
	Run: func(cmd *cobra.Command, tags []string) {
		projectName := cmd.Flag("project").Value.String()
		config := lib.NewConfigFromFile()
		entry := &lib.Entry{
			ID: betterguid.New()[1:],
		}
		err := entry.StartTracking(projectName, tags, config)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			currentTime := au.Cyan(time.Now().Format("15:04"))
			fmt.Printf(
				"Starting Project %s with tags [ %s ] at %s\n",
				projectName,
				strings.Join(tags, " "),
				au.Bold(currentTime),
			)
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
	// Here you will define your flags and configuration settings.
	startCmd.Flags().StringP("project", "p", "", "Name of project to track")
	// config := Config{
	// 	ID:           betterguid.New(),
	// 	Token:        "some strings",
	// 	Name:         "Okonkwo Ikem",
	// 	CurrentEntry: "",
	// 	Projects: []KeyValue{
	// 		KeyValue{ID: betterguid.New(), Name: "core infrastructure"},
	// 		KeyValue{ID: betterguid.New(), Name: "kaizen"},
	// 		KeyValue{ID: betterguid.New(), Name: "allocations"},
	// 		KeyValue{ID: betterguid.New(), Name: "skilltree"},
	// 		KeyValue{ID: betterguid.New(), Name: "path"},
	// 	},
	// 	Tags: []KeyValue{
	// 		KeyValue{ID: betterguid.New(), Name: "core"},
	// 		KeyValue{ID: betterguid.New(), Name: "hack"},
	// 		KeyValue{ID: betterguid.New(), Name: "skill"},
	// 		KeyValue{ID: betterguid.New(), Name: "kaiz"},
	// 		KeyValue{ID: betterguid.New(), Name: "zeit"},
	// 	},
	// 	NewTags: []KeyValue{},
	// }
	// b, _ := json.Marshal(config)
	// ioutil.WriteFile(os.ExpandEnv("$HOME/.zeit/config.json"), b, 0644)
	// startCmd.Flags().String("tag", "t", "New tag of tracked project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
