package goRoutines

import (
	"fmt"
	"strings"
	"time"
)

func SlowPrint(str string) {
	split := strings.Split(str, "")
	for _, char := range split {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(char)
	}
}