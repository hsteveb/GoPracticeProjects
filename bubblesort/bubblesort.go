package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

func main() {

	/* Slice array of size 10 */
	var array = make([]int, 10)

	/* Reader used to scan the line from the terminal */
	reader := bufio.NewReader(os.Stdin)

	/* Read the string and remove the new line character */
	fmt.Println("Please Enter up to 10 Int numbers(use spaces)")
	var line, e = reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")

	if e != nil {
		fmt.Println("Error was found")
		fmt.Println(e)
	} else {
		lines := strings.Split(line, " ")
		for i := 0; i < len(lines); i++ {
			var value, e2 = strconv.Atoi(lines[i])
			if e2 != nil {
				fmt.Println("Error was found")
				fmt.Println(e2)
				return
			}
			array[i] = value
		}

		/* a slice that gets the actual length of the array */
		var arraySlice = array[0:len(lines)]

		fmt.Printf("Unsorted Array: %v\n", arraySlice)
		BubbleSort(arraySlice)
		fmt.Printf("Sorted Array: %v\n", arraySlice)
	}

}

/* Sort the array in increasing order */
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if j+1 < len(array) {
				if array[j] > array[j+1] {
					Swap(array, j)
				}
			}
		}
	}
}

/* Change position of two elements */
func Swap(array []int, position int) {
	var tempElement = array[position]
	array[position] = array[position+1]
	array[position+1] = tempElement
}
