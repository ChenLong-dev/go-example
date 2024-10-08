package tree

import (
	"cgroup/algorithm/stack"
	"fmt"
)

/*
实现二叉搜索树的接口
二叉搜索树的结构：
	- 左子树的所有节点的值都小于根节点的值
	- 右子树的所有节点的值都大于根节点的值
	- 左右子树也分别为二叉搜索树
术语：
	- 节点的度： 一个节点含有的子树的个数称为节点的度
	- 树的度 ： 一棵树中，最大的节点的度称为树的度
	- 叶节点或终端节点：度为零的节点
	- 非终端节点或分支节点： 度不为零的节点
	- 父亲节点或父节点： 若一个节点含有子节点，则这个节点称为其子节点的父节点
	- 孩子节点或子节点：一个节点含有的子树的根节点称为该节点的子节点
	- 兄弟节点： 具有相同父节点的节点互称为兄弟节点
	- 节点的层次： 从根节点开始，根为第一层，根的子节点为第二层，以此类推
	- 深度：对于任意节点n， n的深度为从根到n的唯一路径长，根的深度为0
	- 高度： 对于任意节点n，n的高度为从n到一片树叶的最长路径厂，所有树叶的高度为0
	- 堂兄弟节点：父节点在同一层的节点互称堂兄节点
	- 子孙： 以某节点为根的子树中任一节点都称为该节点的子孙
	- 森林：由m（m>0）颗互不相交的树的集合称为森林

树的分类：
无序树 ： 树中任意节点的子节点之间没有顺序关系，即任意节点的左、右子树可以是任意颗树
有序树 ： 树中任意节点的子节点之间有顺序关系，即任意节点的左子树的值小于或等于根节点的值，右子树的值大于或等于根节点的值
	- 二叉树： 树中每个节点最多含有两个子树的树称为二叉树
		- 完全二叉树： 所有层次都被完全填满，除了最底层，其他层的节点都达到最大个数，并且最底层的叶节点都集中在该层的左边
			- 满二叉树： 所有叶节点都在最后一层，并且除了最底层外，其他层的节点都达到最大个数
		- 平衡二叉树： 二叉树中任意节点的左右子树的高度差的绝对值不超过1，并且左右子树都是一颗平衡二叉树
		- 排序二叉树： 二叉树中任意节点的左子树的值小于或等于根节点的值，右子树的值大于或等于根节点的值，并且左右子树都是一颗排序二叉树
	- 霍夫曼树： 带权路径最短的二叉树，即带权路径长度最小的二叉树，权值较大的节点离根节点更近
	- B树： 一种对外存的平衡查找树，能够在O(log n)时间内查找一个节点，并在O(log n)时间内插入一个节点

排序二叉树：
	- 若任意节点的左子树不为空，则左子树上所有的节点的值都小于它的根节点的值
	- 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值
	- 任意节点的左、右子树也分别为二叉查找树
	- 没有键值相等的节点

*/

// Item 接口
type Item interface {
}

// Node 节点
type Node struct {
	Key   int
	Value Item
	left  *Node
	right *Node
}

// Noder 接口
type Noder interface {
	Insert(key int, value Item)                       // 插入节点
	Min() *Node                                       // 最小值节点
	Max() *Node                                       // 最大值节点
	Search(key int) *Node                             // 查找节点
	Update(key int, value Item)                       // 更新节点
	InOrderTraversal(f func(itemNode *Node))          // 中序遍历
	InOrderTraversalByStack(f func(itemNode *Node))   // 中序遍历（非递归）
	PreOrderTraversal(f func(itemNode *Node))         // 先序遍历
	PreOrderTraversalByStack(f func(itemNode *Node))  // 先序遍历（非递归）
	PostOrderTraversal(f func(itemNode *Node))        // 后序遍历
	PostOrderTraversalByStack(f func(itemNode *Node)) // 后序遍历（非递归）
	Delete(key int) *Node                             // 删除节点
	Print()                                           // 打印树
}

// ItemBinarySearchTree 二叉搜索树
type ItemBinarySearchTree struct {
	Root *Node
}

// NewItemBinarySearchTree 新建二叉搜索树
func (bst *ItemBinarySearchTree) insertNode(node, newNode *Node) {
	if node.Key > newNode.Key {
		if node.left == nil {
			node.left = newNode
		} else {
			bst.insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			bst.insertNode(node.right, newNode)
		}
	}
}

// Insert 插入节点
func (bst *ItemBinarySearchTree) Insert(key int, value Item) {
	newNode := &Node{Key: key, Value: value, left: nil, right: nil}
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		bst.insertNode(bst.Root, newNode)
	}
}

