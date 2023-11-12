package arraysslices

import "fmt"

var slice []int = []int{2,5,3,11}
var arrayOfInts [10]int = [10]int{50,134,2,7,77,99,66}

func ShowSlice() {
	fmt.Println(slice)

	portion1 := arrayOfInts[3:] // slice from 3 to 9
	portion2 := arrayOfInts[:5] // slice from 0 to 4
	portion3 := arrayOfInts[6:7] // slice from 6 to 6

	fmt.Println(portion1)
	fmt.Println(portion2)
	fmt.Println(portion3)
}

func Capacity() {
	elements := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d\n", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))

}