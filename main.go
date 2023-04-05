package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	//var bookings = [50]string{"Nina", "Rossi", "Anks"}
	//var bookings [50]string //alternative declaration syntax for empty array
	var bookings []string //alternative syntax: var bookings = []string{}, bookings := []string{}

	fmt.Println("Welcome to our", conferenceName, "booking application")
	//printf -> format, %v to indicate place holder
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for{ //infinite loop

		//ticketsNotAvailable := remainingTickets == 0 -> can ticketsNotAvailable this as condition
		if remainingTickets == 0 {
			fmt.Println("All tickets sold out! Please come back next year")
			break
		}

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

		if userTickets > remainingTickets{
			fmt.Printf("%v tickets remaining for %v, so you cannot book %v tickets\n", remainingTickets, conferenceName, userTickets)
			continue
		}
		remainingTickets -= userTickets
		//i := 0
		//bookings[i] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will recieve a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names for bookings: %v\n", firstNames)
	}
}