package sort

import "testing"
import "reflect"

var initBubbleSort BubbleSort

func Test_BubbleSort_GetArray(t *testing.T) {

	var intArray = []int{1, 2}
	var bubble = initBubbleSort.Init(intArray)

	if !reflect.DeepEqual(intArray, bubble.GetArray()) {
		t.Fail()
		t.Logf("expected %v, got %v", intArray, bubble.GetArray())
	}
}

func Test_BubbleSort_Sort(t *testing.T) {
	var intArray = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var intSortedArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var bubble = initBubbleSort.Init(intArray)

	bubble.Sort()

	if !reflect.DeepEqual(intSortedArray, bubble.GetArray()) {
		t.Fail()
		t.Logf("expected %v, got %v", intSortedArray, bubble.GetArray())
	}

}
