package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var bookings = make([]UserData, 0)
var remainingTickets = 50
const coferenceTickets = 50

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket( userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket( userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("these are the first names of the bookings: %v\n", firstNames)

			soldOut := remainingTickets == 0 

			if soldOut {
				fmt.Println("Sold out. Come back next year")
			}
		} else {
			if !isValidName {
				fmt.Print("First or Last name is too short\n")
			}
			if !isValidEmail {
				fmt.Print("Email address you entered does not contain @ sign\n")
			}	
			if !isValidTicketNumber {
				fmt.Printf("number of tickets you entered is invalid\n")
			}		
			
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available\n", coferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}			
	for _, booking := range bookings {
		firstNames = append(firstNames,booking.firstName)
	}
	return firstNames	
}

func getUserInput() (string, string, string, int){
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("enter your email: ")
	fmt.Scan(&email)
	fmt.Println("enter num tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
			
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email at the following address: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}

func sendTicket(userTicktes int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicktes, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("############")
	wg.Done()
}