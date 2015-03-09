package cloudcontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pivotalservices/uaaldapimport/functions"
	. "github.com/pivotalservices/uaaldapimport/token"
)

type UserRequest struct {
	Id string `json:"guid"`
}

var Adduser functions.CCAddUserFunc = func(info functions.UserIdInfo) (err error) {
	fmt.Println(fmt.Sprintf("add user id: %s with Uaa id: %s .........", info.User.Uid, info.UserId))
	userRequest := UserRequest{
		Id: info.UserId,
	}
	data, err := json.Marshal(userRequest)
	if err != nil {
		return
	}
	body := bytes.NewBuffer(data)
	_, err = RequestWithToken(info.Token, fmt.Sprintf("%s/v2/users", info.Ccurl), "POST", "application/json", body)
	return err
}
