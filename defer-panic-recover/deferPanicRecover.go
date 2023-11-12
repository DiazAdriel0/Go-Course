package deferpanicrecover

import (
	"fmt"
	"log"
)

func Defer() {
	fmt.Println("First message")
	defer fmt.Println("Defer message")
	fmt.Println("Final message")
}

func PanicExample() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("An error occurs and generate a Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Value 1 found")
	}
}