package linkedList

import "fmt"

/*
特性：
链表的尾节点指向头节点或其他节点，从而形成一个环状结构。将单向链表的首尾相连，使得链表中的任意节点都能通过链式关系找到其他任意节点，包括头节点
2、必须至少有一个头节点
3、在环形链表中插入或删除节点时，需要特别处理，不能破坏原有的环形结构。遍历环形链表时也需要使用特殊的方法，否则可能会陷入死循环。

应用：
环形链表主要用于解决一些具有特殊需求的问题，例如约瑟夫问题。
约瑟夫问题是一个经典的数学问题，描述的是编号为1,2,3,...n的n个人围坐一圈，约定编号为k（1<=k<=n）的人从1开始报数，数到m的那个人出列，它的下一位又从1开始报数，数到m的那个人又出列，依次类推，直到所有人出列为止。这个问题可以通过环形链表来方便地解决。
*/

// cycleNode 环形链表节点
type cycleNode struct {
	data interface{}
	next *cycleNode
}

// cycleLinkedList 环形链表
type cycleLinkedList struct {
	head *cycleNode
	tail *cycleNode
	size int
}

// NewCycleLinkedList 创建一个新的环形链表
func (l *cycleLinkedList) NewCycleLinkedList() *cycleLinkedList {
	return &cycleLinkedList{head: nil, tail: nil, size: 0}
}

// append 向链表尾部添加一个节点
func (l *cycleLinkedList) append(data interface{}) {
	newNode := &cycleNode{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.head.next = l.tail // 首尾相连
		l.tail.next = l.head
	} else {
		l.tail.next = newNode
		newNode.next = l.head
		l.tail = newNode
	}
	l.size++
	return
}

// find 查找一个节点
func (l *cycleLinkedList) find(data interface{}) *cycleNode {
	if l.head == nil {
		return nil
	}
	current := l.head
	for {
		if current.data == data {
			return current
		}
		current = current.next
		if current == l.head {
			break
		}
	}
	return nil
}

// update 更新一个节点的数据
func (l *cycleLinkedList) update(oldData interface{}, newData interface{}) bool {
	current := l.find(oldData)
	if current == nil {
		return false
	}
	current.data = newData
	return true
}

// delete 删除一个节点
func (l *cycleLinkedList) delete(data interface{}) *cycleNode {
	if l.head == nil {
		return nil
	}
	// 特殊情况处理：删除头节点
	if l.head.data == data {
		fmt.Println("xxxxx2 ", data)
		if l.head.next == l.head { // 只有一个节点的情况
			l.head = nil
			l.tail = nil
			fmt.Println("xxxxx3:", data)
		} else { // 删除头节点（有两个以上节点的情况）

		}
		l.size--
		return l.head
	}

	current := l.head
	for {
		if current.next.data == data {
			fmt.Println("xxxxx4:", data)
			current.next = current.next.next
			l.size--
			return current.next
		}
		current = current.next
		if current == l.head {
			break
		}
	}
	return nil
}

// asArray 链表转数组
func (l *cycleLinkedList) asArray() []interface{} {
	result := make([]interface{}, 0)
	if l.head == nil {
		return result
	}
	if l.head == l.tail {
		result = append(result, l.head.data)
		fmt.Println("xxxxx1", result)
		return result
	}

	current := l.head
	for {
		result = append(result, current.data)
		current = current.next
		if current == l.head {
			break
		}
	}
	return result
}

// Append 向链表尾部添加一个节点
func (l *cycleLinkedList) Append(data interface{}) {
	l.append(data)
}

// Find 查找一个节点
func (l *cycleLinkedList) Find(data interface{}) interface{} {
	node := l.find(data)
	if node == nil {
		return nil
	}
	return node.data
}

// Update 更新一个节点的数据
func (l *cycleLinkedList) Update(oldData interface{}, newData interface{}) bool {
	return l.update(oldData, newData)
}

// Delete 删除一个节点
func (l *cycleLinkedList) Delete(data interface{}) bool {
	node := l.delete(data)
	if node == nil {
		return false
	}
	return true
}

// Size 链表长度
func (l *cycleLinkedList) Size() int {
	return l.size
}

// Print 打印链表
func (l *cycleLinkedList) Print() {
	result := l.asArray()
	fmt.Printf("-> size:%d, lenght:%d\n", l.size, len(result))
	for i, v := range result {
		fmt.Printf("[%d-%d] data:%v\n", l.size, i+1, v)
	}
}
