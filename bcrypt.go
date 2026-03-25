package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Hvilket passord ønsker du?")

	var password string

	fmt.Scanln(&password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hashed: ", string(hashedPassword))

	fmt.Println("Passord?")

	var guess string

	fmt.Scanln(&guess)

	match := bcrypt.CompareHashAndPassword(hashedPassword, []byte(guess))

	if match == nil {
		fmt.Println("Match")
	} else {
		fmt.Println("No match")
	}
}
