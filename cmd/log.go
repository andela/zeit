package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
    "log"
    "time"
    "github.com/andela/zeit/lib"
    "regexp"
    "errors"
    "strconv"
    "github.com/kjk/betterguid"
    "os/exec"
)

var (
    today bool
    yesterday bool
    projectName string
    date string
)

func logHours(cmd *cobra.Command, args []string) {
    if len(args) > 0 && len(args) < 2 {
        editEntry(args)
    } else {
        log.Fatal("You can only specify one entry for hours worked")
    }
}

func createEntry(duration string, config *lib.Config, t time.Time) {
    ok, project := config.ContainProject(projectName)
    if !ok {
        log.Fatalf("You cannot log time for %v", projectName)
    }
    entry := &lib.Entry{
        ID: betterguid.New()[1:],
        WorkDuration: duration,
        ProjectID: project.ID,
        ProjectName: project.Name,
        Deleted: false,
        Start: t.Format(JavascriptISOString), //Must specify basic entry
    }
    entry.Save()
    config.AddEntry(entry)
    config.Save()
}

func editEntry(args []string){
    config := lib.NewConfigFromFile()
    _, err := strconv.ParseFloat(args[0], 64)
    if err == nil {
        t, _ := getDateFromArgs()
        entries := config.GetEntriesByDate(t)
        deleteEntries(entries)
        createEntry(args[0], config, t)
    } else {
        log.Fatal("You must enter the time logged as a float. See 'zeit log --help for more info'")
    }
}

func deleteEntries(entries []lib.Entry) {
    removeQuery := "rm "
    for _, entry := range entries {
        entry.Delete()
        removeQuery = removeQuery+ " $HOME/.zeit/" + entry.ID +".json"
    }
    exec.Command(removeQuery).Run()
}

func getDateFromArgs() (time.Time, error) {
    if date != "" {
        return getDateTime()
    } else {
        switch {
        case today:
            return time.Now().UTC(), nil
        case yesterday:
            return time.Now().AddDate(0, 0, -1), nil
        default:
            log.Fatal("You must specify a date to log time. Use zeit log --help for more info")
        }
    }
    return time.Time{}, errors.New("You must specify a date to log time. Use zeit log --help for more info")
}

func getDateTime() (time.Time, error) {
    patt1 := regexp.MustCompile("\\w{3}\\s\\w{2}$")            // Match 'JAN 02'
    patt2 := regexp.MustCompile("\\w{3}\\s\\w{4}$")            // Match 'JAN 2016'
    patt3 := regexp.MustCompile("\\w{3}\\s\\w{2}\\s\\w{4}$") // Match 'JAN 02 2016'
    patt4 := regexp.MustCompile("\\d{2}-\\d{2}-\\d{4}$")     // Match '01-08-3016'
    patt5 := regexp.MustCompile("\\d{2}\\/\\d{2}\\/(\\d{4}$") // Match '01/08/2016'
    switch {
        case patt1.MatchString(date):
            date = fmt.Sprintf("%s %v", date, time.Now().Year())
            return time.Parse("Jan 02 2006", date)
        case patt2.MatchString(date):
            return time.Parse("Jan 2006", date)
        case patt3.MatchString(date):
            return time.Parse("Jan 02 2006", date)
        case patt4.MatchString(date):
            return time.Parse("02-01-06", date)
        case patt5.MatchString(date):
            return time.Parse("02/01/06", date)
        default:
            log.Fatalf("Date %s did not match any of the formats, Check valid formats using 'zeit log--help'", date)
    }
    return time.Time{}, errors.New("You must specify a date to log time. Use 'zeit log --help' for more info")
}




var logCmd = &cobra.Command{
    Use:   "log",
    Short: "Manually Log your time",
    Long:  `Manually Log your time at a specified date. See usage below:

Examples

zeit log [--today] [--yesterday] [--date arg] [--project ProjectName] numberofHours E.g

zeit log --today --project Zhisi 10.00
zeit log --yesterday --project Zhisi 9.30
zeit log --date '01-02-16' --project Skilltree 8.20
    `,
    Run: logHours,
}

func init() {
    logCmd.Flags().BoolVarP(&today, "today", "t", false, "Log time for today")
    logCmd.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "Log time for yesterday")
    logCmd.Flags().StringVarP(&date, "date", "d", "", "Log time at a specified date")
    logCmd.Flags().StringVarP(&projectName, "project", "p", "", "Specify Project to Log time for")
    RootCmd.AddCommand(logCmd)
}
