package lib

import (
	"encoding/json"
	"fmt"
	"github.com/kjk/betterguid"
	"io/ioutil"
	"os"
)

var JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"

type User struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type Config struct {
	CurrentUser  User       `json:"user"`
	CurrentEntry string     `json:"current_entry"`
	Projects     []KeyValue `json:"projects"`
	Tags         []KeyValue `json:"tags"`
	NewTags      []KeyValue `json:"new_tags"`
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
	fmt.Println("Config Saved!")
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

func NewConfigFromFile() *Config {
	config := Config{}
	b, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.zeit/config.json"))
	if err != nil {
		createDirectory()
	} else if err = json.Unmarshal(b, &config); err != nil {
		panic(err)
	}
	return &config
}

func createDirectory() {
	rootPath := os.ExpandEnv("$HOME/.zeit")
	if err := os.MkdirAll(rootPath, 0777); err != nil {
		panic(err)
	} else if _, err := os.Create(rootPath + "/config.json"); err != nil {
		panic(err)
	} else {
		ioutil.WriteFile(rootPath+"/config.json", []byte("{}"), 0777)
	}
}
