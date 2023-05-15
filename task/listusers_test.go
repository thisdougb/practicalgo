//go:build dev

package task

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// simple table driven test layout
func TestGetUsers(t *testing.T) {

	testCases := []struct {
		description string // identify the test
		storedUsers []string
		expectError error
		expectUsers []string
	}{
		{
			description: "no stored users",
			storedUsers: []string{},
			expectError: errors.New("no users found"),
			expectUsers: []string{},
		},
		{
			description: "one stored user",
			storedUsers: []string{"user1"},
			expectError: nil,
			expectUsers: []string{"user1"},
		},
	}

	for _, tc := range testCases {

		mockDatastore := MockDatastore{}
		task := NewListUsersTask(&mockDatastore)

		// setup the mock datastore values for this test case
		mockDatastore.storedUsers = tc.storedUsers

		retval, err := task.getUsers()
		assert.Equal(t, tc.expectUsers, retval, tc.description)
		assert.Equal(t, tc.expectError, err, tc.description)
	}
}
