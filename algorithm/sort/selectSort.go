package sort

import "fmt"

// SelectSort 选择排序：https://www.bilibili.com/video/BV1kjsuenE8v?spm_id_from=333.788.videopod.sections&vd_source=7459db3060f4a09b27ad55e8805e9f7c
func SelectSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ { // 遍历数组, 每轮循环都将最小值放到数组的第 i 个位置，共 n-1 轮
		minIndex := i                // 寻找最小值索引，每个循环开始总是假设第一个元素是最小值
		for j := i + 1; j < n; j++ { // 遍历数组，寻找最小值索引
			if nums[j] < nums[minIndex] { // 如果找到更小的元素，则更新最小值索引
				minIndex = j // 找到最小值索引
			}
		}
		if i == minIndex { // 如果最小值索引和当前索引相同，说明已经有序，不需要再交换
			continue
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i] // 交换最小值到数组的第 i 个位置
		fmt.Printf("第%d轮选择排序结果: %v，最小值索引为%d， 交换元素为：%d<->%d\n", i, nums, minIndex, nums[minIndex], nums[i])
	}
}
