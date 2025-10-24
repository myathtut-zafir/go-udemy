package main

import "fmt"

func main() {
	age := 30
	agePointer := &age

	// fmt.Println("Value of age:", agePointer)
	// fmt.Println("Value of age value:", *agePointer)
	// fmt.Println("Value of age2:", updateAge(agePointer))
	fmt.Println("Value of age:", age)
	updateAge2(agePointer)
	fmt.Println("Value of age after updateAge2:", age)
}

func updateAge(age *int) int {
	return *age - 18
}
func updateAge2(age *int) {
	*age = *age - 18
}
