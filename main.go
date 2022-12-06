package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Golang" // changeable, also "var" can be declared as :=
	const conferenceTickets = 87  // not changeable, also "const" can't be declared as :=
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	// endless loop
	for {
		var bookings [50]string // array in Go is fixed sized, and we can store in it only one type of data
		var bookings2 []string  // this is Slice, which uses array under the hood but Slice is dynamic
		var userName string     // if we don't assign value to variable, Go doesn't understand the type
		var email string
		var userTickets uint

		fmt.Print("Enter your name: ")
		fmt.Scan(&userName) // in user input we must use pointer "&", which points to memory address of var

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("How many tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings[0] = userName                  // adding to an array
		bookings2 = append(bookings2, userName) // adding to a slice

		fmt.Printf("Thank you %v for booking %v tickets. Email was sent to %v", userName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

		firstNames := []string{}
		for _, booking := range bookings2 {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}
