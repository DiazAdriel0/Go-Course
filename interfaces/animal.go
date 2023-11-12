package interfaces

type Animal interface {
	Breath()
	Eat()
	IsCarnivorous() bool
	IsAlive() bool
}