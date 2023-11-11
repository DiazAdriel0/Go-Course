package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() {
	var num int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter a integer:")
	if scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		for err != nil {
			fmt.Println(err.Error())
			fmt.Println("Enter a integer:")
			if scanner.Scan(){
				num, err = strconv.Atoi(scanner.Text())
			}
		}

		for i := 1; i <= 10; i++ {
			fmt.Printf("%d * %d = %d \n", num, i, num * i)
		}
	}
}