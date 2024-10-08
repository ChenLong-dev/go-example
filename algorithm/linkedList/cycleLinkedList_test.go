package linkedList

import "testing"

// 测试环形链表
func TestCycleLinkedList(t *testing.T) {
	list := new(cycleLinkedList).NewCycleLinkedList()

	t.Log("添加节点")
	list.Append(1)
	list.Print()
	if list.Size() != 1 {
		t.Errorf("期望链表长度为1，但实际为%d", list.Size())
	}

	list.Append(2)
	list.Print()
	if list.Size() != 2 {
		t.Errorf("期望链表长度为2，但实际为%d", list.Size())
	}

	list.Append(3)
	list.Print()
	if list.Size() != 3 {
		t.Errorf("期望链表长度为3，但实际为%d", list.Size())
	}

	// 快乐路径测试：查找节点
	t.Log("查找节点")
	if data := list.Find(2); data == nil || data != 2 {
		t.Errorf("期望找到节点2，但未找到")
	}

	// 边界条件测试：查找不存在的节点
	t.Log("查找不存在的节点")
	if data := list.Find(4); data != nil {
		t.Errorf("期望未找到节点4，但却找到了")
	}

	// 快乐路径测试：更新节点
	t.Log("更新节点")
	if success := list.Update(2, 4); !success {
		t.Errorf("期望更新节点2为4，更新失败")
	}

	if data := list.Find(4); data == nil || data != 4 {
		t.Errorf("期望找到更新后的节点4，但未找到")
	}
	list.Print()

	// 边界条件测试：更新不存在的节点
	t.Log("更新不存在的节点")
	if success := list.Update(5, 6); success {
		t.Errorf("期望更新不存在的节点5，更新却成功")
	}
	list.Print()

	// 快乐路径测试：删除节点
	t.Log("删除节点")
	if success := list.Delete(3); !success {
		t.Errorf("期望删除节点3，删除失败")
	}

	if data := list.Find(3); data != nil {
		t.Errorf("期望未找到节点3，但却找到了")
	}
	list.Print()

	// 边界条件测试：删除不存在的节点
	t.Log("边界条件测试：删除不存在的节点")
	if success := list.Delete(5); success {
		t.Errorf("期望删除不存在的节点5，删除却成功")
	}
	list.Print()

	// 边界条件测试：删除最后一个节点
	t.Log("边界条件测试：删除最后一个节点")
	list.Delete(1)
	list.Print()
	list.Delete(4)
	list.Print()

	if list.Size() != 0 {
		t.Errorf("期望链表长度为0，但实际为%d", list.Size())
	}

	//// 边界条件测试：从空链表中删除节点
	//if success := list.Delete(1); success {
	//	t.Errorf("期望从空链表删除失败，删除却成功")
	//}
}
