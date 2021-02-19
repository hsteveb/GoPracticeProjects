package animal

import "strings"

/*Function used to find the animal in the array using the name.
If found it returns the index and if not it returns -1.*/
func FindAnimal(name string, animals []Animal) int {
	var titleName = strings.Title(strings.ToLower(name))
	for index, animal := range animals {
		if animal.GetName() == titleName {
			return index
		}
	}
	return -1
}

/*Function used to make an animal*/
func CreateAnimal(name string, animalType string) Animal {
	var titleName = strings.Title(strings.ToLower(name))

	switch strings.ToLower(animalType) {
	case "cow":
		var cow Cow
		return cow.Init(titleName)
	case "snake":
		var snake Snake
		return snake.Init(titleName)
	case "bird":
		var bird Bird
		return bird.Init(titleName)
	default:
		var nil NilAnimal
		return nil.Init("Nil")
	}
}

/*Function used to call the behavior function from the animal*/
func AnimalBehavior(animal Animal, behavior string) {
	switch strings.ToLower(behavior) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
	}
}
