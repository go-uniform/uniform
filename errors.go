package uniform

import (
	"errors"
)

// ErrNoResults when operation did not return any results
var ErrNoResults = errors.New("no results found")