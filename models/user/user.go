package user

import "persona/models"

type User struct {
	ID        int     `json:"id"`
	FirstName *string `json:"first_name" db:"first_name"`
	LastName  *string `json:"last_name" db:"last_name"`
	Email     *string `json:"email" db:"email"`
}

type Users []User

func FindAll() Users {
	users, err := models.MyDb.Query("SELECT id, first_name, last_name, email FROM users limit 10")
	if err != nil {
		panic(err.Error())
	}

	u := Users{}
	for users.Next() {
		var user User
		err = users.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		u = append(u, user)
	}

	return u
}

func FindByID(id int) User {
	user := User{}
	row := models.MyDb.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		panic(err.Error())
	}

	return user
}

func (u *Users) AddUser(user User) Users {
	*u = append(*u, user)
	return *u
}
