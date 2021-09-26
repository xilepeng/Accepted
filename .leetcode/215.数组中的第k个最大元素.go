/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	return quick_select(nums, 0, n-1, n-k)
}

func quick_select(A []int, start, end, i int) int {
	if piv_pos := random_partition(A, start, end); i == piv_pos {
		return A[i]
	} else if i < piv_pos {
		return quick_select(A, start, piv_pos-1, i)
	} else {
		return quick_select(A, piv_pos+1, end, i)
	}
}

func partition(A []int, start, end int) int {
	piv, i := A[start], start+1
	for j := start + 1; j <= end; j++ {
		if A[j] < piv {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[start], A[i-1] = A[i-1], A[start]
	return i - 1
}

func random_partition(A []int, start, end int) int {
	random := rand.Int()%(end-start+1) + start
	A[start], A[random] = A[random], A[start]
	return partition(A, start, end)
}

// @lc code=end

