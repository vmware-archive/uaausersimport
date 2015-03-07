package config

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Space struct {
	Roles string
}
type Org struct {
	Roles  string
	Spaces map[string]Space
}

type User struct {
	Uid        string
	Externalid string
	Emails     []string
	Orgs       map[string]Org
}
type Config struct {
	Sysdomain string
	Users     []User
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
