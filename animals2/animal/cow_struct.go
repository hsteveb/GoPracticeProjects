package animal

import "fmt"

type Cow struct {
	name string
}

func (cow Cow) Init(name string) Animal {
	cow.name = name
	return cow
}

func (cow Cow) GetName() string {
	return cow.name
}

func (cow Cow) Eat() {
	fmt.Println("Grass")
}

func (cow Cow) Speak() {
	fmt.Println("Moo")
}

func (cow Cow) Move() {
	fmt.Println("Walk")
}
