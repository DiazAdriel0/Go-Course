package exercises

import (
	"strconv"
)


func GreaterThan100(value string) (int, string) {
	int, _ := strconv.Atoi(value)
	if int > 100 {
		return int ,"Is greater than 100"
	} else {
		return int, "Is less or equal to 100"
	}
}