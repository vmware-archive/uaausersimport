package cloudcontroller

import (
	"fmt"

	. "github.com/pivotalservices/uaaldapimport/token"
)

func AssociateOrg(token, orgId, ccurl, userid string) (err error) {
	_, err = RequestWithToken(token, fmt.Sprintf("%s/v2/organizations/%s/users/%s", ccurl, orgId, userid), "PUT", "application/json", nil)
	return
}
