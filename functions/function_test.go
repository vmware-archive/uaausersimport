package functions_test

import (
	"os"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotalservices/uaaldapimport/config"
	. "github.com/pivotalservices/uaaldapimport/functions"
)

var info *Info = &Info{
	Ccurl:    "https://ccurl.sysdomain.com",
	Uaaurl:   "https://uaaurl.sysdomain.com",
	Clientid: "bulkimport",
	Secret:   "test",
}

var _ = Describe("Function", func() {
	file, _ := os.Open("fixtures/users.yml")
	cfg, _ := config.Parse(file)
	var _ = Describe("Map Users", func() {
		It("Should map users with tokenFuncs", func() {
			var tokenFunc TokenFunc = func(*Info) (string, error) {
				return "my_token", nil
			}
			resultFunc := tokenFunc.MapUsers(cfg.Users)
			userInfos, err := resultFunc(info)
			Ω(err).Should(BeNil())
			Ω(len(userInfos)).Should(Equal(2))
			Ω(userInfos[0].Token).Should(Equal("my_token"))
			Ω(reflect.DeepEqual(userInfos[0].User, cfg.Users[0])).Should(Equal(true))
			Ω(reflect.DeepEqual(userInfos[1].User, cfg.Users[1])).Should(Equal(true))
			Ω(userInfos[1].Token).Should(Equal("my_token"))
		})
	})
})
