package cloudcontroller

import (
	"fmt"

	. "github.com/pivotalservices/uaaldapimport/token"
)

var AssociateOrg OrgFunc = func(orgInfo OrgInfo) (spaceInfos []SpaceInfo, err error) {
	_, err = RequestWithToken(orgInfo.Info.Token, fmt.Sprintf("%s/v2/organizations/%s/users/%s", orgInfo.Info.Ccurl, orgInfo.Org.Guid, orgInfo.Info.UserId), "PUT", "application/json", nil)
	if err != nil {
		return
	}
	for _, role := range orgInfo.Org.Roles {
		_, err = RequestWithToken(orgInfo.Info.Token, fmt.Sprintf("%s/v2/organizations/%s/%s/%s", orgInfo.Info.Ccurl, orgInfo.Org.Guid, role, orgInfo.Info.UserId), "PUT", "application/json", nil)
		if err != nil {
			return
		}
	}
	spaceInfos = make([]SpaceInfo, 0)
	for _, space := range orgInfo.Org.Spaces {
		spaceInfo := SpaceInfo{
			Org:   orgInfo,
			Space: space,
		}
		spaceInfos = append(spaceInfos, spaceInfo)
	}
	return
}

var AssociateSpace SpaceFunc = func(spaceInfo SpaceInfo) (err error) {
	for _, role := range spaceInfo.Space.Roles {
		_, err = RequestWithToken(spaceInfo.Org.Info.Token, fmt.Sprintf("%s/v2/spaces/%s/%s/%s", spaceInfo.Org.Info.Ccurl, spaceInfo.Space.Guid, role, spaceInfo.Org.Info.UserId), "PUT", "applicatio    n/json", nil)
		if err != nil {
			return
		}
	}
	return
}
