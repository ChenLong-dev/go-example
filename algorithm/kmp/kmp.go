package kmp

// KMP 实现KMP算法
func KMP(text, pattern string) int {
	if pattern == "" { // 空字符串匹配任意字符串
		return 0
	}
	next := getNext(pattern)           // 获取next数组
	for i, j := 0, 0; i < len(text); { // 遍历text
		if pattern[j] == text[i] { // 匹配成功
			i++
			j++
		} else if j > 0 { // 匹配失败，j>0，j回退
			j = next[j-1]
		} else { // 匹配失败，j=0，i+1
			i++
		}

		if j == len(pattern) { // 匹配成功，返回匹配位置
			return i - j
		}
	}
	return -1
}

// getNext 获取next数组
func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	prefixLen := 0
	for i := 1; i < len(pattern); { // 遍历pattern
		if pattern[i] == pattern[prefixLen] { // 匹配成功
			prefixLen++         // 前缀长度+1
			next[i] = prefixLen // next[i]记录前缀长度
			i++
		} else { // 匹配失败，前缀长度回退
			if prefixLen == 0 { // 前缀长度为0，next[i]记录0
				next[i] = 0 // next[i]记录0
				i++
			} else {
				prefixLen = next[prefixLen-1] // 前缀长度回退
			}
		}
	}
	return next
}
