package variables

import "fmt"

func ShowInt() {
	var int int
	intOf32 := int32(123456789)
	intOf64 := int64(12345678987654321)

	fmt.Println("int = ", int)
	fmt.Println("intof32 = ", intOf32)
	fmt.Println("intof64 = ", intOf64)
}