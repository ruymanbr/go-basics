/*
 * File:    animals.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	24-6-21
 * Course:  	Functions, Methods, and Interfaces in Go (University of California IRVINE)
 * Assignment:  Module 3 - Activity animals.go
 *
 * Summary of File:
 *
 *  The program allows the user to get information about a predefined set of animals. 
 *	Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. 
 *	The user can issue a request to find out one of three things about an animal: 
 *		1) the food that it eats, 
 *		2) its method of locomotion, and 
 *		3) the sound it makes when it speaks. 
 *
 */

package main

import (
	"fmt"
	"strings"
    "bufio"
    "os"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func main() {
	cow 	:= new(Animal)
	bird 	:= new(Animal)
	snake	:= new(Animal)

	initData(cow, bird, snake)

	loopUserRequest(cow, bird, snake)
}

// Eat prints in screen what's in the food field
func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

// Move prints in screen what's in the locomotion field
func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

// Speak prints in screen what's in the noise field
func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

// initData fills the 3 animals struct information
func initData(cow *Animal, bird *Animal, snake *Animal) {	

	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"

	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"

	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"

}

// loopUserRequest iterates infinitely asking user for animal + action to print what the animal does in every case
func loopUserRequest(cow *Animal, bird *Animal, snake *Animal) {

	//var userRequest string
	var userReqSli []string
	
	//var aName, aAction string
	animal := new(Animal)
	var skipProgram, restartProgram bool = false, false

	PrintInstructions()

	for {
		userReqSli = nil

		if (restartProgram == true) {
			restartProgram = false
		}
		
		fmt.Print(">")

		userReqSli, skipProgram = getInput(userReqSli)

		if (skipProgram) {
			break
		}

		userReqSli, restartProgram = validateAmount(userReqSli)

		animal, restartProgram = getAnimal(userReqSli, cow, bird, snake)

		

		if (restartProgram) {
			continue
		}

		getAction(userReqSli, *animal)
		
		
		
		
		
	}
}

// PrintInstructions prints program instructions in screen
func PrintInstructions() {
	fmt.Println("")
	fmt.Println("-----------------INSTRUCTIONS-----------------")					
	fmt.Println("")
	fmt.Println("Enter 1 animal of the list and 1 action separated by 1 space")
	fmt.Println("")
	fmt.Println("           Animal           Actions")
	fmt.Println("")
	fmt.Println("            cow              eat")
	fmt.Println("            bird             move")
	fmt.Println("            snake            speak")
	fmt.Println("")
	fmt.Println("    Example: cow move")
	fmt.Println("")
	fmt.Println("---------(E) Exits - (I) Instructions---------")
	fmt.Println("")
}

// getInput Scans strings from stdin method (keyboard) and saves it in a []string slice. Then returns it
func getInput(userReqSli []string) ([]string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	
 
	for scanner.Scan() {
		scannedTxt := scanner.Text()
		userReqSli = strings.Split(scannedTxt, " ")
		break
	}
	
	for i, userR := range userReqSli {
		userR 			= strings.ToLower(userR)
		userReqSli[i] 	= userR			
	}

	if (userReqSli[0] == "e") {
		fmt.Println("Program will exit now...")
		return userReqSli, true
	} 

	return userReqSli, false
}

// validateInput validates amount of strings from input and returns bool (true if OK)
func validateAmount(userReqSli []string) ([]string, bool) {
	if (len(userReqSli) > 2) {
		fmt.Println("You've entered more than 2 words. Start over")
		return userReqSli, true
	}

	return userReqSli, false
}

// getAnimal switches through a given []string slice and compares and returns Animal
func getAnimal(userReqSli []string, cow *Animal, bird *Animal, snake *Animal) (*Animal, bool) {

	animal := new(Animal)


	if (len(userReqSli) > 2) {
		return animal, true
	}

	switch userReqSli[0] {
		case "cow":
			return cow, false
		case "bird":
			return bird, false
		case "snake":
			return snake, false
		default:
			if (userReqSli[0] == "i") {
				PrintInstructions()
				return animal, true
			} else {
				fmt.Println("You must enter a valid animal. Try again")
				return animal, true
			}				
	}
}


// getAction switches through a given []string slice and compares and returns Animal
func getAction(userReqSli []string, animal Animal) {

	switch userReqSli[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("You must enter a valid action. Try again")
									
	}
}


// lowerStrings returns a lowecase string version of a given string
func lowerStrings(str string) string {
	return strings.ToLower(str)
}