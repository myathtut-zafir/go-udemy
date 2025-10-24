package main

import "fmt"

func main() {
	age := 30
	agePointer := &age

	fmt.Println("Value of age:", agePointer)
	fmt.Println("Value of age value:", *agePointer)
	fmt.Println("Value of age2:", updateAge(age))
}

func updateAge(age int) int {
	return age - 18
}
