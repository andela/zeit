package cmd

import (
	"github.com/andela/zeit/lib"
	"github.com/andela/zeit/utility"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"html/template"
	"os"
)

var numberOfEntries int

func history(cmd *cobra.Command, args []string) {
	config := lib.NewConfigFromFile()
	filePath := os.ExpandEnv("$HOME/.zeit/history.html")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	var entries []lib.Entry
	startIndex := len(config.Entries) - numberOfEntries
	if startIndex < 0 {
		startIndex = 0
	}
	entriesSlice := config.Entries[startIndex:]
	for _, name := range entriesSlice {
		entry, err := config.GetEntryByName(name)
		if err == nil {
			entries = append(entries, *entry)
		}
	}
	t := template.New("History Template")
	t = t.Funcs(template.FuncMap{"dateformat": utility.FormatToDateTime})

	t = template.Must(t.ParseFiles("tpl/history.html"))

	err = t.ExecuteTemplate(file, "history.html", entries)
	if err != nil {
		panic(err)
	} else {
		open.Run(filePath)
	}
}

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View your entries",
	Long:  `View all your entries in the browser. This defaults to 10 latest entries`,
	Run:   history,
}

func init() {
	historyCmd.Flags().IntVarP(&numberOfEntries, "last", "l", 10, "Specify number of entries to view")
	RootCmd.AddCommand(historyCmd)
}
