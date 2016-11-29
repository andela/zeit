package lib

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/kjk/betterguid"
)

var JavascriptISOString = "2006-01-02T15:04:05.999Z07:00"

type Config struct {
	ID           string     `json:"id"`
	Token        string     `json:"token"`
	Name         string     `json:"name"`
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
		panic(err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
