package stack

import (
	"testing"
)

// 测试栈的功能
func TestStack(t *testing.T) {
	// 创建一个新的栈
	stack := NewStack()

	// 测试栈的初始状态
	t.Log("测试栈的初始状态")
	if !stack.IsEmpty() {
		t.Error("栈应该是空的")
	}
	stack.Print()
	if size := stack.Size(); size != 0 {
		t.Errorf("栈的大小应该是 0, 但实际是 %d", size)
	}

	// 测试压入元素
	t.Log("测试压入元素")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)
	stack.Print()

	// 测试栈的状态
	t.Log("测试栈的状态")
	if stack.IsEmpty() {
		t.Error("栈不应该是空的")
	}
	if size := stack.Size(); size != 8 {
		t.Errorf("栈的大小应该是 8, 但实际是 %d", size)
	}

	// 测试栈顶元素
	t.Log("测试栈顶元素")
	if top := stack.Peek(); top != 8 {
		t.Errorf("栈顶元素应该是 8, 但实际是 %v", top)
	}

	// 测试弹出元素
	t.Log("测试弹出元素-8")
	if item := stack.Pop(); item != 8 {
		t.Errorf("弹出的元素应该是 8, 但实际是 %v", item)
	}
	stack.Print()
	if size := stack.Size(); size != 7 {
		t.Errorf("栈的大小应该是 7, 但实际是 %d", size)
	}

	// 测试再次弹出元素
	t.Log("测试弹出元素-7")
	if item := stack.Pop(); item != 7 {
		t.Errorf("弹出的元素应该是 7, 但实际是 %v", item)
	}
	stack.Print()
	if item := stack.Pop(); item != 6 {
		t.Errorf("弹出的元素应该是 6, 但实际是 %v", item)
	}

	// 测试弹空栈
	t.Log("测试弹空栈-6/5/4/3/2/1")
	stack.Pop() // 6
	stack.Pop() // 5
	stack.Pop() // 4
	stack.Pop() // 3
	stack.Pop() // 2
	if item := stack.Pop(); item != nil { // 1
		t.Errorf("弹空栈时应该返回 nil, 但实际是 %v", item)
	}
	stack.Print()

	// 测试栈在弹空后状态
	t.Log("测试栈在弹空后状态")
	if !stack.IsEmpty() {
		t.Error("栈应该是空的")
	}
	if size := stack.Size(); size != 0 {
		t.Errorf("栈的大小应该是 0, 但实际是 %d", size)
	}
}
