package main

import (
	"fmt"
	"strings"
)

func validateUserInputs() bool {
	valid := true
	isValidFirstName := len(firstName) >= 2
	if !isValidFirstName {
		fmt.Print("Name is too short\n\n")
		valid = false
	}

	
	isValidLastName := len(lastName) >= 2
	if !isValidLastName {
		fmt.Print("Name is too short\n\n")
		valid = false
	}
	
	
	isValidEmail := strings.Contains(userEmail, "@")
	if !isValidEmail {
		fmt.Print("Email address does not contain @\n\n")
		valid = false
	}

	
	isValidTicketNumber := userTickets <= remainingTickets
	if !isValidTicketNumber{
		fmt.Printf("%v tickets remaining for %v, so you cannot book %v tickets\n\n", remainingTickets, conferenceName, userTickets)
		valid = false
	}

return valid
}