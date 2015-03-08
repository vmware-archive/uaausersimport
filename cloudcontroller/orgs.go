package cloudcontroller

import (
	"fmt"

	"github.com/pivotalservices/uaaldapimport/functions"
	. "github.com/pivotalservices/uaaldapimport/token"
)

var AssociateOrg functions.OrgFunc = func(orgInfo functions.OrgInfo) (err error) {
	_, err = RequestWithToken(orgInfo.Token, fmt.Sprintf("%s/v2/organizations/%s/users/%s", orgInfo.Ccurl, orgInfo.Org.Guid, orgInfo.UserId), "PUT", "application/json", nil)
	if err != nil {
		return
	}
	for _, role := range orgInfo.Org.Roles {
		_, err = RequestWithToken(orgInfo.Token, fmt.Sprintf("%s/v2/organizations/%s/%s/%s", orgInfo.Ccurl, orgInfo.Org.Guid, role, orgInfo.UserId), "PUT", "application/json", nil)
		if err != nil {
			return
		}
	}
	return
}

var AssociateSpace functions.SpaceFunc = func(spaceInfo functions.SpaceInfo) (err error) {
	for _, role := range spaceInfo.Space.Roles {
		_, err = RequestWithToken(spaceInfo.Token, fmt.Sprintf("%s/v2/spaces/%s/%s/%s", spaceInfo.Ccurl, spaceInfo.Space.Guid, role, spaceInfo.UserId), "PUT", "applicatio    n/json", nil)
		if err != nil {
			return
		}
	}
	return
}
