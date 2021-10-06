package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform/domain"
)

func (m *nosql) CatchErrNoResults(handler func(p diary.IPage)) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			if assertedErr, ok := r.(error); ok {
				err = assertedErr
			}
			if err.Error() != domain.ErrNoResults.Error() {
				panic(err)
			}
		}
	}()

	if handler != nil {
		handler(m.p)
	}
}