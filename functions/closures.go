package functions

import "fmt"

func table(value int) func() int {
	num := value
	sequence := 0
	return func() int {
		sequence++
		return num * sequence
	}
}

func ClosureCall(tableOf int) {
	myTable := table(tableOf)
	for i := 1; i <= 10; i++ {
		fmt.Println(myTable())
	}
}