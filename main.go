package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidEmail, isValidName, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidTicketNumber && isValidEmail {

			bookTicket(userTickets, email, firstName, lastName)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {

				//end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		} else {

			if !isValidName {

				fmt.Println("First Name or last name you entered is too short")

			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				firstNames := []string{}
				//var names []string
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])
				}

			}
		}

	}
}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {

	firstNames := []string{}
	var names []string
	for _, booking := range bookings {

		names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, email string, firstName string, lastName string) []string {

	remainingTickets = remainingTickets - userTickets

	bookings := append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v, you have booked %v tickets  successfully. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}
