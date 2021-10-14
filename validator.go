package uniform

// A definition of the public functions for a request interface
type IValidator interface {
	Error(field, error string)
	Check()
	SilentCheck() map[string][]string

	Required(document map[string]interface{}, fields ...string) map[string]interface{}
	MinimumInt(document map[string]interface{}, minimum int64, fields ...string) map[string]interface{}
	MaximumInt(document map[string]interface{}, maximum int64, fields ...string) map[string]interface{}
	MinimumFloat(document map[string]interface{}, minimum float64, fields ...string) map[string]interface{}
	MaximumFloat(document map[string]interface{}, maximum float64, fields ...string) map[string]interface{}
	RangeInt(document map[string]interface{}, minimum, maximum int64, fields ...string) map[string]interface{}
	RangeFloat(document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
	Mobile(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
	Email(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
	PassportNumber(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
	IdentityNumber(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
	Date(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
	DateTime(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{}
}

type validator struct {
	Errors map[string][]string
}

func NewValidator() IValidator {
	return &validator{
		Errors: make(map[string][]string),
	}
}

func (v validator) Error(field, error string) {
	panic("not implemented yet")
}

func (v validator) Check() {
	panic("not implemented yet")
}

func (v validator) SilentCheck() map[string][]string {
	panic("not implemented yet")
}

func (v validator) Required(document map[string]interface{}, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) MinimumInt(document map[string]interface{}, minimum int64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) MaximumInt(document map[string]interface{}, maximum int64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) MinimumFloat(document map[string]interface{}, minimum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) MaximumFloat(document map[string]interface{}, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) RangeInt(document map[string]interface{}, minimum, maximum int64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) RangeFloat(document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) Mobile(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) Email(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) PassportNumber(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) IdentityNumber(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) Date(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}

func (v validator) DateTime(country string, document map[string]interface{}, minimum, maximum float64, fields ...string) map[string]interface{} {
	panic("not implemented yet")
}