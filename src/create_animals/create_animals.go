/*
 * File:    create_animals.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	25-6-21
 * Course:  	Functions, Methods, and Interfaces in Go (University of California IRVINE)
 * Assignment:  Module 4 - Activity create_animals.go
 *
 * Summary of File:
 *
 *  The program allows the user to create a set of animals and to get information about those animals. 
 *	Each animal has a name and can be either a cow, bird, or snake. 
 *	With each command, the user can either create a new animal of one of the three types, or the user 
 *	can request information about an animal that he/she has already created. 
 *	Each animal has a unique name, defined by the user. Note that the user can define animals 
 *	of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
 *
 */

package main

import (
	"fmt"
	"strings"
    "bufio"
    "os"
    "time"
)

type Animal interface {	
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
	food string
	locomotion string
	noise string
}

type Bird struct {
	name string
	food string
	locomotion string
	noise string
}

type Snake struct {
	name string
	food string
	locomotion string
	noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}


func main() {

	
	var uCows []Cow
	var uBirds []Bird
	var uSnakes []Snake

	loopUserRequest(uCows, uBirds, uSnakes)
}


// initCow fills the Cow struct information
func initCow(name string) Cow {	
	var cow = new(Cow)

	cow.food 			= "grass"
	cow.locomotion 		= "walk"
	cow.noise 			= "moo"
	cow.name 			= name

	return *cow

}

// initBird fills the Bird struct information
func initBird(name string) Bird {	
	var bird = new(Bird)

	bird.food 			= "worms"
	bird.locomotion 	= "fly"
	bird.noise 			= "peep"
	bird.name 			= name

	return *bird

}

// initSnake fills the Snake struct information
func initSnake(name string) Snake {	
	var snake = new(Snake)

	snake.food 			= "mice"
	snake.locomotion 	= "slither"
	snake.noise 		= "hsss"
	snake.name 			= name

	return *snake

}

// loopUserRequest iterates infinitely asking user for a command (either "newanimal" or "query") and initiates processes
func loopUserRequest(uCows []Cow, uBirds []Bird, uSnakes []Snake) {

	var userReqSli []string
	//var animal Animal
	var command, aName, action, addType string
	
	var skipProgram, restartProgram, printAnimals bool = false, false, false

	
	PrintInstructions()

	for{

		userReqSli = nil

		if (restartProgram == true) {
			restartProgram 	= false
			printAnimals 	= false
		}
		
		fmt.Print(">")

		userReqSli, skipProgram, printAnimals = getInput(userReqSli)

		if (skipProgram) {
			break
		}

		if (printAnimals) {

			printUserAnimals(uCows, uBirds, uSnakes)
			continue
		} 

		restartProgram = validateAmount(userReqSli)

		if (restartProgram) {
			waitTimeInstructions(2)
			continue
		}

		command, restartProgram = evaluateInput(userReqSli)
		


		if (restartProgram) {			

			continue
		} else {
			
			aName = userReqSli[1]

			
			switch command {
				case "newanimal":
					addType = userReqSli[2]
					switch addType {
						case "cow":
							uCows = addCow(aName, uCows)
						case "bird":
							uBirds = addBird(aName, uBirds)
						case "snake":
							uSnakes = addSnake(aName, uSnakes)
					}

				case "query":
					action 	= userReqSli[2]

					seekAnimData(aName, action, uCows, uBirds, uSnakes)

				default:

					fmt.Println("Something went wrong. Try again")
					PrintInstructions()
			}
		}

	}
}

// PrintInstructions prints program instructions in screen
func PrintInstructions() {
	fmt.Println("-----------------------INSTRUCTIONS---------------------")					
	fmt.Println("")
	fmt.Println("   Enter 'newanimal' or 'query' followed by commands:")
	fmt.Println("")
	fmt.Println("         (Always 3 words separated by 1 space)")
	fmt.Println("")
	fmt.Println("           newanimal + name + type_of_animal")
	fmt.Println("")
	fmt.Println("                        OR")
	fmt.Println("")
	fmt.Println("               query + name + Information")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("           type_of_animal     Information")
	fmt.Println("")
	fmt.Println("              cow                eat")
	fmt.Println("              bird               move")
	fmt.Println("              snake              speak")
	fmt.Println("")
	fmt.Println("('name' is any string if adding or existing if querying)")
	fmt.Println("")
	fmt.Println("----(L) List Animals -- (E) Exits - (I) Instructions----")
}

