package animal

import "testing"

//import "fmt"

func TestAnimalsStruct(t *testing.T) {

	var cow Animal
	cow.InitAnimals("grass", "walk", "moo")

	// print what the cow eats.
	cow.Eat()

	// print when the cow moves.
	cow.Move()

	// print when the cow speaks.
	cow.Speak()

	cow.Behavior("eat")
}
