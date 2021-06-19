/* 
	Assignment 3 - The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

	Author: Ruymán Borges
	Date: 13-6-21
*/


package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput, formattedUserInput string
	var prefixedI, containsA, suffixedN bool

	fmt.Printf("\nInstructions: You will enter a string and I'll tell you 'Found!' if it begins with 'i', contains 'a' among the middle chars, and ends with 'n'\n")

	fmt.Printf("\nEnter a string and press ENTER\n\n")

	fmt.Scan(&userInput)

	formattedUserInput = strings.ToLower(strings.TrimSpace(userInput))

	prefixedI = strings.HasPrefix(formattedUserInput, "i")
	containsA = strings.Contains(formattedUserInput, "a")
	suffixedN = strings.HasSuffix(formattedUserInput, "n")

	if ( prefixedI && containsA && suffixedN) {
		fmt.Printf("\nFound!\n")
	} else {
		fmt.Printf("\nNot Found!\n")
	}

}