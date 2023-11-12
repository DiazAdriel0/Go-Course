package main

import (
	"fmt"

	arraysslices "github.com/diazadriel0/go-course/arrays-slices"
	"github.com/diazadriel0/go-course/exercises"
	"github.com/diazadriel0/go-course/files"
	"github.com/diazadriel0/go-course/functions"
	"github.com/diazadriel0/go-course/io"
	"github.com/diazadriel0/go-course/iterations"
	"github.com/diazadriel0/go-course/maps"
	"github.com/diazadriel0/go-course/models"
	"github.com/diazadriel0/go-course/users"
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

	// Anonymous func
	functions.Calc()

	// Closure func
	functions.ClosureCall(2)
	functions.ClosureCall(3)

	// Recursion
	functions.Exponential(2)
	functions.Exponential(3)

	// Arrays
	arraysslices.ShowArrays()

	// Slices
	arraysslices.ShowSlice()
	arraysslices.Capacity()

	// Maps
	maps.ShowMaps()

	// Structures
	users.RegisterUser()

	// Interfaces
	adri := new(models.Man)
	exercises.HumansBreathing(adri)
	juli := new(models.Woman)
	exercises.HumansBreathing(juli)
}