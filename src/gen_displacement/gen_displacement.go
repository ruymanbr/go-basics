/*
 * File:    gen_displacement.go
 *
 * Author:  	Ruymán Borges R. (ruyman21@gmail.com)
 * Date:    	23-6-21
 * Course:  	Functions, Methods, and Interfaces in Go (University of California IRVINE)
 * Assignment:  Module 2 - Activity gen_displacement.go
 *
 * Summary of File:
 *
 *  The program prompts the user to enter values for acceleration, initial velocity,
 *	and initial displacement. Then the program should prompts the user to enter a value for
 *	time and computes the displacement after the entered time.
 *
 */
package main

import (
	"fmt"
	"math"
)

func main() {

	var accel, v_ini, s_ini, time float64

	accel, v_ini, s_ini = getVariables(accel, v_ini, s_ini)

	fmt.Println("Enter time to calculate total displacement:")

	fmt.Scan(&time)

	fn := GenDisplaceFn(accel, v_ini, s_ini)

	fmt.Println("Total displacement is...")
	fmt.Println(fn(time))
}

// GenDisplaceFn
func GenDisplaceFn(accel, v_ini, s_ini float64) func(float64) float64 {
	
	fn := func(t float64) float64 {
		displacement := (0.5*accel*math.Pow(t, 2)) + (v_ini*t) + s_ini
		return displacement
	}

	return fn
}

// getVariables prompts user for acceleration, initial velocity, and initial displacement (accel, v_ini, s_ini)
func getVariables(uAccel, uV_ini, uS_ini float64) (float64, float64, float64) {

	fmt.Println("Enter initial acceleration:")
	fmt.Scan(&uAccel)

	fmt.Println("Enter initial velocity:")
	fmt.Scan(&uV_ini)

	fmt.Println("Enter initial displacement:")
	fmt.Scan(&uS_ini)

	return uAccel, uV_ini, uS_ini
}
