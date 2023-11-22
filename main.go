package main

import "fmt"

//Expressions - вирази
func main() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded)
	fmt.Println()

	ageInTenYears := age + 10 // This is Expressions - вирази
	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears)
	fmt.Println()

	isATeenager := age >= 13 // This is Expressions - вирази
	fmt.Println("is a teenager:", isATeenager)

}
