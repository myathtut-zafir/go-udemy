package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (appUser User) PrintUserDetail() {
	fmt.Println(appUser.createdAt, appUser.firstName)
}
func (appUser *User) ClearUserInfo() {
	appUser.firstName = ""
	appUser.lastName = ""
	appUser.birthdate = ""
}

// util function to create a new user
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
