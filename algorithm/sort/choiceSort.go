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

// ChoiceSort 选择排序（升序）
func ChoiceSort(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	for i := 0; i < n; i++ {
		// 寻找最小值索引，每个循环开始总是假设第一个元素是最小值
		minIndex := i
		for j := i; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		// 交换最小值和当前元素
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
		fmt.Printf("第[%d]轮选择排序，最小值索引为:%d，交换元素后数组为:%v\n", i, minIndex, nums)
	}

	return nums
}
