package main

import (
	"log"

	"github.com/jtblin/go-ldap-client"
)

func main() {
	client := &ldap.LDAPClient{
		Base:         "ou=Artek,dc=Corp,dc=Artek,dc=Org",
		Host:         "172.16.0.2",
		Port:         389,
		UseSSL:       false,
		BindDN:       "cn=admin,ou=god,dc=corp,dc=artek, dc=org",
		BindPassword: "Ckj;ysqGfcc4321",
		UserFilter:   "(sAMAccountName=%s)",
		GroupFilter:  "(memberUid=%s)",
		Attributes:   []string{"givenName", "sn", "mail", "uid"},
	}
	// It is the responsibility of the caller to close the connection
	defer client.Close()
	username := "DShkabatur"
	password := "Denis67058343"
	ok, user, err := client.Authenticate(username, password)
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", username, err)
	}
	if !ok {
		log.Fatalf("Authenticating failed for user %s", username)
	}
	log.Printf("User: %+v", user)

	groups, err := client.GetGroupsOfUser(username)
	if err != nil {
		log.Fatalf("Error getting groups for user %s: %+v", username, err)
	}
	log.Printf("Groups: %+v", groups)
}
