package uniform

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testWrapper = func(t *testing.T, expectedPanicMessage string, testCase func()) {
	defer func() {
		if r := recover(); r != nil {
			panicMessage := fmt.Sprint(r)
			assert.Equal(t, expectedPanicMessage, panicMessage)
		} else {
			assert.Equal(t, expectedPanicMessage, "")
		}
	}()

	testCase()
}

func TestValidateRequired(t *testing.T) {
	type TestCase struct {
		Value interface{}

		ExpectedPanicMessage string
		ExpectedValid bool
		ExpectedErrors []string
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		/* Positive tests */
		{
			Value: "test",
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Value: false,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Value: true,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Value: 0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Value: 1,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Value: 0.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Value: 1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},

		/* Negative tests */
		{
			Value: nil,
			ExpectedValid: false,
			ExpectedErrors: []string{ "May not be empty" },
		},
		{
			Value: "",
			ExpectedValid: false,
			ExpectedErrors: []string{ "May not be empty" },
		},
		{
			Value: "   ", // empty spaces should not be seen as populated
			ExpectedValid: false,
			ExpectedErrors: []string{ "May not be empty" },
		},
	}

	for index, test := range tests {
		var valid bool
		var errors []string
		testWrapper(t, test.ExpectedPanicMessage, func() {
			valid, errors = ValidateRequired(test.Value)
		})
		if !assert.Equal(t, test.ExpectedErrors, errors, "Test #%d", index + 1) {
			break
		}
		if !assert.Equal(t, test.ExpectedValid, valid, "Test #%d", index + 1) {
			break
		}
	}
}

func TestValidateMinimumInt(t *testing.T) {
	type TestCase struct {
		Minimum int64
		Value interface{}

		ExpectedPanicMessage string
		ExpectedValid bool
		ExpectedErrors []string
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		/* Positive tests */
		{
			Minimum: 0,
			Value: 0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: 0,
			Value: 0.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: 1,
			Value: 1,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: 1,
			Value: 1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: 1,
			Value: 1.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: -1,
			Value: -1,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: -1,
			Value: -1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: -1,
			Value: -0.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},

		/* Negative tests */
		{
			Minimum: 1,
			Value: 0,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1 or greater" },
		},
		{
			Minimum: 1,
			Value: 0.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1 or greater" },
		},
		{
			Minimum: -1,
			Value: -1.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be -1 or greater" },
		},
		{
			Minimum: 1,
			Value: "test",
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or greater" },
		},
		{
			Minimum: 1,
			Value: false,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or greater" },
		},
		{
			Minimum: 1,
			Value: true,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or greater" },
		},
	}

	for index, test := range tests {
		var valid bool
		var errors []string
		testWrapper(t, "", func() {
			valid, errors = ValidateMinimumInt(test.Minimum, test.Value)
		})
		if !assert.Equal(t, test.ExpectedErrors, errors, "Test #%d", index + 1) {
			break
		}
		if !assert.Equal(t, test.ExpectedValid, valid, "Test #%d", index + 1) {
			break
		}
	}
}

func TestValidateMaximumInt(t *testing.T) {
	type TestCase struct {
		Maximum int64
		Value interface{}

		ExpectedPanicMessage string
		ExpectedValid bool
		ExpectedErrors []string
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		/* Positive tests */
		{
			Maximum: 0,
			Value: 0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: 0,
			Value: 0.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: 1,
			Value: 1,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: 1,
			Value: 1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: 1,
			Value: 0.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: -1,
			Value: -1,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: -1,
			Value: -1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: -1,
			Value: -1.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},

		/* Negative tests */
		{
			Maximum: 0,
			Value: 1,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1 or less" },
		},
		{
			Maximum: 1,
			Value: 1.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1 or less" },
		},
		{
			Maximum: -1,
			Value: -1.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be -1 or less" },
		},
		{
			Maximum: 1,
			Value: "test",
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or less" },
		},
		{
			Maximum: 1,
			Value: false,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or less" },
		},
		{
			Maximum: 1,
			Value: true,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or less" },
		},
	}

	for index, test := range tests {
		var valid bool
		var errors []string
		testWrapper(t, "", func() {
			valid, errors = ValidateMaximumInt(test.Maximum, test.Value)
		})
		if !assert.Equal(t, test.ExpectedErrors, errors, "Test #%d", index + 1) {
			break
		}
		if !assert.Equal(t, test.ExpectedValid, valid, "Test #%d", index + 1) {
			break
		}
	}
}

