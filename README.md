# uaaldapimport


## Problem
When cf integrate with ldap. Currently cloudfoundry does not have any way to assgin user roles (E.g. org and space). So user has to login to the CF env first, then operator assign the roles to them. When they login again, users got see the spaces.

This may not be applicable for the invitation model (login first, say nothing and login back) for operators.

## Resolution

Prepopulate the users with uaa api and cloudcontroller api. So they can have all the roles before user logins.

Prerequisites:

Use uaac to create a client id, who has cloudcontroller.admin and scim.write

Steps (What this progam is doing?):

1. Get token from uaa
2. Add user to the uaa
   * sample users yaml file: [sample file](config/fixtures/users.yml)
   * Sample user config

   ```
   users:
   - uid: jcalabrese@pivotal.io
     externalid: uid=jcalabrese,ou=People,dc=homelab,dc=io
     emails:
     - jcalabrese@pivotal.io
     orgs:
       - guid: 1ff0d5d8-cd7f-4cf6-bf8d-d6632cac2921
         roles:
         - managers
         - auditors
         spaces:
           - guid: 96243054-2eaa-4fb9-99cf-b9c37920ce6b
             roles:
             - managers
             - developers
             - auditors
           - guid: 7207320b-2384-42bd-b893-4afef7f7b209
             roles:
             - managers
             - auditors
       - guid: a8e28250-989a-4214-b839-ee71e1d1b72a
         roles:
         - auditors
         spaces:
           - guid: a6a7f87b-c236-4dc6-8cf1-df7d602b228f
             roles:
             - auditors
           - guid: 0898f4d0-69e2-4ec1-9947-4dc19c980042
             roles:
             - auditors
   ```

3. Add user to the cloudcontroler
4. Associate user roles with the orgs
5. Associate user roles with the spaces

   * Functional Programming (In [main.go](main.go))

   ```
      token.GetToken.MapUsers(cfg.Users).AddUaaUser(uaa.Adduser).AddCCUser(cc.Adduser).MapOrgs(cc.AssociateOrg).MapSpaces(cc.AssociateSpace)
   ```

## How to run

* Install go

* Get the binary

```
go get -u go get -u github.com/pivotalservices/uaaldapimport

```
* Target the cf environment

```
export CF_ENVIRONMENT=environment.yml (change to your environment.yml)
```

* Target the users file

```
export LDAP_USERS=config/fixtures/users.yml (change to your user files)
```
* Enable http traffic dump, optional:

  ```
  export DEBUG_HTTP=true
  ```

  
