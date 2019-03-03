package p7

// Node 节点
type Node struct {
	data        int
	left, right *Node
}

func newNode(data int) *Node {
	return &Node{data: data}
}

func index(data int, in []int) int {
	for i, v := range in {
		if data == v {
			return i
		}
	}
	panic("遍历序列元素出错")
}

// BuildBT 重建二叉树
func BuildBT(pre, in []int) *Node {
	length := len(pre)
	if length == 0 {
		return nil
	}
	rootValue := pre[0]

	root := newNode(rootValue)
	rootPos := index(rootValue, in)

	// 左子树的先序遍历start index
	psl := 1
	pel := rootPos + 1
	// inorder
	insl := 0
	inel := rootPos

	root.left = BuildBT(pre[psl:pel], in[insl:inel])
	root.right = BuildBT(pre[pel:], in[inel+1:])

	return root

}
