package functions

import "fmt"

func Calc() {
	var num1 int = 15
	var num2 int = 10

	calc := func(num3 int,num4 int) int {
		return num1 + num2 + num3 + num4
	}
	fmt.Println(calc(1,2))

	calc = func(num3 int,num4 int) int {
		return num2 * num3 * num4
	}
	fmt.Println(calc(5,10))
}