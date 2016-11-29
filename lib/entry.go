package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/kjk/betterguid"
)

type KeyValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Entry struct {
	ID          string     `json:"id"`
	ProjectID   string     `json:"project_id"`
	ProjectName string     `json:"project_name"`
	Start       string     `json:"start"`
	Stop        string     `json:"stop"`
	Tags        []KeyValue `json:"tags"`
}

func NewEntryFromFile(id string) *Entry {
	entry := Entry{}
	b, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.zeit/") + id + ".json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &entry)
	if err != nil {
		panic(err)
	}
	return &entry
}

func (entry *Entry) Duration() string {
	start, _ := time.Parse(JavascriptISOString, entry.Start)
	duration := time.Now().Sub(start)
	return fmt.Sprintf("%d hours %d minutes", int(duration.Hours()), int(duration.Minutes()))
}

func (entry *Entry) StopTracking(config *Config) {
	entry.Stop = time.Now().UTC().Format(JavascriptISOString)
	entry.Save()
	config.CurrentEntry = ""
	config.Save()
}

func (entry *Entry) StartTracking(projectName string, tags []string, config *Config) error {
	for _, tagName := range tags {
		if !config.ContainTag(tagName) {
			config.AddNewTag(tagName)
		}
		entry.Tags = append(entry.Tags, KeyValue{ID: betterguid.New(), Name: tagName})
	}

	ok, project := config.ContainProject(projectName)
	if !ok {
		return fmt.Errorf("Project %s does not exist or has not been assigned to you\n", projectName)
	}
	entry.Start = time.Now().UTC().Format(JavascriptISOString)
	entry.ProjectID = project.ID
	entry.ProjectName = project.Name
	entry.Save()
	config.CurrentEntry = entry.ID
	config.Save()
	return nil
}

func (entry *Entry) Save() {
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(os.ExpandEnv("$HOME/.zeit/")+entry.ID+".json", jsonEntry, 0644)
	if err != nil {
		panic(err)
	}
}
