package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123456789"
	hashedpassword, err := hashPassword(password)
	if err != nil {
		fmt.Println("Cannot generate password:", err)
	}
	log.Println(string(hashedpassword))
	password = "12345678"
	err = comparePassword([]byte(password), hashedpassword)
	if err != nil {
		fmt.Println("Invalid password:", err)
	} else {
		log.Println("Succesfully logged in!")
	}
}
func hashPassword(password string) ([]byte, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedpassword, nil
}
func comparePassword(password, hashedpassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedpassword, password)
	if err != nil {
		return err
	}
	return nil
}
