package main

import (
	"strings"
)

//Splitting the logic here

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}
