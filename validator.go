package uniform

// A definition of the public functions for a request interface
type IValidator interface {
	Error(field, error string)
	Check()
	SilentCheck() string
}