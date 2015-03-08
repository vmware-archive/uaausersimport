package main

import (
	"fmt"
	"os"

	uaa "github.com/pivotalservices/uaaldapimport/adduser"
	cc "github.com/pivotalservices/uaaldapimport/cloudcontroller"
	config "github.com/pivotalservices/uaaldapimport/config"
	. "github.com/pivotalservices/uaaldapimport/functions"
	token "github.com/pivotalservices/uaaldapimport/token"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() (err error) {
	file, err := os.Open("config/fixtures/users.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg, err := config.Parse(file)
	if err != nil {
		return
	}
	info := &Info{
		Uaaurl:   "https://uaa.sys.homelab.io",
		Ccurl:    "https://api.sys.homelab.io",
		Clientid: "bulkimport",
		Secret:   "test",
	}
	fun := token.GetToken.MapUsers(cfg.Users).AddUaaUser(uaa.Adduser).AddCCUser(cc.Adduser).MapOrgs(cc.AssociateOrg).MapSpaces(cc.AssociateSpace)
	//fun := token.GetToken.MapUsers(cfg.Users)
	return fun(info)
}
