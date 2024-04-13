package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Users []User

var UserDB = Users{
	User{ID: 1, Name: "Jon Snow", Email: "jon.snow@kingdom.com", Password: "123"},
	User{ID: 2, Name: "Tyrion Lannister", Email: "tyron.lannister@kingdom.com", Password: "123"},
	User{ID: 3, Name: "Daenerys Targaryen", Email: "deanerys.targaryen@kingdom.com", Password: "123"},
	User{ID: 4, Name: "Arya Stark", Email: "Arya.Stark@kingdom.com", Password: "123"},
}

func (u *Users) FindAll() Users {
	return *u
}

func (u *Users) FindByID(id int) *User {
	for _, user := range *u {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (u *Users) AddUser(user User) Users {
	*u = append(*u, user)
	return *u
}
