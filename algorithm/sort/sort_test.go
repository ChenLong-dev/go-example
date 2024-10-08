package sort

import (
	"fmt"
	"testing"
)

var arr = []int{35, 63, 48, 11, 86, 24, 53, 9}

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	t.Log("冒泡排序测试")
	fmt.Println("原始数组：", arr)
	data := BubbleSort(arr)
	fmt.Println("排序后数组：", data)
}

// 测试选择排序
func TestChoiceSort(t *testing.T) {
	t.Log("测试选择排序")
	fmt.Println("原始数组：", arr)
	data := ChoiceSort(arr)
	fmt.Println("排序后数组：", data)
}

// 测试插入排序
func TestInsertSort(t *testing.T) {
	t.Log("测试插入排序")
	fmt.Println("原始数组：", arr)
	data := InsertSort(arr)
	fmt.Println("排序后数组：", data)
}

// 测试快速排序
func TestQuickSort(t *testing.T) {
	t.Log("测试快速排序")
	fmt.Println("原始数组：", arr)
	data := QuickSort(arr)
	fmt.Println("排序后数组：", data)
}

// 测试堆排序
func TestHeapSort(t *testing.T) {
	t.Log("测试堆排序")
	fmt.Println("原始数组：", arr)
	data := HeapSort(arr)
	fmt.Println("排序后数组：", data)
}
