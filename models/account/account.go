package account

import (
	"database/sql"
	"fmt"
	"persona/models"
)

type Account struct {
	ID        int     `json:"id"`
	Name      *string `json:"name" db:"name"`
	Password  *string `json:"password" db:"password"`
	CreatedAt *string `json:"created_at" db:"created_at"`
}

type Accounts []Account

func FindAll() (Accounts, error) {
	rows, err := models.MyDb.Query("SELECT * FROM accounts limit 10")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	accounts := Accounts{}
	for rows.Next() {
		var account Account
		err = rows.Scan(&account.ID, &account.Name, &account.Password, &account.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func FindByID(id int) (Account, error) {
	account := Account{}
	row := models.MyDb.QueryRow("SELECT * FROM accounts WHERE id = ?", id)
	err := row.Scan(&account.ID, &account.Name, &account.Password, &account.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Account{}, fmt.Errorf("account not found")
		}
		return Account{}, err
	}

	return account, nil
}

func Create(account Account) (Account, error) {
	err := Validate(account)
	if err != nil {
		return Account{}, err
	}

	stmt, err := models.MyDb.Prepare("INSERT INTO accounts (name, password) VALUES (?, ?)")
	if err != nil {
		return Account{}, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(account.Name, account.Password)
	if err != nil {
		return Account{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Account{}, err
	}

	account.ID = int(id)
	return account, nil
}

func Update(account Account) (Account, error) {
	err := Validate(account)
	if err != nil {
		return Account{}, err
	}

	stmt, err := models.MyDb.Prepare("UPDATE accounts SET name = ?, password = ? WHERE id = ?")
	if err != nil {
		return Account{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(account.Name, account.Password, account.ID)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

func Delete(id int) error {
	stmt, err := models.MyDb.Prepare("DELETE FROM accounts WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func FindByName(name string) (Account, error) {
	account := Account{}
	row := models.MyDb.QueryRow("SELECT * FROM accounts WHERE name = ?", name)
	err := row.Scan(&account.ID, &account.Name, &account.Password, &account.CreatedAt)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}
