package main
import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50 

	fmt.Println("Welcome to our", conferenceName, "booking application")
	//printf -> format, %v to indicate place holder
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2

	fmt.Printf("%v booked %v tickets.\n", userName, userTickets)
}//46:20