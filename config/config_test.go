package config_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/uaaldapimport/config"
)

var _ = Describe("Config", func() {
	file, _ := os.Open("fixtures/users.txt")
	Describe("Parse config", func() {
		Context("with correct fixture", func() {
			It("should be return correct config", func() {
				cfg, err := config.Parse(file)
				fmt.Print(err)
				Ω(err).Should(BeNil())
				Ω(cfg.Sysdomain).Should(Equal("10.244.0.34.xip.io"))
				Ω(len(cfg.Users)).Should(Equal(2))
				Ω(cfg.Users[0].Externalid).Should(Equal("uid=sding,ou=People,dc=homelab,dc=io"))
				Ω(cfg.Users[1].Externalid).Should(Equal("uid=rparrish,ou=People,dc=homelab,dc=io"))
				Ω(cfg.Users[0].Emails[0]).Should(Equal("sding@pivotal.io"))
				Ω(cfg.Users[1].Emails[0]).Should(Equal("rparrish@pivotal.io"))
				Ω(len(cfg.Users[0].Orgs)).Should(Equal(2))
				Ω(cfg.Users[0].Orgs["org1"].Roles).Should(Equal("111"))
				Ω(len(cfg.Users[0].Orgs["org1"].Spaces)).Should(Equal(2))
				Ω(cfg.Users[0].Orgs["org1"].Spaces["space1"].Roles).Should(Equal("110"))
				Ω(cfg.Users[0].Orgs["org2"].Spaces["space2"].Roles).Should(Equal("001"))
				Ω(len(cfg.Users[0].Orgs["org2"].Spaces)).Should(Equal(2))
				Ω(cfg.Users[0].Orgs["org2"].Roles).Should(Equal("011"))
			})
		})
	})
})
