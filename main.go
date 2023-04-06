package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	var bookings []string //alternative syntax: var bookings = []string{}, bookings := []string{}

	for{ //infinite loop

		greetUser(conferenceName, conferenceTickets, remainingTickets)

		firstName, lastName, userEmail, userTickets := getUserInput()

		valid := validateUserInputs(firstName, lastName, userEmail, userTickets, remainingTickets, conferenceName)
		if(!valid){
			continue
		}
		remainingTickets -= userTickets

		bookings = bookTickets(bookings, firstName, lastName, userTickets, userEmail, remainingTickets, conferenceName)

		printBookings(bookings)

		if remainingTickets == 0 {
			fmt.Println("All tickets sold out! Please come back next year")
			break
		}
	}
}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint){
	fmt.Println("Welcome to our", conferenceName, "booking application")
	//printf -> format, %v to indicate place holder
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput()(string, string, string, uint){
	var firstName string
	var lastName string
	var userEmail string		
	var userTickets uint

	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email address: ")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func validateUserInputs(firstName string, lastName string, userEmail string, userTickets uint, remainingTickets uint, conferenceName string) bool {
		valid := true
		isValidFirstName := len(firstName) >= 2
		if !isValidFirstName {
			fmt.Print("Name is to short\n\n")
			valid = false
		}

		
		isValidLastName := len(lastName) >= 2
		if !isValidLastName {
			fmt.Print("Name is to short\n\n")
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

func bookTickets(bookings []string, firstName string, lastName string, userTickets uint, userEmail string, remainingTickets uint, conferenceName string) []string{
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will recieve a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return bookings
}

func printBookings(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {			
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names for bookings: %v\n\n", firstNames)
}