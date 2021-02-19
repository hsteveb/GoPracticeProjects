package animal

import "fmt"

type Snake struct {
	name string
}

func (snake Snake) Init(name string) Animal {
	snake.name = name
	return snake
}

func (snake Snake) GetName() string {
	return snake.name
}

func (snake Snake) Eat() {
	fmt.Println("Mice")
}

func (snake Snake) Speak() {
	fmt.Println("Hsss")
}

func (snake Snake) Move() {
	fmt.Println("Slither")
}
