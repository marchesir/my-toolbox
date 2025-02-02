package main

import (
	"bookit/helper"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {
	const eventName = "Gophers Conference"
	const eventTickets uint = 100
	remainingTickets := eventTickets
	var bookings = make([]UserData, 0)

	fmt.Printf("\n***************************************************************************************\n")
	fmt.Printf("Welcome to BookItNow application configuered for the %v\n", eventName)
	fmt.Printf("Total tickets available for the event: %v\n", eventTickets)
	fmt.Printf("***************************************************************************************\n\n")

	//var bookings []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nHow many tickets would you like to book:__")
		scanner.Scan()
		ticketsOk, tickets := helper.ValidateTickets(strings.TrimSpace(scanner.Text()))
		if !ticketsOk {
			fmt.Println("Booking failed, please try again.")
			continue
		}

		fmt.Print("Enter ticket holders first name:__: ")
		scanner.Scan()
		firstNameOk, firstName := helper.ValidateName(strings.TrimSpace(scanner.Text()))
		if !firstNameOk {
			fmt.Println("Booking failed, please try again.")
			continue
		}

		fmt.Print("Enter ticket holders last name:__: ")
		scanner.Scan()
		secondNameOk, lastName := helper.ValidateName(strings.TrimSpace(scanner.Text()))
		if !secondNameOk {
			fmt.Println("Booking failed, please try again.")
			continue
		}

		fmt.Print("Enter ticket holders email:__: ")
		scanner.Scan()
		emailOk, email := helper.ValidateEmail(strings.TrimSpace(scanner.Text()))
		if !emailOk {
			fmt.Println("Booking failed, please try again.")
			continue
		}

		var userData = UserData{firstName, lastName, email, tickets}

		if tickets > remainingTickets {
			fmt.Print("Sorry, we have sold out of tickets.\n")
			fmt.Println("Exiting...")
			break
		}
		remainingTickets = remainingTickets - tickets
		bookings = append(bookings, userData)
		sendTicket(tickets, firstName, lastName, email)

		var bookingsOutput []string
		for _, booking := range bookings {
			bookingsOutput = append(bookingsOutput, fmt.Sprintf("tickets:%v-%v", booking.tickets, booking.email))
		}

		fmt.Printf("Current bookings: %v\n", bookingsOutput)
		// Ask if the user wants to continue.
		fmt.Print("Do you want to continue (yes/no)? ")
		scanner.Scan()
		if scanner.Text() == "no" {
			fmt.Println("Exiting...")
			break
		}
	}
}

func sendTicket(tickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets to %v %v: ", tickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################################")
}
