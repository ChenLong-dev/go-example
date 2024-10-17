package sort

import "fmt"

// ShellSort 希尔排序：https://www.bilibili.com/video/BV1bm42137UZ?spm_id_from=333.788.videopod.sections&vd_source=7459db3060f4a09b27ad55e8805e9f7c
func ShellSort(nums []int) {
	n := len(nums)
	gap := n / 2  // 初始步长为n/2
	for gap > 0 { // 步长不断缩小
		count := 0
		flag := false              // 标记是否进行过交换
		for i := gap; i < n; i++ { // 遍历数组
			temp := nums[i] // 保存当前元素
			j := i
			for j >= gap && nums[j-gap] > temp { // 遍历前一个元素
				nums[j] = nums[j-gap] // 交换位置
				j -= gap              // 前移一位
				flag = true           // 标记交换
			}
			nums[j] = temp // 插入当前元素
			fmt.Printf("第[%d]轮希尔排序，temp:%d, gap:%d, 排序后的数组为：%v, 是否进行过交换：%v\n", count, temp, gap, nums, flag)
			count += 1
			flag = false // 重置标记
		}
		gap = gap / 2 // 步长缩小一半

	}
	return
}
