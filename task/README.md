## Tasks

The *tasks* module enables actions to run via cli or Go routines in a loop.

#### Types

Defining a type for each task makes it easy to add new tasks, and to isolate the datastore methods.
Though care needs to be taken to keep the type simple, and to avoid duplication.

```
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
	}
)
```
