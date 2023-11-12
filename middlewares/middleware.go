package middlewares

import "fmt"

func sum(a, b int) int {
	return a + b
}

func rest(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func MyFuncWithMiddleware() {
	fmt.Println("Excecute of MyFunc")

	midResponse := operationsMiddleware(sum)(3,8)
	fmt.Printf("The sum result is: %d\n", midResponse)
	midResponse = operationsMiddleware(rest)(5,2)
	fmt.Printf("The rest result is: %d\n", midResponse)
	midResponse = operationsMiddleware(sum)(8,8)
	fmt.Printf("The mult result is: %d\n", midResponse)
}

func operationsMiddleware(f func (int,int) int) func(int,int) int {
	return func(a, b int) int {
		fmt.Printf("The received numbers on middleware are %d and %d\n", a, b)
		return f(a,b)
	}
}