func waitTimeInstructions(seconds int) {
	fmt.Println("Starting Over")
	DurationOfTime := time.Duration(seconds) * time.Second
	f := func() {
		PrintInstructions()
	}

	Timer1 := time.AfterFunc(DurationOfTime, f)
	defer Timer1.Stop()
}


// getInput Scans strings from stdin method (keyboard) and saves it in a []string slice. Then returns it
func getInput(userReqSli []string) ([]string, bool, bool) {
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
		return userReqSli, true, false
	}

	if (userReqSli[0] == "l") {
		fmt.Println("The current animal list is:")
		return userReqSli, false, true
	} 

	return userReqSli, false, false
}

// validateInput validates amount of strings from input and returns bool (true if OK)
func validateAmount(userReqSli []string) bool {
	if (len(userReqSli) > 3) {
		fmt.Println("You've entered more than 3 words")
		return true
	}

	if (len(userReqSli) < 3 && len(userReqSli) > 1) {
		fmt.Println("You must enter 3 words")
		return true
	}

	return false
}

// evaluateInput switches through a given []string slice and gets first one to look for user command. Returns bool that is true for restarting the program
func evaluateInput(userReqSli []string) (string, bool) {

	var command string = userReqSli[0]
	var restartProgram bool

	if (len(userReqSli) > 3) {
		return command, true
	}

	switch command {
		case "newanimal":
			restartProgram = false
		case "query":
			restartProgram = false
		default:
			if command == "i" {
				PrintInstructions()
				restartProgram = true
			} else {
				fmt.Println("You must enter a valid command (newanimal/query). Try again")
				restartProgram = true
			}				
	}
	return command, restartProgram
}

// addCow adds a new Cow type to a given []Cow slice
func addCow(animName string, cows []Cow) []Cow {
	var cow Cow

	cow = initCow(animName)

	cows = append(cows, cow)

	println("Cow added!")

	return cows
}


// addBird adds a new Bird type to a given []Bird slice
func addBird(animName string, birds []Bird) []Bird {

	var bird Bird

	bird = initBird(animName)

	birds = append(birds, bird)

	println("Bird added!")

	return birds
}


// addSnake adds a new Snake type to a given []Snake slice
func addSnake(animName string, snakes []Snake) []Snake {

	var snake Snake

	snake = initSnake(animName)

	snakes = append(snakes, snake)

	println("Snake added!")

	return snakes
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

// printUserAnimals prints all animals user has added
func printUserAnimals(uCows []Cow, uBirds []Bird, uSnakes []Snake) {
	fmt.Println("")
	fmt.Println("------- COWS: -------")
	fmt.Println("")
	for _, uCow := range uCows {
		
		fmt.Println(uCow)
		
	}

	fmt.Println("")
	fmt.Println("------- BIRDS: ------")
	fmt.Println("")
	for _, uBird := range uBirds {
		
		fmt.Println(uBird)
		
	}

	fmt.Println("")
	fmt.Println("------- SNAKES: -----")
	fmt.Println("")
	for _, uSnake := range uSnakes {
		
		fmt.Println(uSnake)
		
	}
	fmt.Println("")
	fmt.Println("-----------------------------")
}

// seekAnimData iterates through 3 given slices and prints out seeked animal name
func seekAnimData(aName string, action string, uCows []Cow, uBirds []Bird, uSnakes []Snake) {


	for _, cow := range uCows {
		if cow.name == aName {
			switch action {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
			}	
		}
	}

	for _, bird := range uBirds {
		if bird.name == aName {
			switch action {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
			}
		}
	}

	for _, snake := range uSnakes {
		if snake.name == aName {
			switch action {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
			}
		}
	}
}