package Storage

import (
	"fmt"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("192.168.64.40:31431")
	cluster.Keyspace = "shrink_url"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "admin",
		Password: "password"}

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
