package main

import "fmt"
import "./animal"
import "bufio"
import "os"
import "strings"

func main() {

	// Create scanner for stdin
	scanner := bufio.NewReader(os.Stdin)

	// Create Cow
	var cow animal.Animal
	cow.InitAnimals("grass", "walk", "moo")

	// Create Bird
	var bird animal.Animal
	bird.InitAnimals("worms", "fly", "peep")

	// Create Snake
	var snake animal.Animal
	snake.InitAnimals("mice", "slither", "hsss")

	var input string = ""
	var error error

	fmt.Println("Type \"exit\" to quit.")

	for input != "exit" {
		fmt.Print(">")
		input, error = scanner.ReadString('\n')

		if error != nil {
			print("Error found exiting....")
			return
		}

		input = strings.Trim(input, "\n")
		inputValues := strings.Split(input, " ")

		if len(inputValues) == 2 {
			if inputValues[0] == "cow" {
				cow.Behavior(inputValues[1])
			} else if inputValues[0] == "bird" {
				bird.Behavior(inputValues[1])
			} else {
				snake.Behavior(inputValues[1])
			}
		}

	}
}
