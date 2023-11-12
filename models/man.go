package models

type Man struct {
	name      string
	age       int8
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	alive     bool
}

func (man *Man) Breathe()       { man.breathing = true }
func (man *Man) Think()         { man.thinking = true }
func (man *Man) Eat()           { man.eating = true }
func (man *Man) Gender() string { return "Man" }
func (man *Man) IsAlive() bool  { return true }
