package uniform

import (
	"github.com/go-diary/diary"
	"reflect"
	"time"
)

// A package shorthand for a map[string]interface
type M map[string]interface{}

// A package shorthand for a map[string]string
type P map[string]string

// A package shorthand for map[string][]string
type Q map[string][]string

// A package shorthand for func(r IRequest, p diary.IPage)
type S func(r IRequest, p diary.IPage)

type DateTime struct {
	UtcTime time.Time

	Timezone           string
	TimezoneAdjustment time.Duration
	TimezoneTime       time.Time

	DaylightSavings           bool
	DaylightSavingsStart      time.Time
	DaylightSavingsEnd        time.Time
	DaylightSavingsAdjustment time.Duration
	LocalTime                 time.Time
}

type Money struct {
	CurrencyCode   string
	CurrencySymbol string
	DateTime       DateTime

	WholeNumber int64
	Precision   byte
	Value       float64

	Display string
}

type Real struct {
	WholeNumber int64
	Precision   byte
	Value       float64
}

type Protected struct {
	Kind      reflect.Kind
	Hash      string
	Encrypted []byte
}

func (p *Protected) CompareHash(value string) bool {
	panic("not yet implemented")
}

func (p *Protected) Decrypt() interface{} {
	panic("not yet implemented")
}

func NewProtectedValue(value interface{}) Protected {
	panic("not yet implemented")
}
