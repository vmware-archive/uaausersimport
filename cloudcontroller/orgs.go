package cloudcontroller

import (
	"fmt"

	"github.com/pivotalservices/uaausersimport/functions"
)

var AssociateOrg functions.OrgFunc = func(orgInfo functions.OrgInfo) (guid string, err error) {
	response, err := orgInfo.RequestFn(orgInfo.Token, fmt.Sprintf("%s/v2/organizations?q=name:%s", orgInfo.Ccurl, orgInfo.Org.Name), "GET", "application/json", nil)
	if err != nil {
		return
	}
	guid, err = parseResponse(response)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Associate user id :%s to org: %s.........", orgInfo.User.Uid, guid))
	_, err = orgInfo.RequestFn(orgInfo.Token, fmt.Sprintf("%s/v2/organizations/%s/users/%s", orgInfo.Ccurl, guid, orgInfo.UserId), "PUT", "application/json", nil)
	if err != nil {
		return
	}
	for _, role := range orgInfo.Org.Roles {
		fmt.Println(fmt.Sprintf("Associate user id :%s to org: %s with %s role.........", orgInfo.User.Uid, guid, role))
		_, err = orgInfo.RequestFn(orgInfo.Token, fmt.Sprintf("%s/v2/organizations/%s/%s/%s", orgInfo.Ccurl, guid, role, orgInfo.UserId), "PUT", "application/json", nil)
		if err != nil {
			return
		}
	}
	return
}
