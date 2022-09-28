package main

import "fmt"

func main() {
	var eventName = "Go Events"
	const eventTickets = 100
	var remainingTickets = 50

	fmt.Println("Welcome to", eventName, "booking application")
	fmt.Println("We have total of", eventTickets, "tickets and", remainingTickets, "are still avaliable")
	fmt.Println("Get your tickets here to attend")
}
