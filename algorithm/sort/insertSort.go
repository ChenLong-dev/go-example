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
		fmt.Printf("第[%d]轮插入排序，排序后的数组为：%v\n", i, nums)

	}
	return nums
}

// DirectInsertSort 直接插入排序（升序）：// DimidiatedInsertSort 折半插入排序：https://www.bilibili.com/video/BV1tf421Q7eh?spm_id_from=333.788.videopod.sections&vd_source=7459db3060f4a09b27ad55e8805e9f7c
func DirectInsertSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ { // 从第二个元素开始
		temp := nums[i]
		j := i - 1
		for ; j >= 0; j-- { // 从第一个元素开始，向前比较
			if nums[j] > temp { // 如果当前元素大于待插入元素，则将当前元素后移一位(升序)
				nums[j+1] = nums[j] // 向右移动元素
			} else { // 如果当前元素小于或等于待插入元素，则停止比较
				break
			}
		}
		nums[j+1] = temp // 插入待插入元素
		fmt.Printf("第[%d]轮直接插入排序，temp:%d, 排序后的数组为：%v\n", i, temp, nums)
	}
	return
}

// DimidiatedInsertSort 折半插入排序（升序）：https://www.bilibili.com/video/BV1E1421b7Eb?spm_id_from=333.788.videopod.sections&vd_source=7459db3060f4a09b27ad55e8805e9f7c
func DimidiatedInsertSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ { // 从第二个元素开始
		temp := nums[i]     // 保存要交换的元素
		left := 0           // 左指针
		right := i - 1      // 右指针
		for left <= right { // 左右指针交叉比较，直到左指针大于右指针
			mid := (left + right) / 2 // 中间指针
			if nums[mid] <= temp {     // 当要排序的元素大于等于了中间值，则左指针右移
				left = mid + 1
			} else { // 当要排序的元素小于了中间值，则右指针左移
				right = mid - 1
			}
			//fmt.Println("mid:", mid, nums[mid])
		}
		for j := i - 1; j >= left; j-- { // 向左移动已排序的元素，直到遇到大于待插入元素的位置
			nums[j+1] = nums[j] // 右移元素
		}
		nums[left] = temp // 插入待插入元素
		fmt.Printf("第[%d]轮折半插入排序，temp:%d, 排序后的数组为：%v\n", i, temp, nums)
	}
	return
}
