package main

import "fmt"
import "./animal"
import "bufio"
import "os"
import "strings"

func main() {

	// Create scanner for stdin
	scanner := bufio.NewReader(os.Stdin)

	var animals = make([]animal.Animal, 0)
	//fmt.Println(len(animals))

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

		if len(inputValues) == 3 {
			switch strings.ToLower(inputValues[0]) {
			case "newanimal":
				if animal.FindAnimal(inputValues[1], animals) != -1 {
					fmt.Println("An animal already has that name, try again. ")
				} else {
					var animal = animal.CreateAnimal(inputValues[1], inputValues[2])
					animals = append(animals, animal)
				}
			case "query":
				var index = animal.FindAnimal(inputValues[1], animals)
				if index == -1 {
					fmt.Println("Animal does not exist.")
				} else {
					animal.AnimalBehavior(animals[index], inputValues[2])
				}
			default:
			}
		}

	}
}
