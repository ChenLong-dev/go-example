package tree

import (
	"fmt"
	"strings"
	"testing"
)

// TestItemBinarySearchTree 单元测试
func TestItemBinarySearchTree(t *testing.T) {
	bst := &ItemBinarySearchTree{}

	// 测试插入节点
	t.Log("测试插入节点")
	bst.Insert(10, "十")
	bst.Insert(5, "五")
	bst.Insert(15, "十五")
	bst.Insert(75, "七十五")
	bst.Insert(20, "二十")
	bst.Insert(30, "三十")
	bst.Insert(40, "四十")
	bst.Insert(60, "六十")
	bst.Insert(80, "八十")
	bst.Print()

	// 测试最小值节点
	t.Log("测试最小值节点")
	minNode := bst.Min()
	if minNode.Key != 5 {
		t.Errorf("期望最小值为 5, 但是得到 %d", minNode.Key)
	}

	// 测试最大值节点
	t.Log("测试最大值节点")
	maxNode := bst.Max()
	if maxNode.Key != 80 {
		t.Errorf("期望最大值为 15, 但是得到 %d", maxNode.Key)
	}

	// 测试查找节点--根节点
	t.Log("测试查找节点--根节点-10")
	searchNode := bst.Search(10)
	if searchNode.Key != 10 {
		t.Errorf("期望查找到的节点为 10, 但是得到 %d", searchNode.Key)
	}

	t.Log("测试查找节点--中间节点-20")
	searchNode = bst.Search(20)
	if searchNode.Key != 20 {
		t.Errorf("期望查找到的节点为 20, 但是得到 %d", searchNode.Key)
	}

	t.Log("测试查找节点--叶子节点-60")
	searchNode = bst.Search(60)
	if searchNode.Key != 60 {
		t.Errorf("期望查找到的节点为 60, 但是得到 %d", searchNode.Key)
	}

	// 测试更新节点--根节点
	t.Log("测试更新节点--根节点-10")
	if !bst.Update(10, "更新的十") {
		t.Error("更新失败")
	}
	bst.Print()
	updatedNode := bst.Search(10)
	if updatedNode.Value != "更新的十" {
		t.Errorf("期望更新后的值为 '更新的十', 但是得到 %v", updatedNode.Value)
	}

	// 测试更新节点--中间节点
	t.Log("测试更新节点--中间节点-20")
	if !bst.Update(20, "更新的二十") {
		t.Error("更新失败")
	}
	bst.Print()
	updatedNode = bst.Search(20)
	if updatedNode.Value != "更新的二十" {
		t.Errorf("期望更新后的值为 '更新的二十', 但是得到 %v", updatedNode.Value)
	}

	// 测试更新节点--叶子节点
	t.Log("测试更新节点--叶子节点-80")
	if !bst.Update(80, "更新的八十") {
		t.Error("更新失败")
	}
	bst.Print()
	updatedNode = bst.Search(80)
	if updatedNode.Value != "更新的八十" {
		t.Errorf("期望更新后的值为 '更新的八十', 但是得到 %v", updatedNode.Value)
	}

	// 测试删除节点--根节点
	t.Log("测试删除节点--根节点-10")
	bst.Delete(10)
	bst.Print()
	deletedNode := bst.Search(10)
	if deletedNode != nil {
		t.Errorf("期望删除节点后找不到 10, 但是找到了 %v", deletedNode)
	}

	// 测试删除节点--中间节点
	t.Log("测试删除节点--根节点-75")
	bst.Delete(75)
	bst.Print()
	deletedNode = bst.Search(75)
	if deletedNode != nil {
		t.Errorf("期望删除节点后找不到 75, 但是找到了 %v", deletedNode)
	}

	// 测试删除节点--叶子节点
	t.Log("测试删除节点--根节点-60")
	bst.Delete(60)
	bst.Print()
	deletedNode = bst.Search(60)
	if deletedNode != nil {
		t.Errorf("期望删除节点后找不到 60, 但是找到了 %v", deletedNode)
	}

	// 边界情况测试：删除不存在的节点
	t.Log("测试删除不存在的节点")
	bst.Delete(999) // 不应崩溃且无需返回值
}

