package container

import (
	"container/heap"
	"fmt"
)

/*
//heap.interface
type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

//sort.Interface
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

方法列表：
func Init(h Interface)
	参数列表：
	- h：实现了heap.Interface接口的堆对象
	功能说明：在对堆h进行操作前必须保证堆已经初始化（即符合堆结构），该方法可以在堆中元素的顺序不符合堆的要求时调用，调用后堆会调整为标准的堆结构，该方法的时间复杂度为：O(n)，n为堆h中元素的总个数。

func Pop(h Interface) interface{}
	参数列表：
	- h：实现了heap.Interface的堆对象
	返回值：
	- interface{}：堆顶的元素
	功能说明：从堆h中取出堆顶的元素并自动调整堆结构。根据h的Less方法实现的不同，堆顶元素可以是最大的元素或者是最小的元素。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

func Push(h Interface, x interface{})
	参数列表：
	- h：实现了heap.Interface的堆对象
	- x：将被存到堆中的元素对象
	功能说明：把元素x存到堆中。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。

func Remove(h Interface, i int) interface{}
	参数列表：
	- h：实现了heap.Interface的堆对象
	- i：将被移除的元素在堆中的索引号
	返回值：interface{}：堆顶的元素
	功能说明：把索引号为i的元素从堆中移除。该方法的时间复杂度为O(log(n))，n为堆中元素的总和。
*/

type myHeap []int // 定义一个堆，存储结构为数组

// Less 实现了heap.Interface中组合的sort.Interface接口的Less方法
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Swap 实现了heap.Interface中组合的sort.Interface接口的Swap方法
func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Len 实现了heap.Interface中组合的sort.Interface接口的Push方法
func (h *myHeap) Len() int {
	return len(*h)
}

// Pop 实现了heap.Interface的Pop方法
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// Push 实现了heap.Interface的Push方法
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

// 按层来遍历和打印堆数据，第一行只有一个元素，即堆顶元素
func (h myHeap) printHeap(desc string) {
	n := 1
	levelCount := 1
	fmt.Println("------- ", desc)
	for n <= h.Len() {
		fmt.Println(h[n-1 : n-1+levelCount])
		n += levelCount
		levelCount *= 2
	}
}

func MyHeap1() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap)
	// 用堆本身的Push方法将数组中的元素依次存入堆中
	for _, value := range data {
		aheap.Push(value)
	}

	// 此时堆数组内容为：13, 25, 1, 9, 5, 12, 11
	// 不是正确的堆结构
	aheap.printHeap("aheap.Push：不是正确的堆结构")
	// 输出：
	//  [13]
	//  [25 1]
	//  [9 5 12 11]

	heap.Init(aheap) // 对堆进行调整，调整后为规范的堆结构

	fmt.Println(*aheap) // 输出：[1 5 11 9 25 12 13]
	aheap.printHeap("aheap.Push：对堆进行调整，调整后为规范的堆结构")
	// 输出：
	//	[1]
	//	[5 11]
	//	[9 25 12 13]
}

func MyHeap2() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap) // 创建空堆

	// 用heap包中的Push方法将数组中的元素依次存入堆中，
	// 每次Push都会保证堆是规范的堆结构
	for _, value := range data {
		heap.Push(aheap, value)
	}
	fmt.Println("*aheap: ", *aheap) // 输出：[1 5 11 25 9 13 12]
	aheap.printHeap("heap.Push：每次Push都会保证堆是规范的堆结构")
	// 输出：
	//  [1]
	//  [5 11]
	//  [25 9 13 12]

	// 依次调用heap包的Pop方法来获取堆顶元素
	fmt.Println("依次调用heap包的Pop方法来获取堆顶元素：")
	for aheap.Len() > 0 {
		value := heap.Pop(aheap)
		aheap.printHeap(fmt.Sprintf("heap.Pop:%d, pop value:%d", aheap.Len(), value))
	}
	// 输出：1 5 9 11 12 13 25
	fmt.Println()
}

func MyHeap3() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap) // 创建空堆

	// 用heap包中的Push方法将数组中的元素依次存入堆中，
	// 每次Push都会保证堆是规范的堆结构
	for _, value := range data {
		heap.Push(aheap, value)
	}
	fmt.Println(*aheap) // 输出：[1 5 11 25 9 13 12]
	aheap.printHeap("heap.Push：每次Push都会保证堆是规范的堆结构")
	// 输出：
	//  [1]
	//  [5 11]
	//  [25 9 13 12]

	value := heap.Remove(aheap, 2) // 删除索引号为2的元素（即数组中的第3个元素）
	fmt.Println("remove index 2 element: ", value)             // 输出：11
	aheap.printHeap("删除索引号为2的元素（即数组中的第3个元素）")
	// 输出：
	//	[1]
	//	[5 12]
	//	[25 9 13]
}


