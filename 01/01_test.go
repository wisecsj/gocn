package p1_test

import (
	"testing"

	. "github.com/wisecsj/gocn/01"
)

var data = []struct {
	in  []int
	out []int
}{
	{[]int{0, 1, 2, 3}, []int{-1}},
	{[]int{0, 1, 1, 0}, []int{0, 1}},
}

func TestFind(t *testing.T) {
	for _, tt := range data {
		r := Find(tt.in)
		if !in(r, tt.out) {
			t.Fatalf("case %v,got %d, want %v", tt.in, r, tt.out)
		}
	}
}

func in(t int, arr []int) bool {
	for _, v := range arr {
		if t == v {
			return true
		}
	}
	return false
}
