package p5_test

import (
	"testing"

	. "github.com/wisecsj/gocn/05"
)

func TestReplaceSpace(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{
			in:  "we are happy",
			out: "we%20are%20happy",
		},
	}

	for _, tc := range testCases {
		r := ReplaceSpace(tc.in)
		if tc.out != r {
			// t.Logf("%v---%v", tc.in, r)
			t.Errorf("%v failed,got %v,should be %v", tc.in, r, tc.out)
		}
	}
}
