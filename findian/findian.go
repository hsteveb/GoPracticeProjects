package main

import "fmt"
import "strings"

func main() {
	var s1 string

	fmt.Println("Enter a string.")

	/*Scan a string with whitespaces but you need to supply double quote marks to make it work*/
	fmt.Scanf("%q", &s1)

	s1 = strings.ToLower(s1)
	containsA := strings.Index(s1, "a")
	length := len(s1)

	/* Checks to see if the string contains 'i', 'a', and 'n' */
	if s1[0] == 'i' && s1[length-1] == 'n' && containsA > 0 && containsA < length-1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
