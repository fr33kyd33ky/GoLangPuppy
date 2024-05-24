package golangpuppy

import (
	golangdog "github.com/fr33kyd33ky/GoLangDog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return golangdog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return golangdog.WhenGrownUp(Barks())
}
