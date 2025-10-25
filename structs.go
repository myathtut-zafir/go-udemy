package main

import (
	"fmt"
	"github.com/pointers/user"
)

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
	appUser, error := user.New(firstName, lastName, birthdate)
	if error != nil {
		fmt.Println("Error creating user:", error)
		return
	}
	appUser.PrintUserDetail()
	appUser.ClearUserInfo()
	appUser.PrintUserDetail()
	// printUser(appUser)
	// printUserPointer(&appUser)

	// ... do something awesome with that gathered data!

	// fmt.Println(appUser.createdAt)
}

// func printUser(appUser user) {
// 	fmt.Println(appUser.createdAt, appUser.firstName)
// }
// func printUserPointer(appUser *user) {
// 	fmt.Println(appUser.createdAt, appUser.firstName)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
