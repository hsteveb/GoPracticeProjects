package animal

import "fmt"

type Animal struct {
	eat, move, speak string
}

func (animal Animal) Eat() {
	fmt.Println(animal.eat)
}

func (animal Animal) Move() {
	fmt.Println(animal.move)
}

func (animal Animal) Speak() {
	fmt.Println(animal.speak)
}

func (animal Animal) Behavior(behavior string) {
	if behavior == "eat" {
		animal.Eat()
	} else if behavior == "move" {
		animal.Move()
	} else {
		animal.Speak()
	}
}

func (animal *Animal) InitAnimals(eat string, move string, speak string) {
	animal.eat = eat
	animal.move = move
	animal.speak = speak
}
