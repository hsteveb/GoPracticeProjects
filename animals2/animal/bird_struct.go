package animal

import "fmt"

type Bird struct {
	name string
}

func (bird Bird) Init(name string) Animal {
	bird.name = name
	return bird
}

func (bird Bird) GetName() string {
	return bird.name
}

func (bird Bird) Eat() {
	fmt.Println("Worms")
}

func (bird Bird) Speak() {
	fmt.Println("Peep")
}

func (bird Bird) Move() {
	fmt.Println("Fly")
}