func TestValidateMinimumFloat(t *testing.T) {
	type TestCase struct {
		Minimum float64
		Value interface{}

		ExpectedPanicMessage string
		ExpectedValid bool
		ExpectedErrors []string
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		/* Positive tests */
		{
			Minimum: 0.0,
			Value: 0.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: 1.0,
			Value: 1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: 1.0,
			Value: 1.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: -1.0,
			Value: -1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Minimum: -1.0,
			Value: -0.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},

		/* Negative tests */
		{
			Minimum: 1.0,
			Value: 0.0,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1.0 or greater" },
		},
		{
			Minimum: 1.0,
			Value: 0.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1.0 or greater" },
		},
		{
			Minimum: -1.0,
			Value: -1.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be -1.0 or greater" },
		},
		{
			Minimum: 1.0,
			Value: "test",
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1.0 or greater" },
		},
		{
			Minimum: 1.0,
			Value: false,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1.0 or greater" },
		},
		{
			Minimum: 1.0,
			Value: true,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1.0 or greater" },
		},
	}

	for index, test := range tests {
		var valid bool
		var errors []string
		testWrapper(t, "", func() {
			valid, errors = ValidateMinimumFloat(test.Minimum, test.Value)
		})
		if !assert.Equal(t, test.ExpectedErrors, errors, "Test #%d", index + 1) {
			break
		}
		if !assert.Equal(t, test.ExpectedValid, valid, "Test #%d", index + 1) {
			break
		}
	}
}

func TestValidateMaximumFloat(t *testing.T) {
	type TestCase struct {
		Maximum float64
		Value interface{}

		ExpectedPanicMessage string
		ExpectedValid bool
		ExpectedErrors []string
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		/* Positive tests */
		{
			Maximum: 0.0,
			Value: 0.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: 1.0,
			Value: 1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: 1.0,
			Value: 0.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: -1.0,
			Value: -1.0,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},
		{
			Maximum: -1.0,
			Value: -1.5,
			ExpectedValid: true,
			ExpectedErrors: nil,
		},

		/* Negative tests */
		{
			Maximum: 0.0,
			Value: 1.0,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1 or less" },
		},
		{
			Maximum: 1.0,
			Value: 1.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be 1 or less" },
		},
		{
			Maximum: -1.0,
			Value: -1.5,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be -1 or less" },
		},
		{
			Maximum: 1.0,
			Value: "test",
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or less" },
		},
		{
			Maximum: 1.0,
			Value: false,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or less" },
		},
		{
			Maximum: 1.0,
			Value: false,
			ExpectedValid: false,
			ExpectedErrors: []string{ "Must be a numeric value", "Must be 1 or less" },
		},
	}

	for index, test := range tests {
		var valid bool
		var errors []string
		testWrapper(t, "", func() {
			valid, errors = ValidateMaximumFloat(test.Maximum, test.Value)
		})
		if !assert.Equal(t, test.ExpectedErrors, errors, "Test #%d", index + 1) {
			break
		}
		if !assert.Equal(t, test.ExpectedValid, valid, "Test #%d", index + 1) {
			break
		}
	}
}

func TestValidateRangeInt(t *testing.T) {

}

func TestValidateRangeFloat(t *testing.T) {

}

func TestValidateMobile(t *testing.T) {

}

func TestValidateEmail(t *testing.T) {

}

func TestValidatePassportNumber(t *testing.T) {

}

func TestValidateIdentityNumber(t *testing.T) {

}

func TestValidateDate(t *testing.T) {

}

func TestValidateDateTime(t *testing.T) {

}