package p7

import (
	"fmt"
	"testing"
)


func traversal(root *Node, rlt *[]int) {
	if root == nil {
		return
	}
	traversal(root.left, rlt)
	traversal(root.right, rlt)
	// 因为预先分配了足够的大小，所以可以保证append
	// 不会因为超出cap从而分配新的底层数组
	*rlt = append(*rlt, root.data)
}
func TestBuildBT(t *testing.T) {
	testCases := []struct {
		pre  []int
		in   []int
		post []int
	}{
		{
			pre:  []int{1, 2, 4, 7, 3, 5, 6, 8},
			in:   []int{4, 7, 2, 1, 5, 3, 8, 6},
			post: []int{7, 4, 2, 5, 8, 6, 3, 1},
		},
	}

	for _, tc := range testCases {
		root := BuildBT(tc.pre, tc.in)
		rlt := make([]int, 0, len(tc.pre))
		traversal(root, &rlt)

		// 因为这里确保是一维切片，所以可以用这种方法
		// 不然推荐用reflect.DeepEqual()
		if fmt.Sprint(tc.post) != fmt.Sprint(rlt) {
			t.Errorf("tc.pre:%v,tc.in:%v,tc.post:%v,your rlt:%v", tc.pre, tc.in, tc.post, rlt)
		}

	}
}
