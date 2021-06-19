/* 
 * File:    makejson.go 
 * 
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com) 
 * Date:    	19-6-21
 * Course:  	Starting with Go (University of California IRVINE)
 * Assignment:  Module 4 - Activity makejson.go
 * 
 * Summary of File: 
 * 
 *  The program prompts the user to first enter a name, and then enter an address. 
 *	Your program should create a map and add the name and address to the map using the 
 *	keys “name” and “address”, respectively. Your program should use Marshal() to create 
 *	a JSON object from the map, and then your program should print the JSON object.
 * 
 */ 

package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
	"ioutil"
	"os"
)

func main() {

	var idMap = make(map[string]string)
	var nameUser, addressUser string
	var mapToJSON = make([]byte)

	fmt.Println("I'm gonna ask you for name and address sequencially. Then I'll create a JSON using that information")

	askNameAdress(nameUser, addressUser)

	addAttributeAndValueToMap(nameUser, addressUser, idMap)

	mapToJSON, _ = createJSONFromMap(idMap)

	printJSON(mapToJSON)

}

// askNameAdress prompts user for 2 inputs and store them in variables
// As in "attribute: value" pair of strings
func askNameAdress(nameUser string, addressUser string) {

	fmt.Println("Please enter your name")
	fmt.Scan(&nameUser)

	fmt.Println("Please enter your address")
	fmt.Scan(&addressUser)

}

// addAttributeAndValueToMap adds an "attr: value" pair for "name" and "address" into a map[string]string
func addAttributeAndValueToMap(uName string, uAddress string, idMap map[string]string)  {
	idMap["name"] = uName

	idMap["address"] = uAddress
}

// createJSONFromMap creates a JSON object from a map[string]string and returns it
func createJSONFromMap(idMap map[string]string) []byte {
	return barr, _ := json.Marshal(idMap)
}

// printJSON prints a JSON byte array in screen
func printJSON(barr []byte) {
	fmt.Println("JSON []byte array is: ")
	fmt.Printf(barr)
}