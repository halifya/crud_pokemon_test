package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

type FloatArray []float64

// Implement the Valuer interface
func (a FloatArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement the Scanner interface
func (a *FloatArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	var s string
	switch v := value.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("invalid type for float_array")
	}
	return json.Unmarshal([]byte(s), a)
}

type StringArray []string

// Implement the Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement the Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	var s string
	switch v := value.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("invalid type for string_array")
	}
	return json.Unmarshal([]byte(s), a)
}

type EvolutionArray []Evolution

// Implement the Valuer interface
func (a EvolutionArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement the Scanner interface
func (a *EvolutionArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	var s string
	switch v := value.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("invalid type for evolution_array")
	}
	return json.Unmarshal([]byte(s), a)
}
