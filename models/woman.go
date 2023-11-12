package models

type Woman struct {
	Man
	pregnant bool
}

func (woman *Woman) Breathe()       { woman.breathing = true }
func (woman *Woman) Think()         { woman.thinking = true }
func (woman *Woman) Eat()           { woman.eating = true }
func (woman *Woman) Gender() string { return "Woman" }
func (woman *Woman) IsAlive() bool  { return true }
