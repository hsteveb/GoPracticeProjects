package animal

import "testing"
import "reflect"
import "fmt"

var name1 = "Bill"
var name2 = "Django"
var name3 = "Luis"
var cowInitializer Cow
var snakeInitializer Snake
var birdInitializer Bird
var cowString = "cow"
var snakeString = "snake"
var birdString = "bird"
var eat = "eat"
var move = "move"
var speak = "speak"

func TestFindAnimal(t *testing.T) {

	var animals = []Animal{cowInitializer.Init(name1), snakeInitializer.Init(name2), birdInitializer.Init(name3)}

	if FindAnimal(name1, animals) != 0 {
		t.Fail()
		t.Logf("Expected %d, got %d", 0, FindAnimal(name1, animals))
	}

	if FindAnimal(name3, animals) != 2 {
		t.Fail()
		t.Logf("Expected %d, got %d", 2, FindAnimal(name3, animals))
	}

	if FindAnimal(name2, animals) != 1 {
		t.Fail()
		t.Logf("Expected %d, got %d", 1, FindAnimal(name2, animals))
	}
}

func TestCreateAnimal(t *testing.T) {
	/*Create cow using CreateAnimal function*/
	var cowUsingInit = cowInitializer.Init(name1)
	var cow = CreateAnimal(name1, cowString)
	_, ok := cow.(Cow)
	if !ok {
		t.Logf("Expected %s, got %s", reflect.TypeOf(cowUsingInit).String(), reflect.TypeOf(cow).String())
	}

	/*Create snake using CreateAnimal function*/
	var snakeUsingInit = snakeInitializer.Init(name1)
	var snake = CreateAnimal(name1, snakeString)
	_, ok = snake.(Snake)
	if !ok {
		t.Logf("Expected %s, got %s", reflect.TypeOf(snakeUsingInit).String(), reflect.TypeOf(snake).String())
	}

	/*Create bird using CreateAnimal function*/
	var birdUsingInit = birdInitializer.Init(name1)
	var bird = CreateAnimal(name1, birdString)
	_, ok = bird.(Bird)
	if !ok {
		t.Logf("Expected %s, got %s", reflect.TypeOf(birdUsingInit).String(), reflect.TypeOf(bird).String())
	}
}

func TestAnimalBehavior(t *testing.T) {

	fmt.Println("Testing cow behavior")
	var cow = CreateAnimal(name1, cowString)
	AnimalBehavior(cow, move)
	AnimalBehavior(cow, speak)
	AnimalBehavior(cow, eat)
	fmt.Println()

	fmt.Println("Testing snake behavior")
	var snake = CreateAnimal(name1, snakeString)
	AnimalBehavior(snake, move)
	AnimalBehavior(snake, speak)
	AnimalBehavior(snake, eat)
	fmt.Println()

	fmt.Println("Testing bird behavior")
	var bird = CreateAnimal(name1, birdString)
	AnimalBehavior(bird, move)
	AnimalBehavior(bird, speak)
	AnimalBehavior(bird, eat)
}
