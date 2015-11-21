package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	uaa "github.com/pivotalservices/uaausersimport/adduser"
	cc "github.com/pivotalservices/uaausersimport/cloudcontroller"
	config "github.com/pivotalservices/uaausersimport/config"
	. "github.com/pivotalservices/uaausersimport/functions"
	token "github.com/pivotalservices/uaausersimport/token"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() (err error) {
	users := os.Getenv("USERS_CONFIG_FILE")
	file, err := os.Open(users)
	if err != nil {
		err = errors.New(fmt.Sprintf("Can not open %s : %s", users, err.Error()))
		return
	}
	cfg, err := config.Parse(file)
	if err != nil {
		return
	}
	info, err := parseEnv()
	if err != nil {
		return
	}
	info.RequestFn = token.RequestWithToken
	fun := token.GetToken.MapUsers(*cfg).AddUaaUser(uaa.Adduser).AddCCUser(cc.Adduser).MapOrgs(cc.AssociateOrg).MapSpaces(cc.AssociateSpace)
	return fun(info)
}

func parseEnv() (*Info, error) {
	env := os.Getenv("CF_ENVIRONMENT")
	file, err := os.Open(env)
	if err != nil {
		err = errors.New(fmt.Sprintf("Can not open %s : %s", env, err.Error()))
		return nil, err
	}
	info := Info{}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &info)
	return &info, err
}
