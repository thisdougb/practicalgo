package task

import (
	"log"

	"github.com/thisdougb/practicalgo/config"
)

type (

	// the main type entry point
	ListUsersTask struct {
		ds    ListUsersTaskInterface
		cfg   *config.Config // access to dynamic config settings module
		debug bool           // useful for isolating unit test output during development
	}

	// An interface defined specifically for this type means we can
	// be more granular (less messy) with the mock datastore methods.
	ListUsersTaskInterface interface {
		GetUsers() ([]string, error)
	}
)

func NewListUsersTask(datastore ListUsersTaskInterface) *ListUsersTask {
	return &ListUsersTask{
		ds: datastore,
	}
}

// ----- Public Methods -----

// CLI entry point
func (t *ListUsersTask) CLI_GetUsers(debug bool) {

	t.debug = debug
	users, _ := t.getUsers()

	log.Println("CLI_GetUsers(): Users:", users)

}

// Typically this is the webhandler entry point
func (t *ListUsersTask) GetUsers() ([]string, error) {
	return t.getUsers()
}

// ----- Private Methods -----
func (t *ListUsersTask) getUsers() ([]string, error) {
	storedUsers, err := t.ds.GetUsers()
	if err != nil {
		if t.debug {
			log.Println("Error in ListUsersTask.getUsers(): ", err.Error())
		}
		return storedUsers, err
	}

	return storedUsers, nil
}
