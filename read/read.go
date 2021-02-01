package main

import "fmt"
import "os"
import "bufio"
import "strings"

/* struct used to put the name of the person */
type name struct {
	firstName string
	lastName  string
}

func main() {

	/* open the file */
	file, error := os.Open("names.txt")

	/* if file cannot open, terminate early */
	if error != nil {
		fmt.Println("Cannot open file")
		return
	}

	/* create scanner so that we can get the names from the file */
	scanner := bufio.NewScanner(file)

	/* create slice with nothing */
	names := make([]name, 0)

	/* scan all the lines and put the values in structs */
	for scanner.Scan() {

		line := scanner.Text()
		lines := strings.Split(line, " ")
		names = append(names, name{firstName: lines[0], lastName: lines[1]})

	}

	/* print all the struct values by putting the first and last names together */
	for _, temp := range names {
		fmt.Printf("%s %s\n", temp.firstName, temp.lastName)
	}

	/* close the file */
	file.Close()
}
