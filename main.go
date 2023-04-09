package main

import (
	"fmt"
)

//package level variables
var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50 
var bookings = make([]UserData, 0)//empty slice of maps
var firstName string
var lastName string
var userEmail string
var userTickets uint

type UserData struct {
	firstName string
	lastName string
	userEmail string
	numberOfTickets uint
}

func main(){

	for{ //infinite loop

		greetUser()

		firstName, lastName, userEmail, userTickets = getUserInput()

		valid := validateUserInputs()
		if(!valid){
			continue
		}

		bookTickets()

		printBookings()

		if remainingTickets == 0 {
			fmt.Println("All tickets sold out! Please come back next year")
			break
		}
	}
}

func greetUser(){
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

func bookTickets() {
	remainingTickets -= userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		userEmail: userEmail,
		numberOfTickets: userTickets, 
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will recieve a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func printBookings() {
	firstNames := []string{}
	for _, booking := range bookings {					
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("List of bookings: %v\n\n", bookings)
}