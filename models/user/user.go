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
	users, err := models.MyDb.Query("SELECT id, first_name, last_name, email FROM users limit 10")
	if err != nil {
		return nil, err
	}

	u := Users{}
	for users.Next() {
		var user User
		err = users.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		u = append(u, user)
	}

	return u, nil
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

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
