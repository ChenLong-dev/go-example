package linkedList

import (
	"testing"
)

// 测试双向链表的单元测试
func TestDoubleLinkedList(t *testing.T) {
	// 创建一个双向链表
	list := NewDoubleLinkedList()

	// 测试添加节点（快乐路径）
	t.Log("测试添加节点（快乐路径）")
	list.Append(1)
	list.Print()
	if list.length() != 1 {
		t.Errorf("期待链表长度为1，实际为%d", list.length())
	}

	// 添加多个节点
	t.Log("添加多个节点")
	list.Append(2)
	list.Append(3)
	list.Print()
	list.PrintReverse()
	if list.length() != 3 {
		t.Errorf("期待链表长度为3，实际为%d", list.length())
	}

	// 测试查找节点（快乐路径）
	t.Log("测试查找节点（快乐路径）")
	node := list.Find(2)
	if node == nil || node.data != 2 {
		t.Error("未能找到数据2")
	}

	// 测试更新节点
	t.Log("测试更新节点")
	if !list.Update(2, 4) {
		t.Error("更新数据2为4失败")
	}
	list.Print()
	if list.Find(2) != nil {
		t.Error("数据2仍然存在，未被正确更新")
	}
	if list.Find(4) == nil {
		t.Error("未能找到更新后的数据4")
	}

	// 测试删除节点
	t.Log("测试删除节点")
	if !list.Delete(3) {
		t.Error("删除数据3失败")
	}
	list.Print()
	if list.Find(3) != nil {
		t.Error("数据3仍然存在，未被正确删除")
	}

	// 测试边界情况 - 删除链表中唯一节点
	t.Log("测试边界情况 - 删除链表中唯一节点")
	list.Delete(1)
	list.Print()
	if list.length() != 1 {
		t.Errorf("期待链表长度为1，实际为%d", list.length())
	}
	if !list.Delete(4) {
		t.Error("删除数据4失败")
	}
	list.Print()
	if list.length() != 0 {
		t.Errorf("期待链表长度为0，实际为%d", list.length())
	}

	// 边界情况 - 查找不存在的节点
	t.Log("边界情况 - 查找不存在的节点")
	if list.Find(5) != nil {
		t.Error("查找不存在的数据5却得到了结果")
	}

	// 边界情况 - 更新不存在的节点
	t.Log("边界情况 - 更新不存在的节点")
	if list.Update(5, 6) {
		t.Error("更新不存在的数据5却成功了")
	}

	// 边界情况 - 从空链表删除
	t.Log("边界情况 - 从空链表删除")
	if list.Delete(1) {
		t.Error("从空链表中删除数据1却成功了")
	}
}
