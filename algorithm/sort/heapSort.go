package sort

import (
	"container/heap"
	"fmt"
	"time"
)

// HeapSort 堆排序(升序)
func HeapSort(nums []int) {
	n := len(nums)
	if n < 1 {
		return
	}
	// 1、构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		adjustMaxHeap(nums, n, i)
		//adjustMinHeap(nums, n, i)
	}
	fmt.Println("1、构建完成的最大堆：", nums)
	fmt.Println("2、循环将首位（最大值）与未排序数据末尾交换，然后重新调整最大堆...")
	// 2、循环将首位（最大值）与未排序数据末尾交换，然后重新调整最大堆
	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // 首位与末尾交换
		fmt.Printf("=== 第%d次交换, 首位 [%d]%d 与末尾 [%d]%d 交换后：%v\n", n-i, 0, nums[0], i, nums[i], nums)
		adjustMaxHeap(nums, i, 0) // 重新调整最大堆
		//adjustMinHeap(nums, i, 0)
	}
	return
}

// 调整使之成为最大堆（降序）
func adjustMaxHeap(nums []int, n, i int) {
	fmt.Printf("---> 调整前的最大堆[%d-%d]：%v\n", n, i, nums)
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
		fmt.Printf("父节点不是最大值，则交换父节点 [%d]%d 与最大值 [%d]%d 交换，调整后的最大堆:%v\n", i, nums[i], maxIndex, nums[maxIndex], nums)
		adjustMaxHeap(nums, n, maxIndex)
	}
	fmt.Printf("<--- 重新调整后的最大堆[%d-%d]：%v\n", n, i, nums)
}

// 调整使之成为最小堆（升序）
func adjustMinHeap(nums []int, n, i int) {
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

/////////////////////////////////////////////////////////////////
// 用小根堆实现超时缓存机制

var cache map[int]*Node

const timeout = 10 // 超时时间

func InitTimeoutCache() {
	cache = make(map[int]*Node)
}

type Node struct {
	deadline int64
	value    interface{}
}

type TimerHeap []*Node

func (pq TimerHeap) Len() int {
	return len(pq)
}

func (pq TimerHeap) Less(i, j int) bool {
	return pq[i].deadline < pq[j].deadline
}

func (pq TimerHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TimerHeap) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *TimerHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func tesTimeoutCache() {
	InitTimeoutCache()
	pq := make(TimerHeap, 0, 5)
	heap.Init(&pq)

	for i := 0; i < 10; i++ {
		node := &Node{
			deadline: time.Now().Unix() + int64(timeout),
			value:    i,
		}
		cache[i] = node
		heap.Push(&pq, node)
		time.Sleep(time.Millisecond * 20)
	}
	ticker := time.NewTicker(time.Millisecond * 5)
	for {
		<-ticker.C
		for {
			currentTimestamp := time.Now().UnixNano()
			if len(pq) <= 0 {
				break
			}
			first := pq[0]
			if currentTimestamp < first.deadline {
				break
			} else {
				delete(cache, first.value.(int))
				heap.Pop(&pq)
				fmt.Println("delete key:", *first)
			}
		}
	}
}
