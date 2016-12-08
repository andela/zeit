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
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"log"
	"github.com/andela/zeit/lib"
	"github.com/kjk/betterguid"
	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var notify bool
var interval int

func startFunction(cmd *cobra.Command, tags []string) {
	projectName := cmd.Flag("project").Value.String()
	config := lib.NewConfigFromFile()
	entry := &lib.Entry{
		ID: betterguid.New()[1:],
	}
	err := entry.StartTracking(projectName, tags, config)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		currentTime := au.Cyan(time.Now().Format("15:04"))
		fmt.Printf(
			"Starting Project %s with tags [ %s ] at %s\n",
			projectName,
			strings.Join(tags, " "),
			au.Bold(currentTime),
		)

		if notify {
			startNotificationTimer()
		}
	}
}

func startNotificationTimer() {
	_, err := getBytesFromScript()
	if err != nil {
		writeToScript()
	}
	startScript()
}

func getBytesFromScript() (os.FileInfo, error) {
	return os.Stat(os.ExpandEnv("$HOME/.zeit/notify.sh"))
}

func writeToScript() {
	workDir, _ := os.Getwd()
	workDir = workDir + "/notify.sh"
	bytes, err := ioutil.ReadFile(workDir)
	if err != nil {
		panic(err)
	} else {
		err := ioutil.WriteFile(os.ExpandEnv("$HOME/.zeit/notify.sh"), bytes, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func startScript() {
	cmd := exec.Command("nohup", os.ExpandEnv("$HOME/.zeit/notify.sh"), strconv.Itoa(interval))
	err := cmd.Start()
	if err == nil {
		processId := cmd.Process.Pid
		storeProcessId(processId)
	}
}

func storeProcessId(processId int) {
	bytes := []byte(strconv.Itoa(processId))
	ioutil.WriteFile(os.ExpandEnv("$HOME/.zeit/pid.txt"), bytes, 0777)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking time",
	Long:  `Full tracking description`,
	Run:   startFunction,
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("project", "p", "", "Name of project to track")
	startCmd.Flags().BoolVarP(&notify, "notify", "n", false, "On/Off Notifications")
	startCmd.Flags().IntVarP(&interval, "interval", "i", 30, "Specify Interval in Minutes default 30 minutes")
}
