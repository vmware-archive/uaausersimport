package adduser

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/pivotalservices/uaaldapimport/config"
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
	transport := NewRoundTripper()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/Users", uaaurl), body)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf(" Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	dump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(dump))
	response, err := transport.RoundTrip(req)
	if err != nil {
		return
	}
	dump, _ = httputil.DumpResponse(response, true)
	fmt.Println(string(dump))
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
