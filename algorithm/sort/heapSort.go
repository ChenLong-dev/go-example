package sort

import "fmt"

// HeapSort 堆排序(升序)
func HeapSort(nums []int) []int {
	n := len(nums)
	if n < 1 {
		return nums
	}
	// 1、构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		adjustMaxHeap(nums, n, i)
		//adjustMinHeap(nums, n, i)
	}
	fmt.Println("1、构建完成的最大堆：", nums)
	fmt.Println("2、开始排序...")
	// 2、循环将首位（最大值）与未排序数据末尾交换，然后重新调整最大堆
	for i := n-1; i > 0; i-- {
		nums[0], nums[n-1] = nums[n-1], nums[0]
		adjustMaxHeap(nums, i, 0)
		//adjustMinHeap(nums, i, 0)
	}
	return nums
}

// 调整使之成为最大堆（降序）
func adjustMaxHeap(nums []int, n, i int){
	maxIndex := i
	left := 2*i + 1
	right := 2 * (i + 1)
	// 如果有左子树，且左子树大于父节点，将最大指针指向左子树
	if left < n && nums[left] > nums[maxIndex] {
		maxIndex = left
	}
	// 如果有右子树，且右子树大于父节点且大约左子树，将最大指针指向右子树
	if right < n && nums[right] > nums[maxIndex] && nums[right] > nums[left] {
		maxIndex = right
	}
	// 如果父节点不是最大值，则交换父节点与最大值交换，并递归调整与父节点交换的位置继续堆化
	if maxIndex != i {
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		fmt.Println("调整后的最大堆:", nums)
		adjustMaxHeap(nums, n, maxIndex)
	}
}

// 调整使之成为最小堆（升序）
func adjustMinHeap(nums []int, n, i int){
	maxIndex := i
	left := 2*i + 1
	right := 2 * (i + 1)
	// 如果有左子树，且左子树小于父节点，将最大指针指向左子树
	if left < n && nums[left] < nums[maxIndex] {
		maxIndex = left
	}
	// 如果有右子树，且右子树小于父节点且大约左子树，将最大指针指向右子树
	if right < n && nums[right] < nums[maxIndex] && nums[right] < nums[left] {
		maxIndex = right
	}
	// 如果父节点不是最大值，则交换父节点与最大值交换，并递归调整与父节点交换的位置继续堆化
	if maxIndex != i {
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		fmt.Println("调整后的最小堆:", nums)
		adjustMinHeap(nums, n, maxIndex)
	}
}