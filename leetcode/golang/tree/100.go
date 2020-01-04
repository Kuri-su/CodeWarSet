package tree

/*
  解题思想:
	前序遍历递归比较两棵树的状态

   时间复杂度: O(logn)
   空间复杂度: O(n)
*/

type TreeNode100 struct {
	Val   int
	Left  *TreeNode100
	Right *TreeNode100
}

func isSameTree(p *TreeNode100, q *TreeNode100) bool {
	return preorderTraversal(p, q)
}

func preorderTraversal(p, q *TreeNode100) bool {
	if p == q && p == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !preorderTraversal(p.Left, q.Left) {
		return false
	}

	if !preorderTraversal(p.Right, q.Right) {
		return false
	}
	return true
}
