//go:build dev

/*
This file is included in the build when the "dev" build tag is used.
This file is used to set up mock values for unit testing.

Its filename is numbered to ensure it is included after the main.init() in the build.
This means that other packages calling this code with the dev tag, will use the mocked version.
*/

package task

import "log"

// Set any state vars for unit test here
var (
	mockPauseValue        int
	mockCurrentUsageValue int
)

func init() {
	log.Println("task.init(): mock init()")

	// replace the pause function to facilitate testing
	// sets var values that we use in unit tests, instead of causing a pause
	pause = func(seconds int) { mockPauseValue = seconds; mockCurrentUsageValue = 0 }
}