// 测试空树的情况
func TestEmptyTree(t *testing.T) {
	bst := &ItemBinarySearchTree{}

	if bst.Min() != nil {
		t.Error("空树的最小值应为 nil")
	}
	if bst.Max() != nil {
		t.Error("空树的最大值应为 nil")
	}
	if bst.Search(10) != nil {
		t.Error("空树中查找任何键应为 nil")
	}
}

func TestOrderTraversal(t *testing.T) {
	bst := &ItemBinarySearchTree{}
	// 初始化
	bst.Insert(5, "A")
	bst.Insert(4, "B")
	bst.Insert(8, "C")
	bst.Insert(3, "D")
	bst.Insert(6, "E")
	bst.Insert(9, "F")
	bst.Insert(1, "G")
	bst.Insert(2, "H")
	bst.Insert(7, "I")

	bst.Print()

 	// 前序遍历
	t.Log("前序遍历-递归")
	orderTraversal := []string{}
	bst.PreOrderTraversal(func(itemNode *Node) {
		fmt.Printf("[%d]%s\t", itemNode.Key, itemNode.Value)
		orderTraversal = append(orderTraversal, itemNode.Value.(string))
	})
	fmt.Printf("\n%s\n", orderTraversal)
	t.Log("前序遍历-非递归（循环迭代）")
	orderTraversalByStack := []string{}
	bst.PreOrderTraversalByStack(func(itemNode *Node) {
		fmt.Printf("[%d]%s\t", itemNode.Key, itemNode.Value)
		orderTraversalByStack = append(orderTraversalByStack, itemNode.Value.(string))
	})
	fmt.Printf("\n%s\n", orderTraversalByStack)
	if strings.Join(orderTraversal, "") != strings.Join(orderTraversalByStack, "") {
		t.Error("前序遍历结果不一致")
	}

	t.Log("中序遍历-递归")
	orderTraversal = []string{}
	bst.InOrderTraversal(func(itemNode *Node) {
		fmt.Printf("[%d]%s\t", itemNode.Key, itemNode.Value)
		orderTraversal = append(orderTraversal, itemNode.Value.(string))
	})
	fmt.Printf("\n%s\n", orderTraversal)
	t.Log("中序遍历-非递归（循环迭代）")
	orderTraversalByStack = []string{}
	bst.InOrderTraversalByStack(func(itemNode *Node) {
		fmt.Printf("[%d]%s\t", itemNode.Key, itemNode.Value)
		orderTraversalByStack = append(orderTraversalByStack, itemNode.Value.(string))
	})
	fmt.Printf("\n%s\n", orderTraversalByStack)
	if strings.Join(orderTraversal, "") != strings.Join(orderTraversalByStack, "") {
		t.Error("中序遍历结果不一致")
	}

	t.Log("后序遍历-递归")
	orderTraversal = []string{}
	bst.PostOrderTraversal(func(itemNode *Node) {
		fmt.Printf("[%d]%s\t", itemNode.Key, itemNode.Value)
		orderTraversal = append(orderTraversal, itemNode.Value.(string))
	})
	fmt.Printf("\n%s\n", orderTraversal)
	t.Log("后序遍历-非递归（循环迭代）")
	orderTraversalByStack = []string{}
	bst.PostOrderTraversalByStack(func(itemNode *Node) {
		fmt.Printf("[%d]%s\t", itemNode.Key, itemNode.Value)
		orderTraversalByStack = append(orderTraversalByStack, itemNode.Value.(string))
	})
	fmt.Printf("\n%s\n", orderTraversalByStack)
	if strings.Join(orderTraversal, "") != strings.Join(orderTraversalByStack, "") {
		t.Error("后序遍历结果不一致")
	}





}
