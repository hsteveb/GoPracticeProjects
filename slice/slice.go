/* Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer. */

package main

import "fmt"
import "sort"
import "strconv"

func main() {
	// create the array slice
	var array = make([]int, 3)

	// count used to keep track of the number of elements in the array
	var count = 0

	// used to capture
	var errorValue error

	// grab the input value from scan
	var input string = ""

	// "while" loop used to get the values and put them on the array and sort the array in increasing order
	for input != "X" {
		var tempInt int
		fmt.Scanln(&input)

		/* create new array if the count is the same size as the capacity of the array slice */
		if count == cap(array) {
			array = increaseArray(cap(array)+3, array)
		}
		tempInt, errorValue = strconv.Atoi(input)

		/* print an error that the string wasn't able to be converted to an int */
		if errorValue != nil {
			print("Could not convert string value\n")
		} else {
			/* If string is able to be converted then put it inside the array, sort the array, increase the count, and print the sorted array */
			array = array[:count+1]
			array[count] = tempInt
			sort.Ints(array)
			count++

			fmt.Printf("Array: %v\n", array)
		}

	}
}

/* Function used to create a new array, pass the old info into the new one, and return the new array*/
func increaseArray(newCapacity int, array []int) []int {
	var tempArray = make([]int, newCapacity)
	for i := 0; i < len(array); i++ {
		tempArray[i] = array[i]
	}
	return tempArray
}
