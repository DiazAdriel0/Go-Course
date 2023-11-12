package exercises

import (
	"fmt"

	"github.com/diazadriel0/go-course/interfaces"
)

func HumansBreathing(human interfaces.Human) {
	human.Breathe()
	fmt.Printf("I'm a %s, and I'm breathing\n", human.Gender())
}