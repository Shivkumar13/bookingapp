package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for their name

		fmt.Println("Enter your first name:")
		fmt.Scanf("%s", &firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole array is: %v\n", bookings)
		fmt.Printf("The first value is: %v\n", bookings[0])

		fmt.Printf("The Array Type: %T\n", bookings)
		fmt.Printf("The Array length: %v\n", len(bookings))

		fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		//var names []string

		for _, booking := range bookings {

			var names = strings.Fields(booking)

			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are : %v\n", firstNames)

	}

}
