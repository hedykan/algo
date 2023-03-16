package chain

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

var resArr []string

func generateParenthesis(n int) []string {
	resArr = make([]string, 0)
	generateItem("", n, n)
	return resArr
}

func generateItem(res string, left, right int) {
	// 非法退出
	if right < left {
		return
	}
	if right < 0 || left < 0 {
		return
	}

	// 合法返回
	if right == 0 && left == 0 {
		resArr = append(resArr, res)
		return
	}

	// 未完继续
	for i := 0; i < 2; i++ {
		if i == 0 {
			generateItem(res+"(", left-1, right)
		} else {
			generateItem(res+")", left, right-1)
		}
	}
}

// @lc code=end
