package main

import "fmt"

func main() {
	var conferenceName = "Golang" // changeable, also "var" can be declared as :=
	const conferenceTickets = 87  // not changeable, also "const" can't be declared as :=
	var remainingTickets = 20

	fmt.Printf("Welcome to %v booking application\n\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string // if we don't assign value to variable, Go doesn't understand the type
	var email string
	var userTickets int

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName) // in user input we must use pointer "&", which points to memory address of var

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v for booking %v tickets. Email was sent to %v", userName, userTickets, email)
}
