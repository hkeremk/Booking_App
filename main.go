package main

import "fmt"

func main() {
	eventName := "Go Events"
	const eventTickets int = 100
	var remainingTickets uint = 100
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", eventName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable\n", eventTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventTickets)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
