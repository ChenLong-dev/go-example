package sort

import (
	"fmt"
)

/*
快速排序原理：
1. 选择一个基准数，通常选择第一个元素或者最后一个元素，也可以选择中间元素。
2. 遍历数组，将小于等于基准数的元素放到左边，大于基准数的元素放到右边。
3. 递归地对左右两边的子数组进行相同的操作。

时间复杂度：O(nlogn)
空间复杂度：O(logn)
*/

// QuickSort 快速排序算法
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	return quickSort(nums, 0, len(nums)-1)
}

// 快速排序的递归函数
func quickSort(nums []int, start, end int) []int {
	if len(nums) < 1 || start <0 || end > len(nums) || start > end {
		return nil
	}
	// 确定分区指示器
	zoneIndex := partition(nums, start, end)
	if zoneIndex > start {
		quickSort(nums, start, zoneIndex-1)
	}
	if zoneIndex < end {
		quickSort(nums, zoneIndex+1, end)
	}
	fmt.Printf("本轮排序后的数组为：%v\n", nums[start:end+1])

	return nums
}

// 快速排序的分区函数
func partition(nums []int, start, end int) int {
	// 只有一个元素时，直接返回，不用分区
	if start == end {
		return start
	}
	// 随机选取一个元素作为分区指示器
	//pivot := start+rand.Intn(end-start+1)
	pivot := start
	// zoneIndex为分区指示器的位置，初始值为分区头元素下标减一
	zoneIndex := start - 1
	fmt.Printf("开始下标[%d]:%d，结束下标[%d]:%d，基准数下标[%d]:%d，分区指示器为:%d\n",
		start, nums[start], end, nums[end], pivot, nums[pivot], zoneIndex)

	// 将基准数和分区尾元素交换位置
	nums[pivot], nums[end] = nums[end], nums[pivot]
	// 遍历分区数组，将小于等于基准数的元素放到分区头，大于基准数的元素放到分区尾
	for i := start; i <= end; i++ {
		// 当前元素小于等于基准数
		if nums[i] <= nums[end] {
			// 分区指示器右移（累加1）
			zoneIndex++
			// 当元素在分区指示器的右边时，交换当前元素和分区指示器元素
			if i>zoneIndex {
				nums[i], nums[zoneIndex] = nums[zoneIndex], nums[i]
			}
		}
	}
	return zoneIndex
}
