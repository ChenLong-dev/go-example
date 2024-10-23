package container

import (
	"container/list"
	"fmt"
)

/*
Element为节点：
type Element struct {
	next, prev *Element
	list *List
	Value interface{}
}

方法列表：
type Element
    func (e *Element) Next() *Element
		返回值：*Element：链表中该节点的下一个节点元素的指针，如果该节点是最后一个节点，则返回nil
		功能说明：获得该节点在链表中的下一个节点元素的指针，如果该节点是最后一个节点，则返回nil

    func (e *Element) Prev() *Element
		返回值：*Element：链表中该节点的上一个节点元素的指针，如果该节点是第一个节点，则返回nil
		功能说明：获得该节点在链表中的上一个节点元素的指针，如果该节点是第一个节点，则返回nil

type List
    func New() *List
		返回值：*List：空链表的指针
		功能说明:创建一个空链表，链表的长度为0，开头和末尾节点都是nil

    func (l *List) Back() *Element
		返回值：*Element，链表中最后一个节点的指针，如果链表长度为0，则为nil

    func (l *List) Front() *Element
		返回值：*Element，链表中第一个节点的指针，如果链表长度为0，则为nil

    func (l *List) Init() *List
		返回值：*List，初始化或者清空后的链表
		功能说明: 初始化或者清空链表，该方法调用后，链表的长度为0

    func (l *List) InsertAfter(v interface{}, mark *Element) *Element
		参数列表：
		- value：要插入的数据的内容
		- mark：链表中的一个节点指针
		返回值：*Element：被插入的节点指针，该节点的Value为数据内容
		功能说明：把数据value插入到mark节点的后面，并返回这个被插入的节点。

    func (l *List) InsertBefore(v interface{}, mark *Element) *Element
		参数列表：
		- value：要插入的数据的内容
		- mark：链表中的一个节点指针
		返回值：*Element：被插入的节点指针，该节点的Value为数据内容
		功能说明：把数据value插入到mark节点的前面，并返回这个被插入的节点。

    func (l *List) Len() int
		返回值：int，链接中节点的个数
		功能说明：获得链接中节点的个数

    func (l *List) MoveAfter(e, mark *Element)
		参数列表：
		- e：链表中的节点
		- mark：链表中的一个节点指针
		功能说明：把节点e移到mark节点的后面

    func (l *List) MoveBefore(e, mark *Element)
		参数列表：
		- e：链表中的节点
		- mark：链表中的一个节点指针
		功能说明：把节点e移到mark节点的前面

    func (l *List) MoveToBack(e *Element)
		参数列表：
		- e：链表中的节点
		功能说明：把节点e移到链表的末尾
    func (l *List) MoveToFront(e *Element)
		参数列表：
		- e：链表中的节点
		功能说明：把节点e移到链表的开头
    func (l *List) PushBack(v interface{}) *Element
		参数列表：
		- value：将被存到链表末尾的任意对象
		返回值：*Element：被存到末尾的节点的指针
		功能说明：把一个对象存到链表末尾，并返回这个节点

    func (l *List) PushBackList(other *List)
		参数列表：
		- ol：将被插入到链表l末尾的链表
		功能说明：把一个链表存到链表末尾

    func (l *List) PushFront(v interface{}) *Element
		参数列表：
		- value：将被存到链表开头的任意对象
		返回值：*Element：被存到开头的节点的指针
		功能说明：把一个对象存到链表开头，并返回这个节点

    func (l *List) PushFrontList(other *List)
		参数列表：
		- ol：将被插入到链表l开头的链表
		功能说明：把一个链表存到链表开头

    func (l *List) Remove(e *Element) interface{}
		参数列表：
		- e：将被删除的节点，该节点必须是属于链表l的
		返回值：interface{}：被删除的节点的内容
		功能说明：删除指定的节点，并返回这个节点的内容
*/

//打印链表
func printList(desc string, l *list.List) {
	ans := make([]interface{}, 0)
	for e := l.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value)
	}
	fmt.Printf("%s: %v\n", desc, ans)
}

func MyList() {
	l := list.New()
	l1 := l.PushFront(1) // [1]
	printList("测试向前添加元素1", l)
	l2 := l.PushFront(2) // [2 1]
	printList("测试向前添加元素2", l)
	l3 := l.PushFront(3) // [3 2 1]
	printList("测试向前添加元素3", l)
	fmt.Println(l1, l2, l3)

	l.MoveToBack(l3)  // [2 1 3]
	printList("测试元素3移到尾部", l)
	l.MoveToFront(l1) // [1 2 3]
	printList("测试元素1移到头部", l)

	l.PushBack(4) // [1 2 3 4]
	printList("测试向后追加元素4", l)

	l.MoveAfter(l1, l2) // [2 1 3 4]
	printList("测试元素1移动到元素2的后面", l)

	l.MoveBefore(l1, l2) // [1 2 3 4]
	printList("测试元素1移动到元素2的前面", l)

	l.InsertAfter(11, l1) // [1 11 2 3 4]
	printList("测试把数据11插入到l1节点[1]的后面，并返回这个被插入的节点", l)

	l.InsertBefore(22, l2) // [1 11 22 2 3 4]
	printList("测试把数据22插入到l2节点[2]的前面，并返回这个被插入的节点", l)

	ll := list.New()
	for i := 50; i < 53; i++ {
		ll.PushBack(i)
	}
	l.PushBackList(ll) // [1 11 22 2 3 4 50 51 52]
	printList("测试在末尾添加list[50 51 52]", l)

	l.PushFrontList(ll) // [50 51 52 1 11 22 2 3 4 50 51 52]
	printList("测试在头部添加list[50 51 52]", l)

	l.Remove(l1) // [50 51 52 11 22 2 3 4 50 51 52]
	printList("测试删除l1节点[1]", l)
}