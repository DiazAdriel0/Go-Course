package exercises

import (
	"strconv"
)


func GreaterThan100(value string) (int, string) {
	int, error := strconv.Atoi(value)

	if error != nil {
		return 0, "An error occurred: " + error.Error()
	}

	if int > 100 {
		return int ,"Is greater than 100"
	} else {
		return int, "Is less or equal to 100"
	}
}