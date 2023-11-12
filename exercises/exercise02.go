package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() string {
	var num int
	var err error
	var table string
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
			table += fmt.Sprintf("%d * %d = %d \n", num, i, num * i)
		}
	}

	return table
}