// Min 最小值节点
func (bst *ItemBinarySearchTree) Min() *Node {
	current := bst.Root
	if current == nil {
		return nil
	}
	for {
		if current.left == nil {
			return current
		} else {
			current = current.left
		}
	}
}

// Max 最大值节点
func (bst *ItemBinarySearchTree) Max() *Node {
	current := bst.Root
	if current == nil {
		return nil
	}
	for {
		if current.right == nil {
			return current
		} else {
			current = current.right
		}
	}
}

// Search 查找节点
func (bst *ItemBinarySearchTree) searchNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.Key > key {
		return bst.searchNode(node.left, key)
	} else if node.Key < key {
		return bst.searchNode(node.right, key)
	} else {
		return node
	}
}

// Search 查找节点
func (bst *ItemBinarySearchTree) Search(key int) *Node {
	return bst.searchNode(bst.Root, key)
}

// Update 更新节点
func (bst *ItemBinarySearchTree) Update(key int, value Item) bool {
	updateNode := bst.searchNode(bst.Root, key)
	if updateNode != nil {
		updateNode.Value = value
		return true
	}
	return false
}

// InOrderTraversal 中序遍历
func (bst *ItemBinarySearchTree) inOrderTraversal(node *Node, f func(itemNode *Node)) {
	if node != nil {
		bst.inOrderTraversal(node.left, f)
		f(node)
		bst.inOrderTraversal(node.right, f)
	}
}

// InOrderTraversal 中序遍历
func (bst *ItemBinarySearchTree) InOrderTraversal(f func(itemNode *Node)) {
	bst.inOrderTraversal(bst.Root, f)
}

// 通过栈实现中序遍历（非递归:循环迭代）
func (bst *ItemBinarySearchTree) inOrderTraversalByStack(f func(itemNode *Node)) {
	stack := stack.NewStack() // stack/stack.go
	root := bst.Root
	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.left
		}
		root = stack.Pop().(*Node)
		f(root)
		root = root.right
	}
}

// InOrderTraversalByStack 中序遍历（非递归:循环迭代）
func (bst *ItemBinarySearchTree) InOrderTraversalByStack(f func(itemNode *Node)) {
	bst.inOrderTraversalByStack(f)
}

// 先序遍历
func (bst *ItemBinarySearchTree) preOrderTraversal(node *Node, f func(itemNode *Node)) {
	if node != nil {
		f(node)
		bst.preOrderTraversal(node.left, f)
		bst.preOrderTraversal(node.right, f)
	}
}

// PreOrderTraversal 先序遍历
func (bst *ItemBinarySearchTree) PreOrderTraversal(f func(itemNode *Node)) {
	bst.preOrderTraversal(bst.Root, f)
}

func (bst *ItemBinarySearchTree) preOrderTraversalByStack(f func(itemNode *Node)) {
	s := stack.NewStack() // stack/stack.go
	root := bst.Root
	for root != nil || !s.IsEmpty() {
		for root != nil {
			f(root)
			s.Push(root)
			root = root.left
		}
		root = s.Pop().(*Node)
		root = root.right
	}
}

func (bst *ItemBinarySearchTree) PreOrderTraversalByStack(f func(itemNode *Node)) {
	bst.preOrderTraversalByStack(f)
}

// 后序遍历
func (bst *ItemBinarySearchTree) postOrderTraversal(node *Node, f func(itemNode *Node)) {
	if node != nil {
		bst.postOrderTraversal(node.left, f)
		bst.postOrderTraversal(node.right, f)
		f(node)
	}
}

// PostOrderTraversal 后序遍历
func (bst *ItemBinarySearchTree) PostOrderTraversal(f func(itemNode *Node)) {
	bst.postOrderTraversal(bst.Root, f)
}

// 通过栈实现后序遍历（非递归:循环迭代）
func (bst *ItemBinarySearchTree) postOrderTraversalByStack(f func(itemNode *Node)) {
	stack := stack.NewStack()
	prevAccess := &Node{}
	root := bst.Root
	// 首次入栈并遍历左子树
	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.left
		}
		// 遍历右子树，并访问节点

		root = stack.Pop().(*Node) // 弹出栈顶节点并将其赋值给 root
		/*
			检查 root 的右子节点是否为 nil 或者已经被访问过 (prevAccess)。
				- 如果条件成立，则执行函数 f(root)，即对当前节点执行操作，记录访问状态，并将 root 设置为 nil（表示该节点已被完全处理）。
				- 如果条件不成立，说明右子树还未被访问，因此将当前节点再次入栈并将 root 更新为其右子节点，继续后续处理。
		*/
		if root.right == nil || root.right == prevAccess {
			f(root)
			prevAccess = root
			root = nil
		} else {
			stack.Push(root)
			root = root.right
		}
	}
}

