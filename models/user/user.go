package user

import "persona/models"

type User struct {
	ID        int     `json:"id"`
	Name      *string `json:"name" db:"name"`
	Email     *string `json:"email" db:"email"`
	Phone     *string `json:"phone" db:"phone"`
	BirthDay  *string `json:"birthday" db:"birthday"`
	CreatedAt *string `json:"created_at" db:"created_at"`
}

type Users []User

func FindAll() (Users, error) {
	rows, err := models.MyDb.Query("SELECT * FROM users limit 10")
	if err != nil {
		return nil, err
	}

	users := Users{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.BirthDay, &user.CreatedAt)
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
	row := models.MyDb.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.BirthDay, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func Create(user User) (User, error) {
	stmt, err := models.MyDb.Prepare("INSERT INTO users (name, email, phone, birthday) VALUES (?, ?, ?, ?)")
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Email, user.Phone, user.BirthDay)
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
	stmt, err := models.MyDb.Prepare("UPDATE users SET name = ?, email = ?, phone = ?, birthday WHERE id = ?")
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Phone, user.BirthDay, user.ID)
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
