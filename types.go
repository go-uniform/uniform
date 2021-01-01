package uniform

import "github.com/go-diary/diary"

// A package shorthand for a map[string]interface
type M map[string]interface{}

// A package shorthand for a map[string]string
type P map[string]string

// A package shorthand for map[string][]string
type Q map[string][]string

// A package shorthand for func(r IRequest, p diary.IPage)
type S func(r IRequest, p diary.IPage)