// PostOrderTraversalByStack 后序遍历（非递归:循环迭代）
func (bst *ItemBinarySearchTree) PostOrderTraversalByStack(f func(itemNode *Node)) {
	bst.postOrderTraversalByStack(f)
}

// 删除节点
func (bst *ItemBinarySearchTree) deleteNode(node *Node, key int) *Node {
	if node == nil { // 判断节点是否为空
		return nil
	}
	if node.Key > key { // 若当前节点的 key 大于要删除的 key，则递归左子树
		node.left = bst.deleteNode(node.left, key)
	} else if node.Key < key { // 若当前节点的 key 小于要删除的 key，则递归右子树
		node.right = bst.deleteNode(node.right, key)
	} else { // 若当前节点的 key 等于要删除的 key，则删除当前节点
		if node.left == nil { // 果当前节点没有左子树（即 left 为 nil），则返回右子树（node.right）。这样，右子树的所有节点都会接替当前节点的位置
			return node.right
		}
		if node.right == nil { // 当前节点没有右子树（即 right 为 nil），则返回左子树（node.left）
			return node.left
		}
		// 处理两个子节点的情况
		/*
			如果当前节点有两个子节点，即左子节点和右子节点都存在，则需要找到右子树中最小的节点（最左侧的节点）。通过遍历右子树找到这个最小节点 minNode。
			将 minNode 的键值赋值给当前节点，用于替代被删除节点的键值。
			然后在右子树中递归地删除 minNode，以确保不会出现重复值。
		*/
		minNode := node.right
		for minNode.left != nil {
			minNode = minNode.left
		}
		node.Key = minNode.Key
		node.right = bst.deleteNode(node.right, minNode.Key)
	}
	return node
}

// Delete 删除节点
func (bst *ItemBinarySearchTree) Delete(key int) *Node {
	return bst.deleteNode(bst.Root, key)
}

func (bst *ItemBinarySearchTree) printNode(node *Node, depth int, local string) {
	if node != nil {
		format := ""
		for i := 0; i < depth; i++ {
			format += "-"
		}
		format += "["
		depth++
		//format = fmt.Sprintf("%s----- [%d\t-%s\t] [", format)
		bst.printNode(node.left, depth, "left")
		fmt.Printf(format+"%d:%v]- [%d-%s]\n", node.Key, node.Value, depth, local)
		bst.printNode(node.right, depth, "right")
	}
}

func (bst *ItemBinarySearchTree) Print() {
	bst.printNode(bst.Root, 0, "root")
}

//func InOrderAndPostOrder2Tree(inorder, postorder []int) *Node {
//
//}

// 传入中序、后序遍历序列，构造二叉树
func InOrderAndPostOrderBuildTree(inorder []int, postorder []int) *Node {
	// 后序遍历序列最后一个数为根节点，以此划分中序遍历序列
	// 划分中序遍历序列，再以此划分后序遍历序列
	// 一层一层，综合中序、后序进行构造

	// 方法一：显式递归，调用 rebuild() 函数
	// 使用哈希表存储中序遍历数组中每个节点的值和对应的索引
	// 查找根节点在中序数组中的位置时，时间复杂度 O(1)

	// 1. 哈希表存储中序遍历序列
	hash := make(map[int]int)
	for i, v := range inorder {
		hash[v] = i
	}

	// 2. 调用递归函数
	return rebuild(inorder, postorder, len(postorder)-1, 0, len(inorder)-1, hash)

	// 方法二：递归
	// 先构造递归函数 traversal
	// 1. 空节点

	// 2. 找后序遍历的最后一个元素，即当前的中间节点

	// 3. 判断是否为叶子节点

	// 4. 切割中序、后序序列

	// 5. 递归

	// 6. 返回

}

// 返回构造的二叉树。rootIdx 表示根节点在后序数组中的索引，l, r 表示在中序数组中的前后切分点
func rebuild(inorder []int, postorder []int, rootIdx int, l, r int, hash map[int]int) *Node {
	if l > r {
		return nil
	}

	// 从后序数组中找到当前中间节点的值，确定其在中序数组中的位置，然后构造出这个节点的实例
	rootV := postorder[rootIdx]   // 从后序序列获取根节点值
	rootIn := hash[rootV]         // 获取中序序列哈希表中 根节点的位置
	root := &Node{Key: rootV} // 创建根节点

	// 建立左右节点
	// 传入的前后分割点发生改变
	root.left = rebuild(inorder, postorder, rootIdx-(r-rootIn)-1, l, rootIn-1, hash)
	root.right = rebuild(inorder, postorder, rootIdx-1, rootIn+1, r, hash)
	return root
}