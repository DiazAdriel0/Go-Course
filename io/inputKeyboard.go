package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var instruction string
var err error

func InputKeyboard() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the instruction: ")
	if scanner.Scan() {
		instruction = scanner.Text()
	}

	fmt.Println("Enter first number: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The input is invalid: " + err.Error())
		}
	}

	fmt.Println("Enter second number: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The inputy is invalid: " + err.Error())
		}
	}


	fmt.Println(instruction, num1 * num2)
}