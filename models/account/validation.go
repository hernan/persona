package account

import (
	"fmt"
	"strings"
)

type Constraint func() error

type Validator struct {
	constraints []Constraint
}

func (validator *Validator) Validate() error {
	errors := []string{}

	for _, constraint := range validator.constraints {
		err := constraint()
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	if len(errors) == 0 {
		return nil
	}

	return fmt.Errorf(strings.Join(errors, ", "))
}

func (validator *Validator) AddConstraint(constraint Constraint) {
	validator.constraints = append(validator.constraints, constraint)
}

func NotEmptyConstraint(value *string, fieldName string) Constraint {
	return func() error {
		if value == nil || strings.TrimSpace(*value) == "" {
			return fmt.Errorf("%s cannot be empty", fieldName)
		}
		return nil
	}
}

func PositiveNonZeroConstraint(value int, fieldName string) Constraint {
	return func() error {
		if value <= 0 {
			return fmt.Errorf("%s must be positive and non-zero", fieldName)
		}
		return nil
	}
}

func Validate(account Account) error {
	validator := Validator{}

	validator.AddConstraint(NotEmptyConstraint(account.Name, "Name"))
	validator.AddConstraint(NotEmptyConstraint(account.Password, "Password"))

	return validator.Validate()
}
