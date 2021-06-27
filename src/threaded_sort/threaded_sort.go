/*
 * File:    threaded_sort.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	27-6-21
 * Course:  	Concurrency in Go (University of California IRVINE)
 * Assignment:  Module 3 - Activity threaded_sort.go
 *
 * Summary of File:
 *
 *  The program sorts an array of integers. It does a partition of the array into 4 parts, each of which 
 *	is sorted by a different goroutine. Each partition will be of approximately equal size. Then 
 *	the main goroutine merges the 4 sorted subarrays into one large sorted array. 
 *	
 *	The program prompts the user to input a series of integers. Each goroutine which sorts ¼ of the array 
 *	prints the subarray that it will sort. When sorting is complete, the main goroutine should print 
 *	the entire sorted list.
 *
 */

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	var userReqSliInt, orderedSliInt []int
	var uSli [][]int
	var oSli1, oSli2, oSli3, oSli4 []int
	var skipProgram, restartProgram bool = false, false
	m_master 	:= make(chan []int)
	m_g1 		:= make(chan []int)
	m_g2 		:= make(chan []int)
	m_g3 		:= make(chan []int)
	m_g4 		:= make(chan []int)
	var chunkSize, resVal int
	var res bool

	PrintInstructions()

	for {
		userReqSliInt = nil
		uSli, oSli1, oSli2, oSli3, oSli4	= nil, nil, nil, nil, nil

		if (restartProgram == true) {
			restartProgram 	= false
			skipProgram 	= false
		}
		fmt.Println("Enter a set of integers separated by 1 space. Then press enter to begin")
		fmt.Println("")
		fmt.Print(">")

		userReqSliInt, skipProgram, restartProgram = getInput()

		if (skipProgram) {
			break;
		}

		if (restartProgram) {
			PrintInstructions()
			continue;
		}

		chunkSize, res, resVal = calculateChunkSize(userReqSliInt, 4)

		if (res) {
			uSli = divideSliceFourWithRes(userReqSliInt, chunkSize, resVal)
		} else {
			uSli = divideSliceFour(userReqSliInt, chunkSize)
		}

		go BubbleSort(uSli[0], m_g1, true)
		go BubbleSort(uSli[1], m_g2, true)
		go BubbleSort(uSli[2], m_g3, true)
		go BubbleSort(uSli[3], m_g4, true)
		

		oSli1 = <- m_g1
		oSli2 = <- m_g2
		oSli3 = <- m_g3
		oSli4 = <- m_g4

		
		go mergeSliInt(oSli1, oSli2, oSli3, oSli4, m_master)
		
		orderedSliInt = <- m_master

		go BubbleSort(orderedSliInt, m_master, false)

		orderedSliInt = <- m_master

		fmt.Println("Sorted final sequence is: ")
		printMe(orderedSliInt)

		restartProgram = checkforRepeat()

		if (restartProgram) {
			PrintInstructions()
			continue;
		} else {
			fmt.Println("Exiting program now...")
			break;
		}
	}

}

// getInput Scans strings from stdin method (keyboard) and saves it in a []string slice. Then returns it
func getInput() ([]int, bool, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var userReqSli []string
	var userReqSliInt []int
	var userInt int
	var err error
	
 
	for scanner.Scan() {
		scannedTxt := scanner.Text()
		userReqSli = strings.Split(scannedTxt, " ")
		break
	}

	if (userReqSli[0] == "e") {
		fmt.Println("Program will exit now...")
		return userReqSliInt, true, false
	}

	if (userReqSli[0] == "i") {
		return userReqSliInt, false, true
	} 
	
	for _, userR := range userReqSli {
		userInt, err	= strconv.Atoi(userR)
		if err == nil {
			userReqSliInt 	= append(userReqSliInt, userInt)	
		} else {
			fmt.Println("Error happened: ", err)
			return userReqSliInt, true, false
		}
		
	}

	return userReqSliInt, false, false
}

// calculateChunkSize  calculates the size of equal(ish) totalNumChunks parts
func calculateChunkSize(userReqSliInt []int, numChunks int) (int, bool, int) {
	var lenArr, resVal, chunkSize int

	lenArr		= len(userReqSliInt)
	chunkSize 	= int(lenArr / numChunks)
	resVal 		= lenArr % numChunks

	if (resVal > 0) {
		return chunkSize, true, resVal
	} else {
		return chunkSize, false, 0	
	}

	
}

// divideSliceFour divide user Slice []int type into equal chunks, exception last chunk, for uneven rest of len(arr)/ num_chunk
func divideSliceFourWithRes(userReqSliInt []int, chunkSize int, resVal int) [][]int {
	var chunks [][]int

	for i := 0; i < len(userReqSliInt); i += (chunkSize + 1) {
		end := i + chunkSize
		if (resVal > 0) {			
			end = i + chunkSize + 1			
			resVal = resVal - 1
		} 

		// necessary check to avoid slicing beyond
		// slice capacity

		if end > len(userReqSliInt) {
			end = len(userReqSliInt)
		}
		
		chunks = append(chunks, userReqSliInt[i:end])
	}

	return chunks
}

// divideSliceFour divide user Slice []int type into equal chunks
func divideSliceFour(userReqSliInt []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(userReqSliInt); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(userReqSliInt) {
			end = len(userReqSliInt)
		}

		chunks = append(chunks, userReqSliInt[i:end])
	}

	return chunks
}

// BubbleSort executes Bubble Sort algorithm in a []int slice
func BubbleSort(sli []int, c chan []int, printBefore bool) {
	var sorted bool = false
	var i, sorted_i int

	if (printBefore) {
		fmt.Println("Array to be sorted: ", sli)
	}	
	
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
	c <- sli

	
}

// Swap switches positions of 2 digits inside a given []int array
func Swap(sli []int, i int) {
	var num1, num2 int

	num1 		= sli[i]
	num2 		= sli[i+1]

	sli[i] 		= num2
	sli[i+1]	= num1
}

// mergeSliInt merges 
func mergeSliInt(sli1 []int, sli2 []int, sli3 []int, sli4 []int, c chan []int) {
	var orderedSliInt []int

	for _, val := range sli1 {
		orderedSliInt = append(orderedSliInt, val)	
	}
	for _, val := range sli2 {
		orderedSliInt = append(orderedSliInt, val)	
	}
	for _, val := range sli3 {
		orderedSliInt = append(orderedSliInt, val)	
	}
	for _, val := range sli4 {
		orderedSliInt = append(orderedSliInt, val)	
	}


	c <- orderedSliInt
}

// PrintInstructions prints program instructions for user
func PrintInstructions() {
	fmt.Println("Welcome to threaded_sort.go (by Ruymán B.R.")
	fmt.Println("The program will sort a []int slice or array using 4 goroutines")
	fmt.Println("")
	fmt.Println("Each goroutine will show which part is sorting before starting")
	fmt.Println("")
	fmt.Println("Quick Commands --> (E) Exits -- (I) Instructions")
	fmt.Println("")
	
}

// printMe prints anything given to it
func printMe (val interface{}) {
	fmt.Println(val)
}

// checkforRepeat prompts the user for a restart in program
func checkforRepeat() bool {
	var ans string
	fmt.Println("")
	fmt.Println("Start Over? (y) for restart, any other letter Exits)")
	fmt.Scan(&ans)

	if strings.ToLower(ans) == "y" {
		return true
	} else {
		return false
	}
}