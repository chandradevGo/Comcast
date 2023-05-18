package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isValidCreditCardNumber(credit string) string {
	creditRemovedHyphen := strings.ReplaceAll(credit, "-", "")

	// Check for 16-digit format
	length16 := regexp.MustCompile(`^[4-6]\d{15}$`).MatchString(credit)
	// Check for 19-digit format
	length19 := regexp.MustCompile(`^[4-6]\d{3}-\d{4}-\d{4}-\d{4}$`).MatchString(credit)
	// Check for consecutive repeated digits
	consecutive := regexp.MustCompile(`(\d)\1\1\1`).MatchString(creditRemovedHyphen)

	if length16 || length19 {
		if consecutive {
			return "Invalid"
		}
	} else {
		return "Invalid"
	}

	return "Valid"
}

func main() {
	var n int
	fmt.Scanln(&n)

	cardNumbers := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&cardNumbers[i])
	}

	for _, cardNumber := range cardNumbers {
		fmt.Println(isValidCreditCardNumber(cardNumber))
	}
}
