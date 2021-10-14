package uniform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommunication(t *testing.T)  {
	type payload struct {
		Message string
	}

	expected := payload{
		Message: "hello world!",
	}
	data, err := encode(expected)
	if err != nil {
		panic(err)
	}

	var output payload
	if err := decode(data, &output); err != nil {
		panic(err)
	}

	assert.Equal(t, expected, output)
}

func TestCommunicationMismatchError(t *testing.T)  {
	type payload struct {
		Value int
	}

	input := payload{
		Value: 123,
	}
	data, err := encode(input)
	if err != nil {
		panic(err)
	}

	var output string
	err = decode(data, &output)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	assert.Contains(t, errMsg, "error decoding")
}