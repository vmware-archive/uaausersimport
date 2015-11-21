package cloudcontroller

import (
	"fmt"

	"github.com/pivotalservices/uaausersimport/functions"
)

var AssociateSpace functions.SpaceFunc = func(spaceInfo functions.SpaceInfo) (err error) {
	response, err := spaceInfo.RequestFn(spaceInfo.Token, fmt.Sprintf("%s/v2/spaces?q=name:%s&q=organization_guid:%s", spaceInfo.Ccurl, spaceInfo.Space.Name, spaceInfo.OrgInfo.Guid), "GET", "application/json", nil)
	if err != nil {
		return
	}
	guid, err := parseResponse(response)
	if err != nil {
		return
	}
	for _, role := range spaceInfo.Space.Roles {
		fmt.Println(fmt.Sprintf("Associate user id :%s to space: %s with %s role.........", spaceInfo.User.Uid, guid, role))
		_, err = spaceInfo.RequestFn(spaceInfo.Token, fmt.Sprintf("%s/v2/spaces/%s/%s/%s", spaceInfo.Ccurl, guid, role, spaceInfo.UserId), "PUT", "applicatio    n/json", nil)
		if err != nil {
			return
		}
	}
	return
}
