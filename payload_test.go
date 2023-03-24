package uniform

import (
    "fmt"
    "testing"
)

func TestRead(t *testing.T) {
    p := payload{
        Request: Request{},
    }
    data := []byte{ 0x00 }
    p.Read(&data)
    fmt.Println(data)
}