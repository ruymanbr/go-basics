/*
 * File:    race_cond.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	26-6-21
 * Course:  	Concurrency in Go (University of California IRVINE)
 * Assignment:  Module 2 - Activity race_cond.go
 *
 * Summary of File:
 *
 *  The program runs two goroutines which have a race condition when executed concurrently. 
 *	Then an explanation is given about what the race condition is and how it can occur.
 *
 */

package main

import (
	"fmt"
)

func main() {
	var anything string
	printInformation()

	fmt.Println("Press any key to start and press enter")
	fmt.Scan(&anything)
	
	goroutine1()	

}
// printInformation prints Activity condition to be replicated and outcome expentancy
func printInformation() {
	fmt.Println("")
	fmt.Println("Welcome to race_cond.go")
	fmt.Println("")
	fmt.Println("We'll run goroutine1 and a second go func() goroutine will be launched from there.")
	fmt.Println("")
	fmt.Println("Both will print 'i' when they are being run by goruntime scheduler")
	fmt.Println("")
	fmt.Println("We may expect discrepancies depending on goroutine scheduler switching them")
	fmt.Println("")
	fmt.Println("A race condition will eventually develop and see some unordered values of 'i' printed")
}

// goroutine1 executes a for loop with a go routine inside it making modifications to the looping index (race condition)
func goroutine1() {
	i:= 0	

	for i < 100 {
		go func() {
			i++
			fmt.Println("Go routine 2 here!: g1 = ", i)
			fmt.Println("g1", i)
		}()
		fmt.Println("g1 = ", i)			
		i++
	} 
}