
package p1

// Find 返回数组中重复的数字（任一个）
func Find(arr []int) int {
	mp := make([]bool, cap(arr))

	for _, v := range arr {
		if mp[v] == true {
			return v
		}
		mp[v] = true
	}
	return -1
}
