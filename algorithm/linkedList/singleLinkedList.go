package linkedList

import "fmt"

/*
特性：
1. 单链表的每个节点只有一个指向下一个节点的指针。
2. 链表的头节点指向第一个节点，尾节点指向最后一个节点。
3. 链表的长度是指链表中节点的个数。
4. 链表的插入、删除操作都可以在O(1)时间内完成。
5. 链表的遍历操作可以在O(n)时间内完成。
6. 链表的节点可以存储任意类型的数据。
应用：
1. 栈、队列、双向链表、循环链表等数据结构的实现。
2. 图的邻接表表示、拓扑排序、最小生成树等算法的实现。
时间复杂度：
1. 链表的插入、删除操作：O(1)
2. 链表的遍历操作：O(n)
空间复杂度：
1. 链表的存储空间：O(n)
2. 链表的节点存储空间：O(1)
总结：
1. 单链表是一种简单的数据结构，适用于实现栈、队列、链表、图等数据结构。
2. 单链表的插入、删除操作可以在O(1)时间内完成，而遍历操作可以在O(n)时间内完成。
*/

// 定义链表节点
type SingleNode struct {
	data interface{} // 节点存储的数据
	next *SingleNode // 指向下一个节点的指针
}

// SingleLinkedList 定义单向链表
type SingleLinkedList struct {
	head *SingleNode // 单向链表头指针
	tail *SingleNode // 单向链表尾指针
	size int         // 链表长度
}

// NewSingleLinkedList 新建单向链表
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{head: nil, tail: nil, size: 0}
}

// 插入节点到链表尾部
func (l *SingleLinkedList) append(data interface{}) {
	newNode := &SingleNode{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

// 查找参数data是否在链表中
func (l *SingleLinkedList) find(data interface{}) *SingleNode {
	current := l.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil
}

// 更新节点数据
func (l *SingleLinkedList) update(oldData interface{}, newData interface{}) bool {
	targetNode := l.find(oldData)
	if targetNode != nil {
		targetNode.data = newData
		return true
	}
	return false
}

// 删除节点数据
func (l *SingleLinkedList) delete(data interface{}) bool {
	if l.head == nil {
		return false
	}
	// 特殊情况：删除头节点
	if l.head.data == data {
		l.head = l.head.next
		l.size--
		if l.head == nil { // 如果链表的头指针为空（链表已经为空链表），则更新尾节点为空
			l.tail = nil
		}
		return true
	}

	// 普通情况：删除中间节点（从头节点的下一个节点开始遍历）
	current := l.head
	for current.next != nil { // 指向头节点的下一个节点，头节点已经处理过了
		if current.next.data == data { // 找到要删除的节点，将其前一个节点的next指针指向当前节点的下一个节点下一个节点
			if current.next == l.tail {
				l.tail = current
			}
			current.next = current.next.next
			l.size--
			return true
		}
	}
	return false
}

// 链表长度
func (l *SingleLinkedList) length() int {
	return l.size
}

// 链表转数组
func (l *SingleLinkedList) asArray() []interface{} {
	result := make([]interface{}, 0)
	current := l.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

// Inseart 插入链表方法
func (l *SingleLinkedList) Inseart(data interface{}) {
	l.append(data)
}

// Find 查找链表方法
func (l *SingleLinkedList) Find(data interface{}) bool {
	return l.find(data) != nil
}

// Uptade 更新链表方法
func (l *SingleLinkedList) Uptade(oldData interface{}, newData interface{}) bool {
	return l.update(oldData, newData)
}

// Delete 删除链表方法
func (l *SingleLinkedList) Delete(data interface{}) bool {
	return l.delete(data)
}

// Length 获取链表长度方法
func (l *SingleLinkedList) Length() int {
	return l.length()
}

// Print 打印链表方法
func (l *SingleLinkedList) Print() {
	fmt.Printf("-> size:%d\n", l.size)
	for i, v := range l.asArray() {
		fmt.Printf("[%d-%d] data:%+v\n", l.size, i+1, v)
	}
}

func Reverse(head *SingleNode) *SingleNode {
	var prev *SingleNode = nil // 初始化前一个节点为 nil
	current := head            // 当前节点指向链表的头
	for current != nil {       // 遍历链表
		next := current.next // 保存当前节点的下一个节点
		current.next = prev  // 反转当前节点的指针
		prev = current       // 将前一个节点更新为当前节点
		current = next       // 移动到下一个节点
	}
	p := prev
	for p != nil {
		fmt.Println("-", p.data)
		p = p.next
	}
	return prev // 返回反转后的链表头
}

func MiddleNode(head *SingleNode) *SingleNode {
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
