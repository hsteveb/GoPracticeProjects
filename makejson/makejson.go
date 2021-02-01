package main

import "fmt"
import "bufio"
import "os"
import "encoding/json"
import "strings"

type info struct {
	name    string
	address string
}

func main() {

	var name string
	var address string
	var errorValue error
	scanner := bufio.NewReader(os.Stdin)

	/* Get name of the person using bufio since scanln doesn't work as intended */
	fmt.Println("Enter name")
	name, errorValue = scanner.ReadString('\n')

	/* Get address the same way as name */
	fmt.Println("Enter address")
	address, errorValue = scanner.ReadString('\n')

	/* check if we get any errors from doing the above stuff */
	if errorValue != nil {
		println("Error getting values")
		return
	}

	/* Remove new line characters from the strings */
	name = strings.Trim(name, "\n")
	address = strings.Trim(address, "\n")

	/* Make a map */
	var personInfo = map[string]string{"name": name, "address": address}

	/* Convert it to json */
	jsonValue, errorValue := json.Marshal(personInfo)

	/* Print json */
	fmt.Println(string(jsonValue))

}
