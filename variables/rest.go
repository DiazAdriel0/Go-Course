package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func RestVariables() {
	Name = "Adri"
	Status = true
	Salary = 1000.85
	Date = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToText(number int) (bool,string) {
	text := strconv.Itoa(number)
	return true, text
}