package arraysslices

import "fmt"

var array [10]int = [10]int{15, 12, 28, 1}
var matrix [20][30]int

func ShowArrays() {
	array[8] = 51
	array[2] = 77

	stringArr := [10]string{"Adri", "Diaz"}

	fmt.Println(array)
	fmt.Println(stringArr)

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	matrix[3][25] = 999
	fmt.Println(matrix)
}