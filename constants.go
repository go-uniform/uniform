package uniform

import "errors"

// A package level reusable error for chain timeouts
var (
	ErrCantReply = errors.New("uniform: no reply channel available")
	ErrTimeout   = errors.New("uniform: timeout")
)
