package goRoutines

import (
	"fmt"
	"strings"
	"time"
)

func SlowPrint(str string, channel1 chan bool) {
	split := strings.Split(str, "")
	for _, char := range split {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(char)
	}

	channel1 <- true
}