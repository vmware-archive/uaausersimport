# uaaldapimport


## Problem
When cf integrates with ldap. Currently cloudfoundry does not have any way to assgin user roles (E.g. org and space). So user has to login to the CF env first, then operator assign the roles to them. When they login again, users can see the spaces.

This may not be applicable for the invitation model for operators, since they have to ask users to login first

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
   - uid: jcalabrese@pivotal.io
     externalid: uid=jcalabrese,ou=People,dc=homelab,dc=io
     emails:
     - jcalabrese@pivotal.io
     orgs:
       - name: org1
         roles:
         - managers
         - auditors
         spaces:
           - name: space1
             roles:
             - managers
             - developers
             - auditors
           - name: space2
             roles:
             - managers
             - auditors
       - name: org2
         roles:
         - auditors
         spaces:
           - name: space1
             roles:
             - auditors
           - name: space2
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

## Future work

* Create an interface (web/command line) help client generate formatted file
* Cross compile the code
