package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateTickets(tickets string) (bool, uint) {
	// Convert string to uint.
	uint64Value, err := strconv.ParseUint(tickets, 10, 0) // 10 is the base, 0 for auto size
	if err != nil {
		fmt.Println("Validation Error: Positive number is required.")
		return false, 0
	}
	return true, uint(uint64Value)
}

func ValidateName(name string) (bool, string) {
	if len(name) <= 2 {
		fmt.Println("Validation Error: Names must be at least 2 characters.")
		return false, ""
	}
	return true, name
}

func ValidateEmail(email string) (bool, string) {
	if !strings.Contains(email, "@") {
		fmt.Println("Validation Error: Email must contain @.")
		return false, ""
	}
	return true, email
}
