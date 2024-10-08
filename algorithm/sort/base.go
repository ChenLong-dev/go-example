package sort

type DataNode struct {
	Data interface{}
}

type Sorter interface {
	BubbleSort(data []*DataNode, fn func(a, b *DataNode) bool) []*DataNode
	SelectionSort(data []DataNode) []DataNode
	InsertionSort(data []DataNode) []DataNode
	MergeSort(data []DataNode) []DataNode
	QuickSort(data []DataNode) []DataNode
}

func ReverseSort(data []int) []int {
	n := len(data)
	i, j := 0, n-1
	for i<j{
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return data
}