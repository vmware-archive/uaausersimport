package functions

import (
	"io"
	"net/http"

	"github.com/pivotalservices/uaaldapimport/config"
)

type RequestTokenFunc func(string, string, string, string, io.Reader) (*http.Response, error)

type Info struct {
	Ccurl     string
	Uaaurl    string
	Clientid  string
	Secret    string
	RequestFn RequestTokenFunc
}

type TokenFunc func(*Info) (string, error)

type UserInfo struct {
	*Info
	Token  string
	User   config.User
	Origin string
}

type UserIdInfo struct {
	UserInfo
	UserId string
}

type UaaAddUserFunc func(UserInfo) (string, error)
type CCAddUserFunc func(UserIdInfo) error

type OrgInfo struct {
	UserIdInfo
	Org  config.Org
	Guid string
}
type SpaceInfo struct {
	OrgInfo
	Space config.Space
}

type OrgFunc func(OrgInfo) (string, error)
type SpaceFunc func(SpaceInfo) error
type UserFuncs func(*Info) ([]UserInfo, error)
type UserIdFuncs func(*Info) ([]UserIdInfo, error)
type OrgFuncs func(*Info) ([]OrgInfo, error)
type SpaceFuncs func(*Info) error

func (tokenFunc TokenFunc) MapUsers(config config.Config) UserFuncs {
	return func(info *Info) ([]UserInfo, error) {
		token, err := tokenFunc(info)
		if err != nil {
			return nil, err
		}
		users := make([]UserInfo, 0)
		for _, user := range config.Users {
			userInfo := UserInfo{
				Info:   info,
				Token:  token,
				User:   user,
				Origin: config.Origin,
			}
			users = append(users, userInfo)
		}
		return users, nil
	}
}

func (userFuncs UserFuncs) AddUaaUser(addUserFunc UaaAddUserFunc) UserIdFuncs {
	return func(info *Info) ([]UserIdInfo, error) {
		userInfos, err := userFuncs(info)
		if err != nil {
			return nil, err
		}
		users := make([]UserIdInfo, 0)
		for _, user := range userInfos {
			id, err := addUserFunc(user)
			if err != nil {
				return nil, err
			}
			userIdInfo := UserIdInfo{
				UserInfo: user,
				UserId:   id,
			}
			users = append(users, userIdInfo)
		}
		return users, nil
	}
}

func (userIdFuncs UserIdFuncs) AddCCUser(addUserFunc CCAddUserFunc) UserIdFuncs {
	return func(info *Info) ([]UserIdInfo, error) {
		userIdInfos, err := userIdFuncs(info)
		if err != nil {
			return nil, err
		}
		for _, user := range userIdInfos {
			err := addUserFunc(user)
			if err != nil {
				return nil, err
			}
		}
		return userIdInfos, nil
	}
}

func (userIdFuncs UserIdFuncs) MapOrgs(orgFunc OrgFunc) OrgFuncs {
	return func(info *Info) ([]OrgInfo, error) {
		userIdInfos, err := userIdFuncs(info)
		if err != nil {
			return nil, err
		}
		orgInfos := make([]OrgInfo, 0)
		for _, userIdInfo := range userIdInfos {
			for _, org := range userIdInfo.User.Orgs {
				orgInfo := OrgInfo{
					UserIdInfo: userIdInfo,
					Org:        org,
				}
				guid, err := orgFunc(orgInfo)
				if err != nil {
					return nil, err
				}
				orgInfo.Guid = guid
				orgInfos = append(orgInfos, orgInfo)
			}
		}
		return orgInfos, nil
	}
}

func (orgFuncs OrgFuncs) MapSpaces(spaceFunc SpaceFunc) SpaceFuncs {
	return func(info *Info) error {
		orgInfos, err := orgFuncs(info)
		if err != nil {
			return err
		}
		for _, orgInfo := range orgInfos {
			spaces := orgInfo.Org.Spaces
			for _, space := range spaces {
				spaceInfo := SpaceInfo{
					Space:   space,
					OrgInfo: orgInfo,
				}
				err := spaceFunc(spaceInfo)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}
