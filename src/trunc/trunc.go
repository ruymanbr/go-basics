/* 
	Assignment 2 - program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

	Author: Ruymán Borges
	Date: 13-6-21
*/

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var floatUserNumber float64
	var floatUserNumber2 float64
	var loopCommandInput string
	var oneOrTwoInputs int
	

	for {

		fmt.Printf("\nAre you going to enter 1 or 2 float numbers? (Enter 1 or 2 and press enter) ")
		fmt.Scan(&oneOrTwoInputs)

		if (oneOrTwoInputs == 1) {
			fmt.Printf("\nEnter 1 float number and press enter: \n")
			fmt.Scan(&floatUserNumber)

			truncatedNumber := int(math.Trunc(floatUserNumber))
			
			switch {
				case &truncatedNumber != nil:
					fmt.Printf("\nYour truncated integer is: %d", truncatedNumber)		
				default:
					fmt.Printf("\nYou didn't enter a number or something went wrong. Try again\n")
					continue
			}

			fmt.Printf("\n(q) and enter to quit -- Any other key and enter to start over\n")
			fmt.Scan(&loopCommandInput)

			if (strings.TrimSpace(loopCommandInput) == "q" || strings.TrimSpace(loopCommandInput) == "Q") {

				break

			} else {

				continue

			}
		} 

		if (oneOrTwoInputs == 2) {
			fmt.Printf("\nEnter 1 or 2 float numbers separated by space and press enter: \n")
			fmt.Scan(&floatUserNumber)
			fmt.Scan(&floatUserNumber2)

			truncatedNumber := int(math.Trunc(floatUserNumber))
			truncatedNumber2 := int(math.Trunc(floatUserNumber2))

			switch {
				case &truncatedNumber2 != nil && &truncatedNumber != nil:
					fmt.Printf("\nYour truncated integers are: %d %d", truncatedNumber, truncatedNumber2)
				case &truncatedNumber2 == nil && &truncatedNumber != nil:
					fmt.Printf("\nYour truncated integer is: %d", truncatedNumber)		
				default:
					fmt.Printf("\nYou didn't enter a number or something went wrong. Try again\n")
					fmt.Printf("Enter 1 or 2 float numbers separated by 1 space and press enter: \n")
			}

			fmt.Printf("\n(q) and enter to quit -- Any other key and enter to start over\n")
			fmt.Scan(&loopCommandInput)

			if (strings.TrimSpace(loopCommandInput) == "q" || strings.TrimSpace(loopCommandInput) == "Q") {

				break

			} else {

				continue

			}
		} else {
			fmt.Printf("\nSomething went wrong or I didn't understand you. Please try again")
		}

		
	}


	
}