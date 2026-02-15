package models

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
