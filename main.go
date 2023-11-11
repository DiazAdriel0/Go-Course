package main

import (
	"fmt"

	"github.com/diazadriel0/go-course/exercises"
	"github.com/diazadriel0/go-course/variables"
)

func main() {
	variables.ShowInt()
	variables.RestVariables()

	status, text := variables.ConvertToText(int(variables.Salary))
	fmt.Println(status)
	fmt.Println(text)

	intValue, isGreater := exercises.GreaterThan100("999")
	fmt.Println("The int value is: ", intValue)
	fmt.Println(isGreater)
}