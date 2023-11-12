package main

import (
	"fmt"

	"github.com/diazadriel0/go-course/exercises"
	"github.com/diazadriel0/go-course/files"
	"github.com/diazadriel0/go-course/io"
	"github.com/diazadriel0/go-course/iterations"
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

	// For
	iterations.Iterator()

	// Exercise 02
	fmt.Println(exercises.MultiplicationTable())

	// Files managment
	/* files.SaveTable()
	// for i := 3; i <= 10; i++{
		files.SumTable()
	// } */

	// Read file
	files.ReadFile()
	files.ReadLines()
}