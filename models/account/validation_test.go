package account

import (
	"testing"
)

func TestValidate(t *testing.T) {
	name := "John Doe"
	password := "password123"
	emptyString := ""

	// Test case 1: Valid account
	account := Account{
		Name:     &name,
		Password: &password,
	}

	err := Validate(account)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Test case 2: Missing name
	account = Account{
		Password: &password,
	}

	err = Validate(account)
	if err == nil {
		t.Error("Expected error, got nil")
	} else if err.Error() != "name is required" {
		t.Errorf("Expected error 'name is required', got: %v", err)
	}

	// Test case 3: Missing password
	account = Account{
		Name: &name,
	}

	err = Validate(account)
	if err == nil {
		t.Error("Expected error, got nil")
	} else if err.Error() != "password is required" {
		t.Errorf("Expected error 'password is required', got: %v", err)
	}

	// Test case 4: Password is empty
	account = Account{
		Name:     &name,
		Password: &emptyString,
	}

	err = Validate(account)
	if err == nil {
		t.Error("Expected error, got nil")
	} else if err.Error() != "password is required" {
		t.Errorf("Expected error 'password is required', got: %v", err)
	}
}
