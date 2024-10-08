package sort

import "fmt"

// InsertSort 插入排序（升序）
func InsertSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	// 当前待排序元素，该元素之前的元素均已排序
	var currentValue int
	for i := 0; i < n; i++ {
		// 已被排序的元素的索引
		preIndex := i
		// 检查当前待排序元素是否超过元素的范围
		if preIndex+1 > n-1 {
			break
		}
		currentValue = nums[preIndex+1]
		fmt.Printf("第[%d]轮插入排序，待排序元素[%d]:%d，已被排序的元素索引:%d，待排数组为：%v\n",
			i, i+1, currentValue, preIndex, nums)
		/*
			在已被排序过数据中倒序查找插入位置，如果当前待排序元素小于已排序元素，
			则将已排序元素后移一位，直到找到插入位置
		*/
		for preIndex >= 0 && currentValue < nums[preIndex] {
			// 元素后移一位
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		// 循环结束后，说明已经找到了插入位置，将当前待排序元素插入到已排序元素的后面
		nums[preIndex+1] = currentValue
		fmt.Printf("第[%d]轮插入排序，排序后的数组为：%v\n", i,nums)

	}
	return nums
}
