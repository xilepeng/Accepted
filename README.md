

[76. 最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/) next

[129. 求根节点到叶节点数字之和](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/)

### 方法一：深度优先搜索
思路与算法

从根节点开始，遍历每个节点，如果遇到叶子节点，则将叶子节点对应的数字加到数字之和。如果当前节点不是叶子节点，则计算其子节点对应的数字，然后对子节点递归遍历。

![](https://assets.leetcode-cn.com/solution-static/129/fig1.png)

![](https://pic.leetcode-cn.com/1603933660-UNWQbT-image.png)

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int

	dfs = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}

	return dfs(root, 0)
}
```

复杂度分析

- 时间复杂度：O(n)，其中 n 是二叉树的节点个数。对每个节点访问一次。

- 空间复杂度：O(n)，其中 n 是二叉树的节点个数。空间复杂度主要取决于递归调用的栈空间，递归栈的深度等于二叉树的高度，最坏情况下，二叉树的高度等于节点个数，空间复杂度为 O(n)。



[958. 二叉树的完全性检验](https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/)

### 方法一：广度优先搜索
1. 按 根左右(前序遍历) 顺序依次检查
2. 如果出现空节点，标记end = true
3. 如果后面还有节点，返回false

```go
func isCompleteTree(root *TreeNode) bool {
	q, end := []*TreeNode{root}, false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			end = true
		} else {
			if end == true {
				return false
			}
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}
	return true
}
```


[468. 验证IP地址](https://leetcode-cn.com/problems/validate-ip-address/)

[剑指 Offer 21. 调整数组顺序使奇数位于偶数前面](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

[322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)

[剑指 Offer 09. 用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

[122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)

[7. 整数反转](https://leetcode-cn.com/problems/reverse-integer/)

[460. LFU 缓存](https://leetcode-cn.com/problems/lfu-cache/)