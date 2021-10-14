package uniform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexOf(t *testing.T) {
	type TestCase struct {
		Haystack      []string
		Needle        string
		CaseSensitive bool

		ExpectedPanicMessage string
		ExpectedOutput       int
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		//Positive tests
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "needle1",
			CaseSensitive: false,

			ExpectedOutput: 0,
		},
		{
			Haystack:      []string{"needle1", "needle1", "needle2"},
			Needle:        "needle1",
			CaseSensitive: false,

			ExpectedOutput: 0,
		},
		{
			Haystack:      []string{"needle1", "needle1", "needle2", "Needle2"},
			Needle:        "Needle2",
			CaseSensitive: true,

			ExpectedOutput: 3,
		},

		//Negative tests
		{
			Haystack:      nil,
			Needle:        "Needle1",
			CaseSensitive: false,

			ExpectedOutput: -1,
		},
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "",
			CaseSensitive: false,

			ExpectedOutput: -1,
		},
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "needle3",
			CaseSensitive: false,

			ExpectedOutput: -1,
		},
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "Needle2",
			CaseSensitive: true,

			ExpectedOutput: -1,
		},
	}

	for index, test := range tests {
		var output int
		testWrapper(t, test.ExpectedPanicMessage, func() {
			output = IndexOf(test.Haystack, test.Needle, test.CaseSensitive)
		})
		if !assert.Equal(t, test.ExpectedOutput, output, "Test #%d", index+1) {
			break
		}
	}
}

func TestContains(t *testing.T) {
	type TestCase struct {
		Haystack      []string
		Needle        string
		CaseSensitive bool

		ExpectedPanicMessage string
		ExpectedOutput       bool
	}

	tests := []TestCase{
		/* Defensive tests */
		// this function will not have any defensive tests

		//Positive tests
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "needle1",
			CaseSensitive: false,

			ExpectedOutput: true,
		},
		{
			Haystack:      []string{"needle1", "needle1", "needle2"},
			Needle:        "needle1",
			CaseSensitive: false,

			ExpectedOutput: true,
		},
		{
			Haystack:      []string{"needle1", "needle1", "needle2", "Needle2"},
			Needle:        "Needle2",
			CaseSensitive: true,

			ExpectedOutput: true,
		},

		//Negative tests
		{
			Haystack:      nil,
			Needle:        "Needle1",
			CaseSensitive: false,

			ExpectedOutput: false,
		},
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "",
			CaseSensitive: false,

			ExpectedOutput: false,
		},
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "needle3",
			CaseSensitive: false,

			ExpectedOutput: false,
		},
		{
			Haystack:      []string{"needle1", "needle2"},
			Needle:        "Needle2",
			CaseSensitive: true,

			ExpectedOutput: false,
		},
	}

	for index, test := range tests {
		var output bool
		testWrapper(t, test.ExpectedPanicMessage, func() {
			output = Contains(test.Haystack, test.Needle, test.CaseSensitive)
		})
		if !assert.Equal(t, test.ExpectedOutput, output, "Test #%d", index+1) {
			break
		}
	}
}

func TestFilter(t *testing.T) {
	type TestCase struct {
		Items         []string
		Filters       []string
		CaseSensitive bool

		ExpectedPanicMessage string
		ExpectedOutput       []string
	}

	tests := []TestCase{
		/* Defensive tests */
		{
			Items:   nil,
			Filters: nil,

			ExpectedOutput: nil,
		},
		{
			Items:   nil,
			Filters: []string{},

			ExpectedOutput: nil,
		},
		{
			Items:   []string{},
			Filters: nil,

			ExpectedOutput: []string{},
		},
		{
			Items:   []string{},
			Filters: []string{},

			ExpectedOutput: []string{},
		},

		//Positive tests
		{
			Items:   []string{"needle1", "needle2"},
			Filters: []string{"needle1"},

			ExpectedOutput: []string{"needle2"},
		},
		{
			Items:   []string{"Needle1", "needle2"},
			Filters: []string{"Needle1"},
			CaseSensitive: true,

			ExpectedOutput: []string{"needle2"},
		},

		//Negative tests
		{
			Items:   []string{"needle1", "needle2"},
			Filters: []string{"needle3"},

			ExpectedOutput: []string{"needle1", "needle2"},
		},
		{
			Items:   []string{"needle1", "needle2"},
			Filters: []string{"Needle1"},
			CaseSensitive: true,

			ExpectedOutput: []string{"needle1", "needle2"},
		},
	}

	for index, test := range tests {
		var output []string
		testWrapper(t, test.ExpectedPanicMessage, func() {
			output = Filter(test.Items, test.Filters, test.CaseSensitive)
		})
		if !assert.Equal(t, test.ExpectedOutput, output, "Test #%d", index+1) {
			break
		}
	}
}
