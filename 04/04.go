package p4

type point struct {
	x int
	y int
}

func newP(x, y int) point {
	return point{x, y}
}

var mp map[point]bool

// Find 二维数组的查找,(x,y)为起始点坐标
func Find(arr [][]int, num int, x, y int, mp map[point]bool) bool {
	if mp == nil {
		mp = make(map[point]bool)
	}
	m := len(arr)
	n := len(arr[0])
	if x >= m || y >= n {
		print("超出边界")
		return false
	}
	p := newP(x, y)
	// 已被访问过
	if mp[p] == true {
		print("被访问过")
		return false
	}
	mp[p] = true
	if num < arr[x][y] {
		return false
	} else if num == arr[x][y] {
		return true
	} else {
		return Find(arr, num, x+1, y, mp) || Find(arr, num, x, y+1, mp)
	}
}
