package p4_test

import (
	"testing"

	. "github.com/wisecsj/gocn/04"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		arr   [][]int
		num   int
		exist bool
	}{
		// 存在，值介于max和min之间
		{
			arr: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			num:   7,
			exist: true,
		},
		// 存在，值等于max
		{
			arr: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			num:   15,
			exist: true,
		},
		// 存在，值等于min
		{
			arr: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			num:   1,
			exist: true,
		},
		// 不存在，值大于max
		{
			arr: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			num:   20,
			exist: false,
		},
		// 不存在，值小于min
		{
			arr: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			num:   0,
			exist: false,
		},
		// 不存在，介于max和min之间
		{
			arr: [][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			num:   3,
			exist: false,
		},
	}
	for _, tC := range testCases {
		r := Find(tC.arr, tC.num, 0, 0, nil)
		if r != tC.exist {
			t.Errorf("%v whether exists in %v,should be %v,but got %v", tC.num, tC.arr, tC.exist, r)
		}
	}
}
