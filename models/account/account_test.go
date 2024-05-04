package account

import (
	"database/sql"
	"errors"
	"persona/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func String(s string) *string {
	return &s
}

func TestFindAll(t *testing.T) {
	// Mock the database query
	mockRows := sqlmock.NewRows([]string{"id", "name", "password", "created_at"}).
		AddRow(1, "John Doe", "password123", "2022-01-01").
		AddRow(2, "Jane Smith", "password456", "2022-01-02")

	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	mock.ExpectQuery("^SELECT (.+) FROM accounts limit 10$").
		WillReturnRows(mockRows)

	// Call the function
	accounts, err := FindAll()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify the results
	if len(accounts) != 2 {
		t.Errorf("Expected 2 accounts, got: %d", len(accounts))
	}

	expectedAccount1 := Account{
		ID:        1,
		Name:      String("John Doe"),
		Password:  String("password123"),
		CreatedAt: String("2022-01-01"),
	}
	if !reflect.DeepEqual(accounts[0], expectedAccount1) {
		t.Errorf("Expected account 1 to be %+v, got: %+v", expectedAccount1, accounts[0])
	}

	expectedAccount2 := Account{
		ID:        2,
		Name:      String("Jane Smith"),
		Password:  String("password456"),
		CreatedAt: String("2022-01-02"),
	}
	if !reflect.DeepEqual(accounts[1], expectedAccount2) {
		t.Errorf("Expected account 2 to be %v, got: %v", expectedAccount2, accounts[1])
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFindAllError(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	var ErrDatabase = errors.New("database error")

	mock.ExpectQuery("^SELECT (.+) FROM accounts limit 10$").
		WillReturnError(ErrDatabase)

	// Call the function
	_, err := FindAll()
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFindByID(t *testing.T) {
	// Mock the database query
	mockRows := sqlmock.NewRows([]string{"id", "name", "password", "created_at"}).
		AddRow(1, "John Doe", "password123", "2022-01-01")

	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	mock.ExpectQuery("^SELECT (.+) FROM accounts WHERE id = \\?$").
		WithArgs(1).
		WillReturnRows(mockRows)

	// Call the function
	account, err := FindByID(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify the results
	expectedAccount := Account{
		ID:        1,
		Name:      String("John Doe"),
		Password:  String("password123"),
		CreatedAt: String("2022-01-01"),
	}
	if !reflect.DeepEqual(account, expectedAccount) {
		t.Errorf("Expected account to be %+v, got: %+v", expectedAccount, account)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFindByIDNotFound(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	mock.ExpectQuery("^SELECT (.+) FROM accounts WHERE id = \\?$").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	// Call the function
	_, err := FindByID(1)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCreate(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	mock.ExpectPrepare("^INSERT INTO accounts \\(name, password\\) VALUES \\(\\?, \\?\\)$").
		ExpectExec().
		WithArgs("John Doe", "password123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function
	account, err := Create(Account{
		Name:     String("John Doe"),
		Password: String("password123"),
	})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify the results
	expectedAccount := Account{
		ID:        1,
		Name:      String("John Doe"),
		Password:  String("password123"),
		CreatedAt: nil,
	}
	if !reflect.DeepEqual(account, expectedAccount) {
		t.Errorf("Expected account to be %+v, got: %+v", expectedAccount, account)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCreateValidation(t *testing.T) {
	// Call the function
	_, err := Create(Account{
		Name:     nil,
		Password: String("password123"),
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Call the function
	_, err = Create(Account{
		Name:     String("John Doe"),
		Password: nil,
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Call the function
	_, err = Create(Account{
		Name:     String("John Doe"),
		Password: String(""),
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestCreateError(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	var ErrDatabase = errors.New("database error")

	mock.ExpectPrepare("^INSERT INTO accounts \\(name, password\\) VALUES \\(\\?, \\?\\)$").
		ExpectExec().
		WithArgs("John Doe", "password123").
		WillReturnError(ErrDatabase)

	// Call the function
	_, err := Create(Account{
		Name:     String("John Doe"),
		Password: String("password123"),
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	mock.ExpectPrepare("^UPDATE accounts SET name = \\?, password = \\? WHERE id = \\?$").
		ExpectExec().
		WithArgs("John Doe", "password123", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function
	account, err := Update(Account{
		ID:       1,
		Name:     String("John Doe"),
		Password: String("password123"),
	})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify the results
	expectedAccount := Account{
		ID:        1,
		Name:      String("John Doe"),
		Password:  String("password123"),
		CreatedAt: nil,
	}
	if !reflect.DeepEqual(account, expectedAccount) {
		t.Errorf("Expected account to be %+v, got: %+v", expectedAccount, account)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUpdateError(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	var ErrDatabase = errors.New("database error")

	mock.ExpectPrepare("^UPDATE accounts SET name = \\?, password = \\? WHERE id = \\?$").
		ExpectExec().
		WithArgs("John Doe", "password123", 1).
		WillReturnError(ErrDatabase)

	// Call the function
	_, err := Update(Account{
		ID:       1,
		Name:     String("John Doe"),
		Password: String("password123"),
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUpdateValidation(t *testing.T) {
	// Call the function
	_, err := Update(Account{
		ID:       1,
		Name:     nil,
		Password: String("password123"),
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Call the function
	_, err = Update(Account{
		ID:       1,
		Name:     String("John Doe"),
		Password: nil,
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Call the function
	_, err = Update(Account{
		ID:       1,
		Name:     String("John Doe"),
		Password: String(""),
	})
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestDelete(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	mock.ExpectPrepare("^DELETE FROM accounts WHERE id = \\?$").
		ExpectExec().
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function
	err := Delete(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestDeleteError(t *testing.T) {
	// Mock the database query
	mockDB, mock, _ := sqlmock.New()
	models.MyDb = mockDB
	defer mockDB.Close()

	var ErrDatabase = errors.New("database error")

	mock.ExpectPrepare("^DELETE FROM accounts WHERE id = \\?$").
		ExpectExec().
		WithArgs(1).
		WillReturnError(ErrDatabase)

	// Call the function
	err := Delete(1)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
