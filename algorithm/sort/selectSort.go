package sort

import "fmt"

/*
选择排序的原理：
通过将未排序部分的最小（或最大）元素选择出来放到已排序部分的后面，从而逐步构建有序序列。

具体步骤如下：
1、初始状态: 从待排序的序列中找出最小（或最大）元素。
2、选择最小元素: 将找到的最小元素与未排序序列的第一个元素交换位置，这样第一个元素就被放置到了已排序的区域。
3、缩小范围: 然后，再在剩余的未排序元素中重复步骤1和步骤2，找到新的最小元素并进行交换。
4、重复过程: 重复上述过程，直到所有元素都被排序完成。

缺点：性能上不如许多其他排序算法
优点：算法简单且不需要额外的存储空间（原地排序）

时间复杂度：O(n^2）
空间复杂度：O(1)
稳定性：不稳定

适用场景：
适用于少量数据排序，不适用于大量数据排序。
*/

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
