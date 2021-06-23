/*
 * File:    read.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	19-6-21
 * Course:  	Starting with Go (University of California IRVINE)
 * Assignment:  Module 4 - Activity read.go
 *
 * Summary of File:
 *
 *  The program reads information from a file and represents it in a slice of structs. 
 *  It assumes that there is a text file which contains a series of names. Each line of the text 
 *	file has a first name and a last name, in that order, separated by a single space on the line. 
 *
 *	the program will define a name struct which has two fields, fname for the first name, 
 *	and lname for the last name. Each field will be a string of size 20 (characters).
 *
 *	Also, it will prompt the user for the name of the text file, and it will successively read 
 *	each line of the text file and create a struct which contains the first and last names found
 *	in the file. Each struct created will be added to a slice, and after all lines have been read
 *	from the file, your program will have a slice containing one struct for each line in the file. 
 *	After reading all lines from the file, your program should iterate through your slice of structs 
 *	and print the first and last names found in each struct.
 *
 */
package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
)

type FullName struct {
	fname string
	lname string
}

func main() {
	

	var names []FullName
	var lines []string
	//var barr = make([]byte, 1)
	var textFile string
	var f *os.File


	fmt.Println("This program opens a file and will read name and last name for every line in a text file. ")
	fmt.Println("Once it's done It'll print in screen the whole set of 'Name LastName' ")
	
	fmt.Println("Please enter text file name including extension (.txt) now:")

	fmt.Scan(&textFile)

	f = ReadFile(f, textFile)

	lines = ExtractLines(f)

	names = PopulateSliceStruct(f, names, lines)

	PrintSliceStruct(names)
}

// readFile reads a file and returns it as os.File type
func ReadFile(f *os.File, textFile string) *os.File {
	f, _ = os.Open(textFile)

	return f
}

// ExtractLines saves each separate line from a given file into a []string slice and returns it
func ExtractLines(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	//fmt.Println("txtlines are: ", txtlines)
	/*
	for _, names := range txtlines {
		fmt.Println(names)
	}
	*/

	return txtlines
}

// BreakNameLine separate strings using blank space as string separator
func BreakNameLine(line string) []string {
	stringFrag := strings.Split(line, " ")

	return stringFrag
}

// populateSliceStruct fill []FullName fields for every fname, lname in os.File type file
// fname is separated by 1 space byte from lname
func PopulateSliceStruct(f *os.File, names []FullName, lines []string) []FullName {
	
	var nameFrags []string
	name := new(FullName)	
	
	for _, line := range lines {
		nameFrags = BreakNameLine(line)

		name.fname = nameFrags[0]
		name.lname = nameFrags[1]

		names = append(names, *name)
	}	
	
	
	
	return names

}

// printSliceStruct prints fname, lname pairs inside []FullName struct type in screen
func PrintSliceStruct(names []FullName) {
	fmt.Println("Content of []FullName struct is...")
	fmt.Println("-------------- NAMES --------------")
	for _,name := range names {		
		fmt.Printf(name.fname)
		fmt.Println(" ", name.lname)
	}
	fmt.Println("-----------------------------------")
}

