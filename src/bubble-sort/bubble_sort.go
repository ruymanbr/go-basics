/*
 * File:    bubble_sort.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	22-6-21
 * Course:  	Functions, Methods, and Interfaces in Go (University of California IRVINE)
 * Assignment:  Module 1 - Activity bubble_sort.go
 *
 * Summary of File:
 *
 *  The program prompts the user to type in a sequence of up to 10 integers. 
 *	Then it prints the integers out on one line, in sorted order, from least to
 *	greatest using the bubble sort algorithm (based on https://www.geeksforgeeks.org/bubble-sort/).
 *  
 *	The program includes a function called BubbleSort() which takes a slice of integers as an argument
 *	and returns nothing. The BubbleSort() function modifies the slice so that the elements are in sorted
 *	order.
 *
 *	A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position 
 *	of two adjacent elements in the slice. The Swap() function performs this operation. It takes 
 *	two arguments, a slice of integers and an index value i which indicates a position in the slice. 
 *	The Swap() function returns nothing, but it swaps the contents of the slice in position i with 
 *	the contents in position i+1.
 *
 */
package main

import (
	"fmt"
)

func main() {
	var nArr = make([]int, 10)
	
	
	fmt.Println("Welcome to the Bubble Sort program.")

	nArr = LoopInputNum(nArr)

	

	fmt.Println("Now the sorting part. Let's begin...")	

	BubbleSort(nArr)
	
	

	fmt.Println("Ordered array is:")
	PrintArr(nArr)
}

// LoopInputNum iterates through an slice adding int value inputs into it in input order
func LoopInputNum(nArr []int) ([]int) {
	var userInput, i int

	for i = 0; i<10; i++ {
		fmt.Println("Enter number ", i+1)
		fmt.Scan(&userInput)
		nArr[i] = userInput
	}

	return nArr
}

// PrintArr(nArr) prints a []int array in screen
func PrintArr(nArr []int) {	
	fmt.Println(nArr)
}

// BubbleSort executes Bubble Sort algorithm in a []int slice
func BubbleSort(sli []int) {
	var sorted bool = false
	var i, sorted_i int
	
	for sorted == false {
		sorted_i = 0
		for i = 0; i < (len(sli) - 1); i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				sorted_i += 1
			}
		}

		if sorted_i == 0 {
			sorted = true
		}
	}
		
}

// Swap switches positions of 2 digits inside a given []int array
func Swap(sli []int, i int) {
	var num1, num2 int

	num1 		= sli[i]
	num2 		= sli[i+1]

	sli[i] 		= num2
	sli[i+1]	= num1
}