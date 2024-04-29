package user

import "persona/models"

type User struct {
	ID        int     `json:"id"`
	FirstName *string `json:"first_name" db:"first_name"`
	LastName  *string `json:"last_name" db:"last_name"`
	Email     *string `json:"email" db:"email"`
	Password  *string `json:"-" db:"password"`
}

type Users []User

func FindAll() (Users, error) {
	rows, err := models.MyDb.Query("SELECT id, first_name, last_name, email FROM users limit 10")
	if err != nil {
		return nil, err
	}

	users := Users{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	rows.Close()
	return users, nil
}

func FindByID(id int) (User, error) {
	user := User{}
	row := models.MyDb.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func Create(user User) (User, error) {
	stmt, err := models.MyDb.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return User{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}

	user.ID = int(id)

	return user, nil
}

func Update(user User) (User, error) {
	stmt, err := models.MyDb.Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?")
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func Delete(id int) error {
	stmt, err := models.MyDb.Prepare("DELETE FROM users WHERE id = ?")
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
