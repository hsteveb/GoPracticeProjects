package sort

/*BubbleSort struct used to save the array*/
type BubbleSort struct {
	intArray []int
}

/*Init used to add an int to the bubblesort */
func (bs BubbleSort) Init(array []int) Algorithm {
	bs.intArray = array
	return bs
}

/*GetArray returns the int array of the bubblesort */
func (bs BubbleSort) GetArray() []int {
	return bs.intArray
}

/*Sort sort the array using bubblesort*/
func (bs BubbleSort) Sort() {
	for i := 0; i < len(bs.intArray); i++ {
		for j := 1; j < len(bs.intArray); j++ {
			if bs.intArray[j-1] > bs.intArray[j] {
				var temp = bs.intArray[j]
				bs.intArray[j] = bs.intArray[j-1]
				bs.intArray[j-1] = temp
			}
		}
	}
}
