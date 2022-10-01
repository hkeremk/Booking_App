package main

import (
	"BookingApp/helper"
	"fmt"
	"strings"
)

var eventName = "Go Events"

const eventTickets int = 100

var remainingTickets uint = 100
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUSerInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUSerInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our event is booked out. Come back for the next event.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", eventName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable\n", eventTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUSerInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}
