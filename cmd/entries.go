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
	"os"
	"log"
	"time"
	"regexp"
	"html/template"
	"github.com/spf13/cobra"
	"github.com/skratchdot/open-golang/open"
	"github.com/andela/zeit/utility"
	"github.com/andela/zeit/lib"
)

var start string
var end string
var JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"

func entries(cmd *cobra.Command, args []string) {
	if start != "" && end != "" {
        startDate, endDate := getStartAndEndDate(start, end)
        entriesRange := getEntriesRange(startDate, endDate)
        previewEntries(entriesRange)
    } else{
        log.Fatalf("You must specify both a start and an end date flags")
    }
}

func getEntriesRange(start time.Time, end time.Time) []lib.Entry {
	config := lib.NewConfigFromFile()
	var entries []lib.Entry
	for _, entryName := range config.Entries {
		entry, err := config.GetEntryByName(entryName)
		entryStartTime, _ := time.Parse(JavascriptISOString, entry.Start)
		if err == nil && (entryStartTime.Equal(start) || entryStartTime.After(start)) && (entryStartTime.Equal(end) || entryStartTime.Before(end)) {
			entries = append(entries, *entry)
		}
	}
	return entries
}

func previewEntries(entries []lib.Entry) {
	filePath, file := getOrCreateFile()
	t := template.New("History Template")
	t = t.Funcs(template.FuncMap{"dateformat": utility.FormatToDateTime})

	t = template.Must(t.ParseFiles("tpl/history.html"))

	err := t.ExecuteTemplate(file, "history.html", entries)
	if err != nil {
		panic(err)
	} else {
		open.Run(filePath)
	}
}

func getOrCreateFile() (string, *os.File) {
	filePath := os.ExpandEnv("$HOME/.zeit/history.html")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return filePath, file
}



func getStartAndEndDate(args ...string) (time.Time, time.Time) {
	var dateHolder []time.Time
	patt1 := regexp.MustCompile("(\\w{3})\\s(\\w{2})$") // Match 'JAN 02'
	patt2 := regexp.MustCompile("(\\w{3})\\s(\\w{4})$") // Match 'JAN 2016'
	patt3 := regexp.MustCompile("(\\w{3})\\s(\\w{2})\\s(\\w{4})$") // Match 'JAN 02 2016'
	patt4 := regexp.MustCompile("(\\d{2})-(\\d{2})-(\\d{4})$") // Match '01-08-3016'
	patt5 := regexp.MustCompile("(\\d{2})\\/(\\d{2})\\/(\\d{4})$") // Match '01/08/2016'
	for _, arg := range args {
		switch {
			case len(patt1.FindStringSubmatch(arg)) > 0:
				arg = fmt.Sprintf("%s %v", arg, time.Now().Year())
				t, _ := time.Parse("Jan 02 2006", arg)
				dateHolder = append(dateHolder, t)
			case len(patt2.FindStringSubmatch(arg)) > 0:
				t, _ := time.Parse("Jan 2006", arg)
				dateHolder = append(dateHolder, t)
			case len(patt3.FindStringSubmatch(arg)) > 0:
				t, _ := time.Parse("Jan 02 2006", arg)
				dateHolder = append(dateHolder, t)
			case len(patt4.FindStringSubmatch(arg)) > 0:
				t, _ := time.Parse("02-01-06", arg)
				dateHolder = append(dateHolder, t)
			case len(patt5.FindStringSubmatch(arg)) > 0:
				t, _ := time.Parse("02/01/06", arg)
				dateHolder = append(dateHolder, t)
			default:
				log.Fatalf("Date %s did not match any of the formats, please check valid formats using 'zeit entries --help'", arg);
		}
	}
	return dateHolder[0], dateHolder[1]
}

var entryCmd = &cobra.Command{
	Use:   "entries",
	Short: "View all entries matching specified range",
	Long: `View all entries matching the specified range specified by --start and --end
USAGE

zeit entries --start [start date] --end [stop date] E.g 

zeit entries --start 'JAN 02' --end 'NOV 03 2016' OR		
zeit entries --start 'JAN 2016' --end 'NOV 2016' OR
zeit entries --start 'JAN 02 2016' --end 'NOV 30 2016' OR
zeit entries --start '01-08-2016' --end '02-09-2016' OR
zeit entries --start '01/08/2016' --end '02/09/2016' OR
	`,
	Run: entries,
}

func init() {
	entryCmd.Flags().StringVarP(&start, "start", "s", "", "Specify the start date")
	entryCmd.Flags().StringVarP(&end, "end", "e", "", "Specify the end date")
	RootCmd.AddCommand(entryCmd)
}
