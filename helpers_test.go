package uniform

import (
	"testing"
)

type message struct {
	Message string
}

func TestCommunication(t *testing.T)  {
	input := message{
		Message: "hello world!",
	}
	data, err := encode(input)
	if err != nil {
		panic(err)
	}

	var output message
	if err := decode(data, &output); err != nil {
		panic(err)
	}

	if output != input {
		t.Error("output model is not the same as the input model")
	}
}

func TestCommunicationMismatchError(t *testing.T)  {
	input := message{
		Message: "hello world!",
	}
	data, err := encode(input)
	if err != nil {
		panic(err)
	}

	var output string
	err = decode(data, &output)
	if err == nil || err.Error() != "gob: decoding into local type *string, received remote type message = struct { Message string; }" {
		t.Error("expecting mismatch error but did not receive one")
	}
}