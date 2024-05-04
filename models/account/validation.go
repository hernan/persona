package account

import "errors"

// TODO: Implement a similar validation to rails models, should return an array of errors
// is some of the conditions are not met
func Validate(account Account) error {
	if account.Name == nil {
		return errors.New("name is required")
	}

	if account.Password == nil || *account.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
