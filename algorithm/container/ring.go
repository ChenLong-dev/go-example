package container

import (
	"container/ring"
	"fmt"
)

/*
方法列表：
type Ring
	func New(n int) *Ring
		参数列表：
		- n：环形双向链表的节点的个数
		返回值：*Ring：空链表的指针
		功能说明:创建一个有n个节点的环形双向链表

	func (r *Ring) Do(f func(interface{}))
		参数列表：
		- `f`：一个回调函数，该函数的参数为环形双向链表中的节点的`Value`字段值
		功能说明：正向遍历环形双向链表，并对每个链表中的元素执行回调函数`f`，如果这个回调函数`f`会修改链表`r`，那这个回调函数的行为将不可确定

	func (r *Ring) Len() int
		返回值：
		- `int`：环形链表中元素的个数
		功能说明:遍历环形双向链表来统计链表中的元素的个数

	func (r *Ring) Link(s *Ring) *Ring
		参数列表：
		- `s`：环形双向链表
		返回值：`*Ring`：`s`和`r`相连前`r.Next()`的值，也是相连后`s.Next()`的值
		功能说明:
		- 把一个环形双向链表`s`与环形双向链表`r`相链接，链接后`r.Next()`为s，并返回相连前时`r.Next()`的值。`r`不能为空。
		- 如果`s`和`r`不是同一个环形链表，则相连后，只产生一个环形链表，并返回相连前时`r.Next()`的值，也是相连后`s.Prev()`的值。
		- 如果`s`和`r`是同一个环形链表，但`s != r`时，相连后，产生两个环形链表，其中一个是由`r`和`s`之间的节点构成（不包括`r`和`s`），返回值为相连前时`r.Next()`的值，即`r`和`s`之间的节点（不包括`r`和`s`）构成的环形链表的表头节点。
		- 如果`s`和`r`是同一个环形链表，且`s == r`时，相连后，产生两个环形链表，其中一个是由`r`指向的节点构成的长度为1的环形链表，其他节点构成另一个环形链表，返回值为相连前时`r.Next()`的值，即其他节点构成的环形链表的表头节点。

	func (r *Ring) Move(n int) *Ring
		参数列表：
		- `n`：指针`r`在双向链表上移动的位置的个数。n>0时，为正向移动；反之为反向移动。
		返回值：*Ring：移动结束后，指针`r`指向的节点
		功能说明：指向节点`r`的指针，正向或者逆向移动`n % r.Len()`个节点，并返回这个指针移动后指向的节点。但是`r.Move(n)`不对改变r的值，`r`不能为空

	func (r *Ring) Next() *Ring
		返回值： `*Ring`：指向下一个节点的指针
		功能说明：获得指向下一个节点的指针

	func (r *Ring) Prev() *Ring
		返回值：`*Ring`：指向上一个节点的指针
		功能说明：获得指向上一个节点的指针

	func (r *Ring) Unlink(n int) *Ring
		参数列表：
		- `n`：要被移除的节点的数个
		功能说明：从节点`r`的下一个节点（包含该节点）开始移除`n % r.Len()`个节点。如果`n % r.Len() == 0`，则链表不会有改变。`r`不能为空。
*/


func printRing(desc string, r *ring.Ring) {
	ans := make([]interface{}, 0)
	r.Do(func(v interface{}) {
		ans = append(ans, v)
	})
	fmt.Printf("%s: %v\n", desc, ans)
}

func makeN(n int, begin int) *ring.Ring {
	r := ring.New(n)
	for i := begin; i < n+begin; i++ {
		r.Value = i
		r = r.Next()
	}
	return r
}

// Link两个不同的环形链表
func linkDiffRing() {
	r1 := makeN(5, 0) // [0 1 2 3 4]
	printRing("创建一个容量为5的环形链表，从元素0开始", r1)

	r2 := makeN(5, 10) // [10 11 12 13 14]
	printRing("创建一个容量为5的环形链表，从元素10开始", r1)

	r1.Link(r2)
	fmt.Println("r1.Value: ", r1.Value) // 输出：0
	fmt.Println("r2.Value: ", r2.Value) // 输出：10

	printRing("Link两个不同的环形链表，把一个环形双向链表`s`与环形双向链表`r`相链接，r1", r1) // [0 10 11 12 13 14 1 2 3 4]
	printRing("Link两个不同的环形链表，把一个环形双向链表`s`与环形双向链表`r`相链接, r2", r2) // [10 11 12 13 14 1 2 3 4 0]
}

// 如果两个链表是同一个链表，但是两个指针不是指这个链表中的节点时，即s != r时
func linkSameRing() {
	r1 := makeN(10, 0)
	printRing("创建一个容量为10的环形链表，从元素0开始", r1) // [0 1 2 3 4 5 6 7 8 9]

	r2 := r1.Move(5)
	fmt.Println(r1.Value, r2.Value) //0 5
	printRing("正向或者逆向移动`n % r.Len()`个节点, r1", r1) // [0 1 2 3 4 5 6 7 8 9]
	printRing("正向或者逆向移动`n % r.Len()`个节点, r2", r2) // [5 6 7 8 9 0 1 2 3 4]

	r3 := r1.Link(r2)

	printRing("两个链表是同一个链表，但是两个指针不是指这个链表中的节点时，即s != r时，相连后，产生两个环形链表，r1", r1) // [0 5 6 7 8 9]
	printRing("两个链表是同一个链表，但是两个指针不是指这个链表中的节点时，即s != r时，相连后，产生两个环形链表，r3", r3) // [1 2 3 4]
}

// 如果两个链表是同一个链表，且两个指针指向这个链表中的同一个节点时，即s == r时
func linkSameElement() {
	r1 := makeN(10, 0)
	printRing("创建一个容量为10的环形链表，从元素0开始", r1) // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println("r1.Value:", r1.Value) // 输出：0

	r2 := r1.Link(r1)

	printRing("两个链表是同一个链表，且两个指针指向这个链表中的同一个节点时，即s == r时, r1", r1) // [0]
	printRing("两个链表是同一个链表，且两个指针指向这个链表中的同一个节点时，即s == r时, r2", r2) // [1 2 3 4 5 6 7 8 9]
}

func MyRing() {
	linkDiffRing()
	linkSameRing()
	linkSameElement()
}