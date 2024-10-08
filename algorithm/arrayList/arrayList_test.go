package arrayList

import (
	"testing"
)

// 测试 ArrayList 的单元测试
func TestArrayList(t *testing.T) {
	list := &ArrayList{}

	// 测试添加元素
	t.Log("测试添加元素")
	list.Add(1)
	list.Add("two")
	list.Add(3.0)
	list.Print()
	if list.size != 3 {
		t.Errorf("期望大小为3，实际大小为%d", list.size)
	}

	// 测试按索引删除元素，正常情况
	t.Log("测试按索引删除元素，正常情况")
	removedElement := list.RemoveByIndex(1)
	if removedElement != "two" {
		t.Errorf("期望删除元素为'two'，实际删除元素为'%v'", removedElement)
	}
	list.Print()
	if list.size != 2 {
		t.Errorf("期望大小为2，实际大小为%d", list.size)
	}

	// 测试按索引删除元素，边界情况（删除最后一个元素）
	t.Log("测试按索引删除元素，边界情况（删除最后一个元素）")
	removedElement = list.RemoveByIndex(1)
	if removedElement != 3.0 {
		t.Errorf("期望删除元素为3.0，实际删除元素为'%v'", removedElement)
	}
	list.Print()
	if list.size != 1 {
		t.Errorf("期望大小为1，实际大小为%d", list.size)
	}

	// 测试按索引删除元素，边界情况（索引超出范围）
	t.Log("测试按索引删除元素，边界情况（索引超出范围）")
	removedElement = list.RemoveByIndex(10)
	if removedElement != nil {
		t.Errorf("期望删除元素为nil，实际删除元素为'%v'", removedElement)
	}
	list.Print()

	// 测试按值删除元素，正常情况
	t.Log("测试按值删除元素，正常情况")
	success := list.RemoveByValue(1)
	list.Print()
	if !success {
		t.Errorf("期望删除元素1成功，实际删除结果为false")
	}
	if list.size != 0 {
		t.Errorf("期望大小为0，实际大小为%d", list.size)
	}

	// 测试按值删除元素，边界情况（删除不存在的元素）
	t.Log("测试按值删除元素，边界情况（删除不存在的元素）")
	success = list.RemoveByValue(2)
	list.Print()
	if success {
		t.Errorf("期望删除不存在的元素失败，实际删除结果为true")
	}
}

func TestReverseArrayList(t *testing.T) {
	list := &ArrayList{}
	list.Add(1)
	list.Add("two")
	list.Add(3.0)
	list.Add(4)
	list.Add("five")
	list.Add(6.0)
	list.Reverse()
	list.Print()
	if list.size != 3 {
		t.Errorf("期望大小为3，实际大小为%d", list.size)
	}
	list.Reverse()
	list.Print()
}
