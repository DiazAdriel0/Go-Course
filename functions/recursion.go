package functions

import "fmt"

func Exponential(value int) int {
	if value > 1000000000 {
		return value
	}

	fmt.Println(value)
	return Exponential(value * 2)
}

