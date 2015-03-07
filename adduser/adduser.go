package adduser

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pivotalservices/uaaldapimport/config"
	. "github.com/pivotalservices/uaaldapimport/token"
)

type Email struct {
	Value string `json:"value"`
}

type UserRequest struct {
	UserName   string  `json:"userName"`
	Emails     []Email `json:"emails"`
	Origin     string  `json:"origin"`
	Externalid string  `json:"externalid"`
}

type UserCreateResponse struct {
	Id string
}

var NewRoundTripper = func() http.RoundTripper {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func Adduser(token, uaaurl string, user *config.User) (guid string, err error) {
	emails := make([]Email, 0)
	for _, value := range user.Emails {
		email := Email{
			Value: value,
		}
		emails = append(emails, email)
	}
	userRequest := UserRequest{
		UserName:   user.Uid,
		Externalid: user.Externalid,
		Emails:     emails,
		Origin:     "ldap",
	}
	data, err := json.Marshal(userRequest)
	if err != nil {
		return
	}
	body := bytes.NewBuffer(data)
	response, err := RequestWithToken(token, fmt.Sprintf("%s/Users", uaaurl), "POST", "application/json", body)
	if err != nil {
		return
	}
	return parse(response)
}

func parse(response *http.Response) (guid string, err error) {
	body := response.Body
	defer body.Close()
	uResponse := UserCreateResponse{}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &uResponse)
	if err != nil {
		return
	}
	guid = uResponse.Id
	return
}
