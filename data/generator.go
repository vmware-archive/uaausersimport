package main

import (
	"fmt"
	"github.com/pivotalservices/uaaldapimport/config"
	"gopkg.in/yaml.v2"
	"os"
)

func main() {
	cfg := config.Config{}
	cfg.Users = make([]config.User, 0)
	var yes string
	for {
		user := config.User{}
		var uid string
		fmt.Println("uid:")
		fmt.Scanf("%s", &uid)
		user.Uid = uid
		fmt.Println("dn:")
		var externalid string
		fmt.Scanf("%s", &externalid)
		user.Externalid = externalid
		fmt.Println("email:")
		var email string
		fmt.Scanf("%s", &email)
		emails := make([]string, 0)
		emails = append(emails, email)
		user.Emails = emails
		//enter orgs
		user.Orgs = make([]config.Org, 0)
		for {
			var orgName string
			fmt.Println("Org name:")
			fmt.Scanf("%s", &orgName)
			org := config.Org{}
			org.Name = orgName
			fmt.Println("Org Manager Role: default is no [y]")
			org.Roles = make([]string, 0)
			fmt.Scanf("%s", &yes)
			if isYes(yes) {
				org.Roles = append(org.Roles, "managers")
			}

			fmt.Println("Org Audit Role: default is no [y]")
			fmt.Scanf("%s", &yes)
			if isYes(yes) {
				org.Roles = append(org.Roles, "auditors")
			}
			org.Spaces = make([]config.Space, 0)
			fmt.Println("Creating roles for spaces? [y]")
			fmt.Scanf("%s", &yes)
			if isYes(yes) {
				for {
					space := config.Space{}
					fmt.Println("Space name:")
					fmt.Scanf("%s", &space.Name)
					space.Roles = make([]string, 0)
					fmt.Println("Space Manager Role: default is no [y]")
					fmt.Scanf("%s", &yes)
					if isYes(yes) {
						space.Roles = append(space.Roles, "managers")
					}
					fmt.Println("Space developer Role: default is no [y]")
					fmt.Scanf("%s", &yes)
					if isYes(yes) {
						space.Roles = append(space.Roles, "developers")
					}
					fmt.Println("Space auditor Role: default is no [y]")
					fmt.Scanf("%s", &yes)
					if isYes(yes) {
						space.Roles = append(space.Roles, "auditors")
					}
					org.Spaces = append(org.Spaces, space)
					fmt.Println("Done with space adding?")
					fmt.Scanf("%s", &yes)
					if isYes(yes) {
						break
					}
				}
			}
			user.Orgs = append(user.Orgs, org)
			fmt.Println("Done with org?")
			fmt.Scanf("%s", &yes)
			if isYes(yes) {
				break
			}
		}
		cfg.Users = append(cfg.Users, user)
		fmt.Println("Done with User?")
		fmt.Scanf("%s", &yes)
		if isYes(yes) {
			break
		}
	}
	err := writeToFile(&cfg)
	if err != nil {
		fmt.Println(err)
	}
}

func writeToFile(cfg *config.Config) (err error) {
	file, err := os.Create("users.yml")
	if err != nil {
		return
	}
	d, err := yaml.Marshal(cfg)
	if err != nil {
		return
	}
	_, err = file.Write(d)
	return
}

func isYes(yes string) bool {
	if yes == "y" {
		return true
	}
	return false
}
