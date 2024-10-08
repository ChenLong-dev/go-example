package hashTable

import "fmt"

// AraySize 定义数组大小
const AraySize = 5

// HashTable 定义哈希表结构
type bucketNode struct {
	key  string
	next *bucketNode
}

// bucket 定义哈希表桶结构
type bucket struct {
	head *bucketNode
	size int
}

// HashTable 定义哈希表结构
type HashTable struct {
	buckets []*bucket
	size    int
}

// 初始化哈希表
func (b *bucket) insert(key string) {
	// 获取当前桶的头节点的下一个节点（即当前的第一个有效节点）。为了在插入新节点时保持原有链表的结构。
	next := b.head.next
	// 创建新的桶节点，插入到当前桶的头节点的下一个节点之前。
	node := &bucketNode{key: key, next: next}
	// 将新节点插入桶中
	b.head.next = node
	b.size++
}

// 查找哈希表中是否存在指定键值
func (b *bucket) search(key string) *bucketNode {
	current := b.head.next
	for current != nil {
		if current.key == key {
			return current
		}
	}
	return nil
}

// 删除哈希表中指定键值
func (b *bucket) delete(key string) *bucketNode {
	current := b.head.next
	for current.next != nil {
		if current.key == key {
			current.next = current.next.next
			b.size--
			return current
		}
		current = current.next
	}
	return nil
}

// Print 打印哈希表中所有键值
func (b *bucket) Print() {
	current := b.head.next
	for current != nil {
		fmt.Printf("%v\n", current.key)
		current = current.next
	}
}

// hash 哈希函数
func hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % AraySize
}

// InitHashTable 初始化哈希表
func InitHashTable() *HashTable {
	table := &HashTable{}
	for i := range table.buckets {
		table.buckets[i] = &bucket{
			head: &bucketNode{next: nil},
			size: 0,
		}
	}
	return table
}

// Insert 插入键值对
func (t *HashTable) Insert(key string) {
	t.buckets[hash(key)].insert(key)
	t.size++
}

// Search 查找键值对
func (t *HashTable) Search(key string) *bucketNode {
	return t.buckets[hash(key)].search(key)
}

// Delete 删除键值对
func (t *HashTable) Delete(key string) *bucketNode {
	bn := t.buckets[hash(key)].delete(key)
	if bn != nil {
		t.size--
		return bn
	}
	return nil
}

// Print 打印哈希表
func (t *HashTable) Print() {
	for i := range t.buckets {
		fmt.Printf("bucket %v:\n", i)
		t.buckets[i].Print()
	}
}
