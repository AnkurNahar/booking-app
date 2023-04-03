package main
import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 
	//var bookings = [50]string{"Nina", "Rossi", "Anks"}
	var bookings [50]string //alternative declaration syntax for empty array

	fmt.Println("Welcome to our", conferenceName, "booking application")
	//printf -> format, %v to indicate place holder
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	i := 0
	bookings[i] = firstName + " " + lastName
	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will recieve a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Println(bookings[i])
}//46:20