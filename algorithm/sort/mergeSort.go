package sort

import "fmt"

// MergeSort 归并排序：https://www.bilibili.com/video/BV1em1oYTEFf?spm_id_from=333.788.videopod.sections&vd_source=7459db3060f4a09b27ad55e8805e9f7c
func MergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	mid := n / 2 // 取中间位置
	left := MergeSort(nums[:mid]) // 递归左边
	right := MergeSort(nums[mid:]) // 递归右边
	result := merge(left, right) // 合并
	fmt.Printf("归并排序前的数组：%v, L-%v, R-%v,合并数组：L-%v, R-%v, 归并排序后的数组：%v\n", nums, nums[:mid], nums[mid:], left, right, result)
	return result
}

func merge(left, right []int) []int {
	if len(left) == 0 { // 左边数组为空, 直接返回右边数组
		return right
	}
	if len(right) == 0 { // 右边数组为空, 直接返回左边数组
		return left
	}

	result := make([]int, 0)
	l, r := 0, 0
	for l < len(left) && r < len(right) { // 合并两个数组
		if left[l] <= right[r] { // 左边数组的元素小于等于右边数组的元素
			result = append(result, left[l]) // 追加左边数组的元素
			l++ // 左边数组指针后移
		} else { // 左边数组的元素大于右边数组的元素
			result = append(result, right[r]) // 追加右边数组的元素
			r++ // 右边数组指针后移
		}
	}

	if l < len(left) { // 左边数组还有剩余元素
		result = append(result, left[l:]...) // 追加左边数组剩余元素
	}
	if r < len(right) { // 右边数组还有剩余元素
		result = append(result, right[r:]...) // 追加右边数组剩余元素
	}
	return result
}
