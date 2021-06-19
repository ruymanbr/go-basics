/* 
 * File:    slice.go 
 * 
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com) 
 * Date:    	18-6-21
 * Course:  	Starting with Go (University of California IRVINE)
 * Assignment:  Module 3 - Activity slice.go
 * 
 * Summary of File: 
 * 
 *  The program slice.go prompts the user to enter integers and stores the integers in a sorted slice. 
 *	The program should be written as a loop. Before entering the loop, the program should create 
 * 	an empty integer slice of size (length) 3. During each pass through the loop, 
 *	the program prompts the user to enter an integer to be added to the slice. 
 * 	The program adds the integer to the slice, sorts the slice, and prints the contents of the 
 *	slice in sorted order. The slice must grow in size to accommodate any number of integers 
 *	which the user decides to enter. The program should only quit (exiting the loop) 
 *	when the user enters the character ‘X’ instead of an integer.
 * 
 */ 

package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {

	var s1 = make([]int, 3, 4)
	var userInput string
	var exitProgram bool = false

	
	fmt.Println("You will enter 1 integer at a time and the program will show you the sorted slice at every step.")

	exitProgram = insertFirstThreeIntegers(s1, userInput)	
		


	for exitProgram == false {
		fmt.Println("You can keep adding integers to the slice. Enter (X) to EXIT: ")

		exitProgram, s1 = increasingSlice(s1, userInput)
	}

	
}


/* 
 * function insertFirstThreeIntegers (sliceX []int, uInput string) bool
 *  
 * Function:	insertFirstThreeIntegers 
 *  
 * Description:	User is prompted to enter 3 digits. Function enters it in a predefined slice with 
 * 				underlying array of size 4. Values are added in input order to the slice. Function calls *				another function printSlice which will sort and print the sorted slice
 *  
 * Parameters:	[]int, string - a Slice of integers and a string
 *  
 * Return:		bool
 *  
 * Examples of Usage: 
 *  
 *  >> var s1 = make([]int, 3, 4)
 *  >> var userInput string
 *  >> insertFirstThreeIntegers(s1, userInput)
 *  
 *  (for loops for range of slice s1, which is 3 in this case ans user inputs 3 numbers)
 *  e.g.: 3, 7 and 2
 * 
 *  
 *
 *  ans:
 * 			[ 3 7 2 ]  	// The actual slice after input
 *			[ 2 3 7 ]  	// Printed version on screen after sorting it
 *
 *	return: true 		// After all 3 integers are entered. "false" if "x" is entered through the
 *						// process
 * 
 *
 */
func insertFirstThreeIntegers(sliceX []int, uInput string) bool {

	exitProgram := false

	for i, _ := range sliceX {



		fmt.Println("Enter 1 integer now and press ENTER: enter (X) to EXIT")
		fmt.Scan(&uInput)

		if !(checkIfNumber(uInput)) {
			if (evalInput(uInput) == "x") {
				fmt.Println("You have successfully EXITED the program. Have a nice day!")
				exitProgram = true
				
				break
			}

			for {
				fmt.Println("It MUST be an integer. Please try again or enter (X) to EXIT")
				fmt.Scan(&uInput)

				if (checkIfNumber(uInput)) {
					intInput, _ := strconv.Atoi(uInput)
					sliceX[i] = intInput
					break
				}
			}
			
		} else {
			intInput, _ := strconv.Atoi(uInput)
			sliceX[i] = intInput
		}
		printSlice(sliceX)

	}

	return exitProgram
}


/* 
 * function increasingSlice (sliceX []int, uInput string) (bool, []int)
 *  
 * Function:	increasingSlice 
 *  
 * Description:	User is prompted to enter 1 more digit recursively until "x" or "X" is entered at any 
 *				time to exit. It includes the input in original slice at every step of the loop.
 * 				The new value is added in the slice at the end of it. 
 * 				Function calls another function printSlice which will sort and print the sorted slice
 *  
 * Parameters:	[]int, string 	- a Slice of integers and a string
 *  
 * Return:		bool, []int 	- true/false and a slice of integers
 *  
 * Examples of Usage: 
 *  
 *  >> s1 := []int{ 3, 7, 2}
 *  >> var userInput string
 *  >> increasingSlice(s1, userInput)
 *  
 *  (function checks for number validity and then adds it at the end. Then calls printSlice which sorts *		and print the sorted slice)
 *  e.g.: 	77
 *			22
 *			4
 * 
 *  
 *
 *  ans:
 * 			[ 3 7 2 77 22 4]  // The actual slice after input
 *			[ 2 3 4 7 22 77]  // Printed version on screen after sorting it *
 *	
 *
 *	return: false 		// Until "x" is entered. Then return: true (and exits)
 * 
 *
 */
