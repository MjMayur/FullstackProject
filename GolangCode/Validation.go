package main

import (
	"errors"
	"regexp"
)

func validationEmail(email string) error {
	// Regular expression for validating an email

	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// func validationPass(password string) error {
// 	// Regular expression for validating an email

// 	var passRezex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`
// 	re := regexp.MustCompile(passRezex)
// 	fmt.Println("-------> Its Working")
// 	if !re.MatchString(password) {
// 		return errors.New("invalid email format")
// 	}
// 	return nil
// }
