package datastore

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/thisdougb/practicalgo/config"
)

type Datastore struct {
	host     string
	port     string
	username string
	password string
	database string
	tls      bool
	ctx      context.Context
	client   *redis.Client
	cfg      *config.Config // access to dynamic config settings module
	debug    bool           // useful for isolating unit test output during development
}

func NewRedisDatastore(host string, port string, username string, password string, tls bool) *Datastore {
	return &Datastore{
		host,
		port,
		username,
		password,
		"0", // default database
		tls,
		context.Background(),
		nil,
		nil, // config.cfg
		false}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Defining the keys used in the datastore this way is clearer, and allows key reuse.
const (
	UsersKey = "%s:users"
)
