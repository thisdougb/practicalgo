//go:build dev

package task

import "errors"

// Setting state for individual unit tests is cleaner this way, rather
// than overriding methods per test case.
type MockDatastore struct {
	storedUsers []string
}

func (m *MockDatastore) init() {}

func (m *MockDatastore) GetUsers() ([]string, error) {
	if len(m.storedUsers) > 0 {
		return m.storedUsers, nil
	}
	return []string{}, errors.New("no users found")
}
