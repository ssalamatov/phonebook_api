package user

import "fmt"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) String() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func NewUser(firstName string, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}
