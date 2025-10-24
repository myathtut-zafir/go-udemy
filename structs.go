package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (appUser user) printUserDetail() {
	fmt.Println(appUser.createdAt, appUser.firstName)
}
func (appUser *user) clearUserInfo() {
	appUser.firstName = ""
	appUser.lastName = ""
	appUser.birthdate = ""
}

// util function to create a new user
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user
	// appUser := user{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }
	appUser, error := newUser(firstName, lastName, birthdate)
	if error != nil {
		fmt.Println("Error creating user:", error)
		return
	}
	appUser.printUserDetail()
	appUser.clearUserInfo()
	appUser.printUserDetail()
	// printUser(appUser)
	// printUserPointer(&appUser)

	// ... do something awesome with that gathered data!

	// fmt.Println(appUser.createdAt)
}

func printUser(appUser user) {
	fmt.Println(appUser.createdAt, appUser.firstName)
}
func printUserPointer(appUser *user) {
	fmt.Println(appUser.createdAt, appUser.firstName)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
