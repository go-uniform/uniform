package uniform

type Request struct {
	Model      interface{}
	Parameters P
	Context    M
	Error      string
}
