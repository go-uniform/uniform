package uniform

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// A definition of the public functions for a request interface
type IValidator interface {
	Error(field string, errors ...string)
	Validate()
	Check() Q

	Required(document M, fields ...string) M
	MinimumInt(document M, minimum int64, fields ...string) M
	MaximumInt(document M, maximum int64, fields ...string) M
	MinimumFloat(document M, minimum float64, fields ...string) M
	MaximumFloat(document M, maximum float64, fields ...string) M
	RangeInt(document M, minimum, maximum int64, fields ...string) M
	RangeFloat(document M, minimum, maximum float64, fields ...string) M
	Mobile(country string, document M, minimum, maximum float64, fields ...string) M
	Email(country string, document M, minimum, maximum float64, fields ...string) M
	PassportNumber(country string, document M, minimum, maximum float64, fields ...string) M
	IdentityNumber(country string, document M, minimum, maximum float64, fields ...string) M
	Date(country string, document M, minimum, maximum float64, fields ...string) M
	DateTime(country string, document M, minimum, maximum float64, fields ...string) M
}

type validator struct {
	Errors map[string][]string
	Lock sync.Mutex
}

func NewValidator() IValidator {
	return &validator{
		Errors: make(map[string][]string),
		Lock: sync.Mutex{},
	}
}

func RequestValidator(request M, fields ...string) IValidator {
	validator := NewValidator()

	emptyKeys := make([]string, 0)
	for key, value := range request {
		if IsEmpty(value) {
			emptyKeys = append(emptyKeys, key)
		}

		if Contains(fields, key, true) {
			if stringValue, ok := value.(string); ok && stringValue != "" {
				request[key] = strings.TrimSpace(stringValue)
			} else if stringPointerValue, ok := value.(*string); ok && stringPointerValue != nil {
				trimmedString := strings.TrimSpace(*stringPointerValue)
				request[key] = trimmedString
			}
		} else {
			validator.Error(key, "Unexpected field")
		}
	}

	return validator
}

func (v *validator) Error(field string, errors ...string) {
	if field == "" || errors == nil || len(errors) == 0 {
		return
	}

	v.Lock.Lock()
	defer v.Lock.Unlock()

	if value, exists := v.Errors[field]; !exists || value == nil {
		v.Errors[field] = []string{}
	}

	v.Errors[field] = append(v.Errors[field], errors...)
}

func (v *validator) Validate() {
	v.Lock.Lock()
	defer v.Lock.Unlock()

	if v.Errors != nil && len(v.Errors) > 0 {
		jsonObjectData, err := json.Marshal(v.Errors)
		if err != nil {
			panic(err)
		}
		jsonObjectString := string(jsonObjectData)
		validationErrorMessage := fmt.Sprintf("validation:%s", jsonObjectString)
		panic(validationErrorMessage)
	}
}

func (v *validator) Check() Q {
	v.Lock.Lock()
	defer v.Lock.Unlock()

	if v.Errors == nil {
		v.Errors = Q{}
	}
	return v.Errors
}

func (v *validator) Required(document M, fields ...string) M {
	if fields == nil || len(fields) == 0 {
		return document
	}

	for _, field := range fields {
		value := document[field]
		valid, errors := ValidateRequired(value)
		if !valid {
			if errors != nil {
				panic("invalid required validation without any errors")
			}
			v.Error(field, errors...)
		}
	}

	return document
}

func (v *validator) MinimumInt(document M, minimum int64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) MaximumInt(document M, maximum int64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) MinimumFloat(document M, minimum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) MaximumFloat(document M, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) RangeInt(document M, minimum, maximum int64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) RangeFloat(document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) Mobile(country string, document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) Email(country string, document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) PassportNumber(country string, document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) IdentityNumber(country string, document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) Date(country string, document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}

func (v *validator) DateTime(country string, document M, minimum, maximum float64, fields ...string) M {
	panic("not implemented yet")
}