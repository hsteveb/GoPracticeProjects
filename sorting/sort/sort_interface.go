package sort

/*Algorithm Interface used for all the sorting algorithms */
type Algorithm interface {
	Init(array []int) Algorithm
	Sort()
	GetArray() []int
}
