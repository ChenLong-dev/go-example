package linkedList

import (
	"fmt"
	"testing"
)


func TestSingleLinkedList(t *testing.T) {
	// 创建单链表实例
	singleLinkedList := NewSingleLinkedList()

	t.Log("测试用例 1: 插入元素并检查长度")
	singleLinkedList.Inseart(1)
	singleLinkedList.Print()
	if singleLinkedList.Length() != 1 {
		t.Errorf("期望长度 1，但得到 %d", singleLinkedList.Length())
	}

	t.Log("测试用例 2: 再插入一个元素并检查链表内容")
	singleLinkedList.Inseart(2)
	singleLinkedList.Print()
	if singleLinkedList.Length() != 2 {
		t.Errorf("期望长度 2，但得到 %d", singleLinkedList.Length())
	}
	if !singleLinkedList.Find(2) {
		t.Errorf("期望找到元素 2，但未找到")
	}

	t.Log("测试用例 3: 更新元素")
	if !singleLinkedList.Uptade(1, 3) {
		t.Errorf("期望更新元素 1 成功，但失败")
	}
	singleLinkedList.Print()
	if singleLinkedList.Find(1) {
		t.Errorf("期望未找到元素 1，但找到")
	}
	if !singleLinkedList.Find(3) {
		t.Errorf("期望找到元素 3，但未找到")
	}

	t.Log("测试用例 4: 删除元素")
	if !singleLinkedList.Delete(2) {
		t.Errorf("期望删除元素 2 成功，但失败")
	}
	singleLinkedList.Print()
	if singleLinkedList.Find(2) {
		t.Errorf("期望未找到元素 2，但找到")
	}

	t.Log("测试用例 5: 删除头节点")
	if !singleLinkedList.Delete(3) {
		t.Errorf("期望删除头节点 3 成功，但失败")
	}
	singleLinkedList.Print()
	if singleLinkedList.Length() != 0 {
		t.Errorf("期望长度 0，但得到 %d", singleLinkedList.Length())
	}

	t.Log("测试用例 6: 在空链表中删除元素")
	if singleLinkedList.Delete(4) {
		t.Errorf("期望删除元素 4 失败，但成功")
	}
	singleLinkedList.Print()

	// 测试用例 7: 合并多个操作
	singleLinkedList.Inseart(5)
	singleLinkedList.Inseart(6)
	singleLinkedList.Print()
	if !singleLinkedList.Find(5) || !singleLinkedList.Find(6) {
		t.Errorf("期望找到元素 5 和 6，但未找到")
	}
	if singleLinkedList.Length() != 2 {
		t.Errorf("期望长度 2，但得到 %d", singleLinkedList.Length())
	}

	// 测试用例 8: 清空链表并检查长度
	singleLinkedList.Delete(5)
	singleLinkedList.Delete(6)
	singleLinkedList.Print()
	if singleLinkedList.Length() != 0 {
		t.Errorf("期望长度 0，但得到 %d", singleLinkedList.Length())
	}
}

func TestSingleLinkedReverse(t *testing.T) {
	// 创建单链表实例
	t.Log("创建链表并打印")
	// 创建一个链表: 1 -> 2 -> 3 -> 4 -> 5
	//head := &SingleNode{data: 1}
	//head.next = &SingleNode{data: 2}
	//head.next.next = &SingleNode{data: 3}
	//head.next.next.next = &SingleNode{data: 4}
	//head.next.next.next.next = &SingleNode{data: 5}
	list := NewSingleLinkedList()
	list.Inseart(1)
	list.Inseart(2)
	list.Inseart(3)
	list.Inseart(4)
	list.Inseart(5)
	list.Inseart(6)
	list.Print()
	p := list.head
	m := MiddleNode(p)
	fmt.Printf("中间节点(奇数)为: %d\n", m.data)

	t.Log("反转链表")
	rev := Reverse(list.head)
	for rev != nil {
		fmt.Printf("%d -> ", rev.data)
		rev = rev.next
	}
	fmt.Println()

	t.Log("添加节点")
	list.Inseart(6)
	pp := list.head
	m = MiddleNode(pp)
	fmt.Printf("中间节点(偶数)为: %d\n", m.data)

}