func increasingSlice(sliceX []int, uInput string) (bool, []int) {

	fmt.Println("Enter 1 integer now and press ENTER: enter (X) to EXIT")
	fmt.Scan(&uInput)

	
	if !(checkIfNumber(uInput)) {
		if evalInput(uInput) == "x" {
			fmt.Println("You have successfully EXITED the program. Have a nice day!")
			return true, sliceX;
		}

		for {
			fmt.Println("It MUST be an integer. Please try again or enter (X) to EXIT")
			fmt.Scan(&uInput)

			intInput, _ := strconv.Atoi(uInput)
			
			if checkIfNumber(uInput) {
				sliceX = append(sliceX, intInput)
				break
			}
		}
		
	} else {
		intInput, _ := strconv.Atoi(uInput)
		sliceX = append(sliceX, intInput)
	}
	printSlice(sliceX)

	return false, sliceX;
}

/* 
 * function evalInput (uInput string) string
 *  
 * Function:	evalInput 
 *  
 * Description:	Function lower case a string and returns it
 *  
 * Parameters:	string
 *  
 * Return:		string	- a lowercase string version
 *  
 * Examples of Usage: 
 *  
 *  >> userInput := "X"
 *  >> evalInput(uInput)
 *  
 *
 *  ans:
 * 			"x"  	// The lower case string
 *
 *	return: "x" 	
 * 
 *
 */
func evalInput(uInput string) string {
	
	return strings.ToLower(strings.TrimSpace(uInput))
}

/* 
 * function sortSlice(slice []int)
 *  
 * Function:	sortSlice 
 *  
 * Description:	Function gets an integer slice and sorts it in ascendant order
 * 				It doesn't return anything. It just orders the slice
 *  
 * Parameters:	[]int 		- a Slice of integers
 *  
 * Return:		
 *  
 * Examples of Usage: 
 *  
 *  >> s1 := []int{ 3, 7, 2}
 *  >> sortSlice(s1)
 *  
 *  (function loops through the slice ordering its elements)
 *  
 *
 *  ans:
 *			[ 2 3 7 ]   // Not printed in screen
 *	
 *
 *	return: 	 		// No return
 * 
 *
 */
func sortSlice(tempSliceX []int) {
	
	sort.Slice(tempSliceX, func(i, j int) bool {
	    return tempSliceX[i] < tempSliceX[j]	    
	})
}

/* 
 * function printSlice (slice []int)
 *  
 * Function:	printSlice 
 *  
 * Description:	Function copy a slice into a temporary slice to prevent modifying the original one. 
 *				Then it calls sortSlice function using the temporary slice and prints it
 *  
 * Parameters:	[]int 		- a Slice of integers
 *  
 * Return:				 	- No return
 *  
 * Examples of Usage: 
 *  
 *  >> s1 := []int{ 3, 7, 2}
 *  >> printSlice(s1)
 *  
 *
 *  ans:
 *			len=3 cap= 4 [ 2 3 7 ] 	// Printed in the screen
 *	
 *
 *	return:  						// No return
 * 
 *
 */
func printSlice(sliceX []int) {
	var tempSlice = make([]int, len(sliceX))
	copy(tempSlice, sliceX)

	sortSlice(tempSlice)

	fmt.Printf("len=%d cap=%d %v", len(sliceX), cap(sliceX), tempSlice)
}

/* 
 * function checkIfNumber (uInput string) bool
 *  
 * Function:	checkIfNumber 
 *  
 * Description:	Function tries to convert a string into a integer. If there is no pointer to any error 
 *				after trying, it'll return true, else it will return false
 *  
 * Parameters:	string 		- a string
 *  
 * Return:				 	- true (if succeded) / false (if not)
 *  
 * Examples of Usage: 
 *  
 *  >> uInput := "77"
 *  >> var isItNumber bool
 *  >> isItNumber = checkIfNumber(uInput)
 *  
 *  
 *
 *  ans:
 *							 		// Not printed, only returned answer
 *									// You could add fmt.Printf("%q looks like a number.", v) to check 
 *									// in screen
 *	
 *
 *	return: true
 * 
 *
 */
func checkIfNumber (uInput string) bool {
	if _, err := strconv.Atoi(uInput); err == nil {
	    
	    return true
	}

	return false
}
