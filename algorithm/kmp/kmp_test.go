package kmp

import (
	"testing"
)

// 测试 KMP 函数
func TestKMP(t *testing.T) {
	// Happy Path 测试用例
	tests := []struct {
		text     string
		pattern  string
		expected int
		desc     string
	}{
		{"hello world", "world", 6, "正常匹配"},   		// 正常匹配
		{"abcabcabcd", "abcd", 6, "匹配在文本的末尾"}, 	// 匹配在文本的末尾
		{"aaaaa", "aa", 0, "多个重复字符"},          	// 多个重复字符
		{"aaaaa", "aaaaa", 0, "完全匹配"},         		// 完全匹配
		{"abcdef", "cd", 2, "中间匹配"},           		// 中间匹配
	}

	// 执行 Happy Path 测试
	for _, tt := range tests {
		t.Log("执行 Happy Path 测试：" + tt.desc)
		t.Run(tt.text+"_"+tt.pattern, func(t *testing.T) {
			result := KMP(tt.text, tt.pattern)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}

	// 边界情况测试用例
	edgeCases := []struct {
		text     string
		pattern  string
		expected int
		desc     string
	}{
		{"", "", 0, "空文本和空模式"},                   	// 空文本和空模式
		{"", "a", -1, "空文本，非空模式"},                // 空文本，非空模式
		{"abc", "", 0, "非空文本，空模式"},               // 非空文本，空模式
		{"abc", "abcd", -1, "模式比文本长"},            	// 模式比文本长
		{"abcabc", "abcabc", 0, "完全匹配"},          	// 完全匹配
		{"thequickbrownfox", "quick", 3, "中间匹配"}, 	// 中间匹配
	}

	// 执行边界情况测试
	for _, tt := range edgeCases {
		t.Log("执行边界情况测试：" + tt.desc)
		t.Run(tt.text+"_"+tt.pattern, func(t *testing.T) {
			result := KMP(tt.text, tt.pattern)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
