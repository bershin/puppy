package puppy

import "github.com/bershin/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof!!"
}

func BigBark() string {
	return dog.Bark(Bark())
}

func BigBarks() string {
	return dog.Bark(Barks())
}
