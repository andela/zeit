package lib

import (
    "encoding/json"
    "github.com/kjk/betterguid"
    "io/ioutil"
    "os"
    "time"
    "github.com/andela/zeit/utility"
)

var JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"

type User struct {
    Id    string `json:"id"`
    Token string `json:"token"`
}

type BasicEntry struct {
    Id string
    Date string
}

type Config struct {
    CurrentUser  User       `json:"user"`
    CurrentEntry string     `json:"current_entry"`
    Projects     []KeyValue `json:"projects"`
    Tags         []KeyValue `json:"tags"`
    NewTags      []KeyValue `json:"new_tags"`
    Entries      []BasicEntry `json:"entries"`
}

func (c *Config) Save() {
    b, err := json.Marshal(c)
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile(os.ExpandEnv("$HOME/.zeit/config.json"), b, 0644)
    if err != nil {
        panic(err)
    }
}

func (c *Config) ContainTag(name string) bool {
    contain := false
    for _, tag := range c.Tags {
        if tag.Name == name {
            contain = true
            break
        }
    }
    return contain
}

func (c *Config) AddNewTag(name string) {
    tag := KeyValue{ID: betterguid.New(), Name: name}
    c.NewTags = append(c.NewTags, tag)
}

func (c *Config) ContainProject(name string) (bool, KeyValue) {
    contains := false
    var project KeyValue
    for _, p := range c.Projects {
        if p.Name == name {
            project = p
            contains = true
            break
        }
    }

    return contains, project
}

func (c *Config) AddEntry(entry *Entry) {
    basicEntry := BasicEntry{
        Id: entry.ID,
        Date: entry.Start,
    }
    c.Entries = append(c.Entries, basicEntry)
}

func (c *Config) GetCurrentEntry() (*Entry, error) {
    entry := &Entry{}
    if c.CurrentEntry == "" {
        return nil, nil
    }
    currentEntryPath := os.ExpandEnv("$HOME/.zeit/" + c.CurrentEntry + ".json")
    bytes, err := ioutil.ReadFile(currentEntryPath)
    if err != nil {
        return nil, err
    } else {
        err = json.Unmarshal(bytes, entry)
    }
    return entry, err
}

func (c *Config) GetEntry(entryName string) (*Entry, error) {
    entry := &Entry{}
    currentEntryPath := os.ExpandEnv("$HOME/.zeit/" + entryName + ".json")
    bytes, err := ioutil.ReadFile(currentEntryPath)
    if err != nil {
        return entry, err
    } else {
        err = json.Unmarshal(bytes, entry)
    }
    return entry, err
}

func (c *Config) GetEntryByName(entryName string) (*Entry, error) {
    return c.GetEntry(entryName)
}

func (c *Config) GetEntriesByDate(entryDate time.Time) []Entry {
    var entries []Entry
    for _, basicEntry := range c.Entries {
        _entryDate, _ := time.Parse(JavascriptISOString, basicEntry.Date)
        if utility.EqualDates(entryDate, _entryDate) {
            entry, _ := c.GetEntry(basicEntry.Id)
            entries = append(entries, *entry)
        }
    }
    return entries
}


func NewConfigFromFile() *Config {
    config := Config{}
    b, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.zeit/config.json"))
    if err != nil {
        config := getMockConfig()
        createDirectory(config)
        return config
    } else if err = json.Unmarshal(b, &config); err != nil {
        panic(err)
    }
    return &config
}

func createDirectory(config *Config) {
    rootPath := os.ExpandEnv("$HOME/.zeit")
    if err := os.MkdirAll(rootPath, 0777); err != nil {
        panic(err)
    } else {
        bytes, err := json.Marshal(config)
        if err != nil {
            panic(err)
        } else {
            ioutil.WriteFile(rootPath+"/config.json", bytes, 0777)
        }
    }
}
