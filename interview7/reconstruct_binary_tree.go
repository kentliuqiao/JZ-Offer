package interview7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reconstruct(pre, mid []int) *TreeNode {
	if len(pre) <= 0 || len(mid) <= 0 || len(pre) != len(mid) {
		return nil
	}

	return reconstructBinaryTree(pre, 0, len(pre)-1, mid, 0, len(mid)-1)
}

func reconstructBinaryTree(pre []int, preStart, preEnd int, mid []int, midStart, midEnd int) *TreeNode {
	if preStart > preEnd || midStart > midEnd {
		return nil
	}

	// root node
	root := &TreeNode{Val: pre[preStart]}

	for i := midStart; i <= midEnd; i++ {
		if mid[i] == root.Val {
			root.Left = reconstructBinaryTree(pre, preStart+1, preStart+i-midStart, mid, midStart, i-1)
			root.Right = reconstructBinaryTree(pre, preStart+i-midStart+1, preEnd, mid, i+1, midEnd)

			break
		}
	}

	return root
}

func reconstruct2(preorder []int, inorder []int) *TreeNode {
	for k, v := range inorder {
		if v == preorder[0] {
			return &TreeNode{
				Val:   v,
				Left:  reconstruct2(preorder[1:k+1], inorder[:k]),
				Right: reconstruct2(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}
