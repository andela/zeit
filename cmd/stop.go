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
    "io/ioutil"
    "os"
    "os/exec"
    "time"

    "github.com/andela/zeit/lib"
    au "github.com/logrusorgru/aurora"
    "github.com/spf13/cobra"
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
            cleanUpAllRunningNotifications()
        } else {
            fmt.Print("You are not logging time\n")
        }
    },
}

func cleanUpAllRunningNotifications() {
    err := checkScriptExists()
    if err != nil {
        os.Exit(1)
    }
    writeScriptToFile()
    killRunningScript()
}

func writeScriptToFile() {
    bytes, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.zeit/pid.txt"))
    script := fmt.Sprintf("kill -9 %s", string(bytes))
    if err != nil {
        os.Exit(1)
    } else {
        err = ioutil.WriteFile(os.ExpandEnv("$HOME/.zeit/stop.sh"), []byte(script), 0777)
        if err != nil {
            os.Exit(1)
        }
    }
}

func checkScriptExists() error {
    _, err := os.Stat(os.ExpandEnv("$HOME/.zeit/pid.txt"))
    return err
}

func killRunningScript() {
    cmd := exec.Command("nohup", os.ExpandEnv("$HOME/.zeit/stop.sh"))
    err := cmd.Start()
    if err == nil {
        fmt.Println("All timers succesfully closed!")
    }
}

func init() {
    RootCmd.AddCommand(stopCmd)

}
