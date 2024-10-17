package sort

import (
	"fmt"
	"reflect"
	"testing"
)


var dst = []int{7, 16, 20, 27, 28, 36, 36, 44, 55, 60, 67}

// 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	t.Log("冒泡排序测试")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	BubbleSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("冒泡排序测试失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试选择排序
//func TestChoiceSort(t *testing.T) {
//	t.Log("测试选择排序")
//	fmt.Println("原始数组：", src)
//	data := ChoiceSort(src)
//	fmt.Println("排序后数组：", data)
//	if !reflect.DeepEqual(src, dst) {
//		t.Errorf("冒泡排序测试失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
//	}
//}

// 测试插入排序
//func TestInsertSort(t *testing.T) {
//	t.Log("测试插入排序")
//	fmt.Println("原始数组：", src)
//	data := InsertSort(src)
//	fmt.Println("排序后数组：", data)
//	if !reflect.DeepEqual(src, dst) {
//		t.Errorf("冒泡排序测试失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
//	}
//
//}

func TestDirectInsertSort(t *testing.T) {
	t.Log("测试直接插入排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	DirectInsertSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试直接插入排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}
func TestDimidiatedInsertSort(t *testing.T) {
	t.Log("测试折半插入排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	DimidiatedInsertSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试折半插入排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试快速排序
func TestQuickSort(t *testing.T) {
	t.Log("测试快速排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	QuickSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试快速排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试选择排序
func TestSelectSort(t *testing.T) {
	t.Log("测试选择排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	SelectSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试选择排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试堆排序
func TestHeapSort(t *testing.T) {
	t.Log("测试堆排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	HeapSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试堆排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试希尔排序
func TestShellSort(t *testing.T) {
	t.Log("测试希尔排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	ShellSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试希尔排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试合并排序
func TestMergeSort(t *testing.T) {
	t.Log("测试合并排序")
	var src = []int{36, 27, 20, 60, 55, 7, 28, 36, 67, 44, 16}
	fmt.Println("原始数组：", src)
	src = MergeSort(src)
	fmt.Println("排序后数组：", src)
	if !reflect.DeepEqual(src, dst) {
		t.Errorf("测试合并排序失败，结果不正确，结果为：%v， 应该等于：%v", src, dst)
	}
}

// 测试用小根堆实现超时缓存机制
//func TestTimeoutCache(t *testing.T) {
//	t.Log("测试超时缓存机制")
//	tesTimeoutCache()
//	t.Log("测试超时缓存机制成功")
//}
