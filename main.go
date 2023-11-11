package main

import (
	"fmt"

	"github.com/diazadriel0/go-course/exercises"
	"github.com/diazadriel0/go-course/io"
	"github.com/diazadriel0/go-course/variables"
)

func main() {
	// Define variables
	variables.ShowInt()
	variables.RestVariables()

	// Asign variables
	status, text := variables.ConvertToText(int(variables.Salary))
	fmt.Println(status)
	fmt.Println(text)

	// Exercise 01
	intValue, isGreater := exercises.GreaterThan100("999")
	fmt.Println("The int value is: ", intValue)
	fmt.Println(isGreater)

	// Error
	fmt.Println(exercises.GreaterThan100("fffffff"))

	// Input/output
	io.InputKeyboard()
}