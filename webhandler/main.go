package webhandler

import (
	"github.com/thisdougb/practicalgo/config"
	"github.com/thisdougb/practicalgo/task"
)

type (
	Handler struct {
		ds    Datastore
		cfg   *config.Config // access to dynamic config settings module
		debug bool
	}

	Datastore interface {
		task.ListUsersTaskInterface
	}
)

func New(datastore Datastore) *Handler {
	return &Handler{
		ds: datastore,
	}
}
