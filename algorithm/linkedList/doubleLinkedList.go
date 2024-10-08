package linkedList

import "fmt"

// 定义双向链表节点
type doubleNode struct {
	data interface{}
	prev *doubleNode
	next *doubleNode
}

// DoubleLinkedList 定义双向链表
type DoubleLinkedList struct {
	head *doubleNode
	tail *doubleNode
	size int
}

// NewDoubleLinkedList 创建一个双向链表
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{head: nil, tail: nil, size: 0}
}

// 向链表尾部添加一个节点
func (l *DoubleLinkedList) append(data interface{}) {
	newNode := &doubleNode{data: data, prev: nil, next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.size++
}

// 向链表头部添加一个节点
func (l *DoubleLinkedList) prepend(data interface{}) {
	newNode := &doubleNode{data: data, prev: nil, next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

// 查找链表中是否存在指定的数据
func (l *DoubleLinkedList) find(data interface{}) *doubleNode {
	current := l.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil
}

// 删除链表中指定的数据
func (l *DoubleLinkedList) delete(data interface{}) bool {
	targetNode := l.find(data)
	if targetNode == nil {
		return false
	}
	if targetNode.prev == nil { // 头节点
		l.head = targetNode.next
	} else { // 中间节点
		targetNode.prev.next = targetNode.next
	}

	if targetNode.next == nil { // 尾节点
		l.tail = targetNode.prev
	} else { // 中间节点
		targetNode.next.prev = targetNode.prev
	}
	l.size--
	return true
}

// 更新链表中指定的数据
func (l *DoubleLinkedList) update(oldData, newData interface{}) bool {
	targetNode := l.find(oldData)
	if targetNode != nil {
		targetNode.data = newData
		return true
	}

	return false
}

// 链表转数组
func (l *DoubleLinkedList) asArray() []interface{} {
	result := make([]interface{}, 0)
	current := l.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func (l *DoubleLinkedList) asArrayReverse() []interface{} {
	result := make([]interface{}, 0)
	current := l.tail
	for current != nil {
		result = append(result, current.data)
		current = current.prev
	}
	return result
}

// 链表长度
func (l *DoubleLinkedList) length() int {
	return l.size
}

// Append 向链表尾部添加一个节点
func (l *DoubleLinkedList) Append(data interface{}) {
	l.append(data)
}

// Prepend 向链表头部添加一个节点
func (l *DoubleLinkedList) Prepend(data interface{}) {
	l.prepend(data)
}

// Find 查找链表中是否存在指定的数据
func (l *DoubleLinkedList) Find(data interface{}) *doubleNode {
	return l.find(data)
}

// Delete 删除链表中指定的数据
func (l *DoubleLinkedList) Delete(data interface{}) bool {
	return l.delete(data)
}

// Update 更新链表中指定的数据
func (l *DoubleLinkedList) Update(oldData, newData interface{}) bool {
	return l.update(oldData, newData)
}

// Print 打印链表
func (l *DoubleLinkedList) Print() {
	result := l.asArray()
	fmt.Printf("-> size: %d， length: %d\n", l.size, len(result))
	for i, value := range result {
		fmt.Printf("[%d-%d] %v\n", l.size, i+1, value)
	}
}

// PrintReverse 打印链表（反序）
func (l *DoubleLinkedList) PrintReverse() {
	result := l.asArrayReverse()
	fmt.Printf("<- size: %d， length: %d\n", l.size, len(result))
	for i, value := range result {
		fmt.Printf("[%d-%d] %v\n", l.size, i+1, value)
	}
}