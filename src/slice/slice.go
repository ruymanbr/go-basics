/* 
	Assignment 4 - Module 3 - Activity slice.go
	The program slice.go prompts the user to enter integers and stores the integers in a sorted slice. 
	The program should be written as a loop. Before entering the loop, the program should create 
	an empty integer slice of size (length) 3. During each pass through the loop, 
	the program prompts the user to enter an integer to be added to the slice. 
	The program adds the integer to the slice, sorts the slice, and prints the contents of the 
	slice in sorted order. The slice must grow in size to accommodate any number of integers 
	which the user decides to enter. The program should only quit (exiting the loop) 
	when the user enters the character ‘X’ instead of an integer.

	Author: Ruymán Borges
	Date: 18-6-21
*/

	/* 
	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
	    return a[i] < a[j]
	})
	for _, v := range a {
	    fmt.Println(v)
	}

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
	//var intInput int

	
	fmt.Println("\nYou will enter 1 integer at a time and the program will show you the sorted slice at every step.\n")

	exitProgram = insertFirstThreeIntegers(s1, userInput)	
		


	for exitProgram == false {
		fmt.Println("\nYou can keep adding integers to the slice. Enter (X) to EXIT: \n")

		exitProgram, s1 = increasingSlice(s1, userInput)
	}

	
}

func insertFirstThreeIntegers(sliceX []int, uInput string) bool {

	exitProgram := false

	for i, _ := range sliceX {



		fmt.Println("\nEnter 1 integer now and press ENTER: enter (X) to EXIT\n")
		fmt.Scan(&uInput)

		if !(checkIfNumber(uInput)) {
			if (evalInput(uInput) == "x") {
				fmt.Println("\nYou have successfully EXITED the program. Have a nice day!\n")
				exitProgram = true
				
				break
			}

			for {
				fmt.Println("\nIt MUST be an integer. Please try again or enter (X) to EXIT\n")
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
		//sortSlice(sliceX)
		printSlice(sliceX)

	}

	return exitProgram
}


func increasingSlice(sliceX []int, uInput string) (bool, []int) {

	fmt.Println("\nEnter 1 integer now and press ENTER: enter (X) to EXIT\n")
	fmt.Scan(&uInput)

	//intInput, _ := strconv.Atoi(uInput)
	if !(checkIfNumber(uInput)) {
		if evalInput(uInput) == "x" {
			fmt.Println("\nYou have successfully EXITED the program. Have a nice day!\n")
			return true, sliceX;
		}

		for {
			fmt.Println("\nIt MUST be an integer. Please try again or enter (X) to EXIT\n")
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

	//sortSlice(sliceX)
	printSlice(sliceX)

	return false, sliceX;
}

func evalInput(uInput string) string {
	
	return strings.ToLower(strings.TrimSpace(uInput))
}

func sortSlice(tempSliceX []int) {
	
	sort.Slice(tempSliceX, func(i, j int) bool {
	    return tempSliceX[i] < tempSliceX[j]	    
	})
}

func printSlice(sliceX []int) {
	var tempSlice = make([]int, len(sliceX))
	copy(tempSlice, sliceX)

	sortSlice(tempSlice)

	fmt.Printf("len=%d cap=%d %v\n", len(sliceX), cap(sliceX), tempSlice)
}

func checkIfNumber (uInput string) bool {
	if _, err := strconv.Atoi(uInput); err == nil {
	    //fmt.Printf("%q looks like a number.\n", v)
	    return true
	}

	return false
}