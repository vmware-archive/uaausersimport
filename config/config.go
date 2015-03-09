package config

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Space struct {
	Name  string
	Roles []string
}
type Org struct {
	Name   string
	Roles  []string
	Spaces []Space
}

type User struct {
	Uid        string
	Externalid string
	Emails     []string
	Orgs       []Org
}
type Config struct {
	Users []User
}

func Parse(reader io.Reader) (*Config, error) {
	config := Config{}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	return &config, err
}
