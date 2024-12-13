package main

// TreeNode  AVL 树节点结构体
type TreeNode struct {
	Val    int       // 节点值
	Height int       // 节点高度
	Left   *TreeNode // 左子节点引用
	Right  *TreeNode // 右子节点引用
}

// NewTreeNode 构造方法
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil, // 左子节点指针
		Right: nil, // 右子节点指针
		Val:   v,   // 节点值
	}
}

/* 获取节点高度 */
func (t *TreeNode) height(node *TreeNode) int {
	// 空节点高度为 -1 ，叶节点高度为 0
	if node != nil {
		return node.Height
	}
	return -1
}

/* 更新节点高度 */
func (t *TreeNode) updateHeight(node *TreeNode) {
	lh := t.height(node.Left)
	rh := t.height(node.Right)
	// 节点高度等于最高子树高度 + 1
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

/* 获取平衡因子 */
func (t *TreeNode) balanceFactor(node *TreeNode) int {
	// 空节点平衡因子为 0
	if node == nil {
		return 0
	}
	// 节点平衡因子 = 左子树高度 - 右子树高度
	return t.height(node.Left) - t.height(node.Right)
}

/* 右旋操作 */
func (t *TreeNode) rightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	// 以 child 为原点，将 node 向右旋转
	child.Right = node
	node.Left = grandChild
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

/* 左旋操作 */
func (t *TreeNode) leftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	// 以 child 为原点，将 node 向左旋转
	child.Left = node
	node.Right = grandChild
	// 更新节点高度
	t.updateHeight(node)
	t.updateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

/* 执行旋转操作，使该子树重新恢复平衡 */
func (t *TreeNode) rotate(node *TreeNode) *TreeNode {
	// 获取节点 node 的平衡因子
	// Go 推荐短变量，这里 bf 指代 t.balanceFactor
	bf := t.balanceFactor(node)
	// 左偏树
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			// 右旋
			return t.rightRotate(node)
		} else {
			// 先左旋后右旋
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}
	// 右偏树
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			// 左旋
			return t.leftRotate(node)
		} else {
			// 先右旋后左旋
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}
	// 平衡树，无须旋转，直接返回
	return node
}

/* 插入节点 */
func (t *TreeNode) insert(val int) {
	t = t.insertHelper(t, val)
}

/* 递归插入节点（辅助函数） */
func (t *TreeNode) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	/* 1. 查找插入位置并插入节点 */
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		// 重复节点不插入，直接返回
		return node
	}
	// 更新节点高度
	t.updateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.rotate(node)
	// 返回子树的根节点
	return node
}

func main() {

}
