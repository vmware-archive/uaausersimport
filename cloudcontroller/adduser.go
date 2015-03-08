package cloudcontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/pivotalservices/uaaldapimport/token"
)

type UserRequest struct {
	Id string `json:"guid"`
}

var Adduser TokenFunc = func(info *Info) (infoRet *Info, err error) {
	userRequest := UserRequest{
		Id: info.UserId,
	}
	data, err := json.Marshal(userRequest)
	if err != nil {
		return
	}
	body := bytes.NewBuffer(data)
	_, err = RequestWithToken(info.Token, fmt.Sprintf("%s/v2/users", info.Ccurl), "POST", "application/json", body)
	return info, err
}
