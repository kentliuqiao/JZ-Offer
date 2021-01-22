package interview8

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func getNext(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	var next *TreeNode
	if node.Right != nil {
		// 右子树不为空，则其下一节点为该节点右子树的最左节点
		next = node.Right
		for next.Left != nil {
			next = next.Left
		}

		return next
	} else {
		// 右子树为空，则其下一节点分为以下情况
		// 1. 沿其x父节点一直向上遍历，直到找到一个其中的父节点a为该节点a的父节点aP的左子节点，那么此父节点a的父节点aP就是x节点的下一个中序节点
		// 2. 如果沿着方式1的路劲一直遍历到根节点，由于根节点的父节点为nil，故x节点没有下一个节点，x就是中序便利的最后一个节点
		// next = node.Parent
		// for next == next.Parent.Right {
		// 	next = next.Parent
		// }
		// if next.Parent != nil {
		// 	next = next.Parent
		// } else {
		// 	next = nil
		// }
		for node.Parent != nil {
			if node.Parent.Left == node {
				return node.Parent
			} else {
				node = node.Parent
			}
		}
	}

	return nil
}
