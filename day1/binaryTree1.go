package day1

/*
给定一个非负整数n，代表二叉树的节点个数。返回能形成多少种不同的二叉树结构
*/

func binaryNumber(n int)int{
	if n < 2{
		return 1
	}
	res := 0
	for i := 1; i <= n; i++{  // i代表头节点的值
		leftWays:= binaryNumber(i-1) // 左子树的种数
		rightWays:= binaryNumber(n-i) // 右子树的种数
		res += leftWays * rightWays // 左子树的种数 * 右子树的种数
	}
	return res
}

func binaryNumberDp(n int)int{
	if n < 2{
		return 1
	}
	dp := make([]int, n+1) // dp[i]代表i个节点的二叉树有多少种结构
	dp[0] = 1 // 0个节点的二叉树有1种结构
	for i := 1; i <= n; i++{ // i代表头节点的值
		for j := 1; j <= i; j++{ // j代表头节点的值
			dp[i] += dp[j-1] * dp[i-j] // 左子树的种数 * 右子树的种数
		}
	}
	return dp[n]
}