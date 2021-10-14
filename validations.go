package uniform

import (
	"time"
)

var ValidateRequired = func(value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateMinimumInt = func(minimum int64, value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateMaximumInt = func(maximum int64, value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateMinimumFloat = func(minimum float64, value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateMaximumFloat = func(maximum float64, value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateRangeInt = func(minimum, maximum int64, value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateRangeFloat = func(minimum, maximum float64, value interface{}) (bool, []string) {
	panic("not yet implemented")
}

var ValidateMobile = func(country string, value interface{}) (bool, []string, interface{}) {
	panic("not yet implemented")
}

var ValidateEmail = func(value interface{}) (bool, []string, interface{}) {
	panic("not yet implemented")
}

var ValidatePassportNumber = func(country string, value interface{}) (bool, []string, interface{}) {
	panic("not yet implemented")
}

var ValidateIdentityNumber = func(country string, value interface{}) (bool, []string, interface{}) {
	panic("not yet implemented")
}

var ValidateDate = func(american bool, value interface{}) (bool, []string, time.Time) {
	panic("not yet implemented")
}

var ValidateDateTime = func(american bool, value interface{}) (bool, []string, time.Time) {
	panic("not yet implemented")
}
