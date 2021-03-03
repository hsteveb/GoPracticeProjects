package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "./sort"
import "sync"
import "strconv"

func main() {

	var wp sync.WaitGroup
	var bubblesortInit sort.BubbleSort
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter int numbers to be sorted, min 4.")
	fmt.Println("Ex: 1, 2, 3, 4")

	line, error := scanner.ReadString('\n')
	checkError(error)

	line = strings.Trim(line, "\n")
	stringArray := strings.Split(line, ", ")

	intArray, err := atoiArray(stringArray)
	checkError(err)

	twodArray := makeSubarrays(&intArray)

	for i := 0; i < 4; i++ {
		wp.Add(1)
		go doSorting(i, bubblesortInit, twodArray[i], &wp)
	}

	wp.Wait()

	mainArray := bubblesortInit.Init(intArray)
	mainArray.Sort()
	fmt.Printf("Sorted Array: %v\n", mainArray.GetArray())
}

/*Where concurrency happens to do the sorting of the arrays. */
func doSorting(name int, initAlgoritm sort.Algorithm, array []int, wp *sync.WaitGroup) {
	fmt.Printf("Thread %d sorting subarray: %v\n", name, array)
	var tempAlgorithm = initAlgoritm.Init(array)
	tempAlgorithm.Sort()
	wp.Done()
}

/*Creates a 2d array that partitions it into 4 subarrays.*/
func makeSubarrays(array *[]int) [][]int {

	subArrayLength := len(*array) / 4
	subArrayRemainder := len(*array) % 4
	var twodArray = make([][]int, 4)

	remainderCountdown := 0

	j := 0
	for i := 0; i < len(*array); {
		var size = subArrayLength
		if remainderCountdown < subArrayRemainder {
			size++
			remainderCountdown++
		}

		twodArray[j] = make([]int, size)
		twodArray[j] = (*array)[i : i+size]

		i = i + size
		j++
	}

	return twodArray
}

/*Convert string array to int array. */
func atoiArray(array []string) ([]int, error) {
	var len = len(array)
	intArray := make([]int, len)

	for i := 0; i < len; i++ {
		value, err := strconv.Atoi(array[i])

		if err != nil {
			return []int{}, err
		}
		intArray[i] = value
	}
	return intArray, nil
}

/*checks if we found an error and exit the problem if we do */
func checkError(err error) {
	if err != nil {
		fmt.Printf("Error found %v \nexiting... \n", err)
		os.Exit(3)
	}
}
