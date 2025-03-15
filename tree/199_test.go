package tree

import "testing"

// TreeNode 定义二叉树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createTree 根据层序遍历数组创建二叉树
func createTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		current := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if i < len(vals) && vals[i] != nil {
			current.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		// 处理右子节点
		if i < len(vals) && vals[i] != nil {
			current.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

func TestRight(t *testing.T) {
	// 给定的层序遍历数组，其中 nil 表示空节点
	vals := []interface{}{1, 2, 3, nil, 5, nil, 4}

	// 创建二叉树
	root := createTree(vals)

	res := make([]int, 0)
	travel := make([]*TreeNode, 0)

	travel = append(travel, root)
	res = append(res, root.Val)
	res = TravelHelp(root, travel, res)

	for _, v := range res {
		println(v)
	}
}

func TravelHelp(node *TreeNode, travel []*TreeNode, res []int) []int {
	for len(travel) > 0 {

		l := len(travel)
		println(travel[0].Left)
		println(travel[0].Right)
		for i := 0; i < l; i++ {
			if travel[i].Left != nil {
				travel = append(travel, travel[i].Left)
			}
			if travel[i].Right != nil {
				travel = append(travel, travel[i].Right)
			}
		}
		if len(travel) > l {
			res = append(res, travel[len(travel)-1].Val)
		}

		travel = travel[l:]
	}
	for _, v := range res {
		println(v)
	}
	return res
}

func check(nums []int) {
	nums[0] = 10
}
