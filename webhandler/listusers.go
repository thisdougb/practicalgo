package webhandler

import (
	"fmt"
	"net/http"

	"github.com/thisdougb/practicalgo/task"
)

func (h *Handler) ListUsersEndpoint(w http.ResponseWriter, r *http.Request) {

	listUsersTask := task.NewListUsersTask(h.ds)
	storedUsers, _ := listUsersTask.GetUsers()

	if len(storedUsers) == 0 {
		fmt.Fprintf(w, "No users found.\n")
		return
	}

	fmt.Fprintf(w, "Stored users are: %s\n", storedUsers)
}
