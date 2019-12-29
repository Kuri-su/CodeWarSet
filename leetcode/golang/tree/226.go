package tree

/*
   解题思想:
	广搜

   时间复杂度: O(n)
   空间复杂度: O(1)

   过程:
*/

type TreeNode226 struct {
	Val   int
	Left  *TreeNode226
	Right *TreeNode226
}

var queue []*TreeNode226

func invertTree(root *TreeNode226) *TreeNode226 {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	queue = append(queue, root)
	do()
	return root
}

func do() {
	remove := func() *TreeNode226 {
		node := queue[0]
		queue = queue[1:]
		return node
	}
	var node *TreeNode226
	if len(queue) > 0 {
		node = remove()
	} else {
		return
	}
	// 交换
	node.Right, node.Left = node.Left, node.Right

	if node.Right != nil {
		queue = append(queue, node.Right)
	}
	if node.Left != nil {
		queue = append(queue, node.Left)
	}
	do()
}
