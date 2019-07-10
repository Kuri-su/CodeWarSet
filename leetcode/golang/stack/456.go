package stack

/*
* 利用 栈 处理此问题

* 解题思想:
    * 理清题目意思..
    * 序列中 A < B < C
    * 值  aA < aC < aB  翻译一下 数字 A < B <C , 在数组中  aB 的值最大, aC的值小于 aB , 但是大于aA
    *  ------A----B---C------
             |    |   |
             |     ↘ ↙
             v     ↙ ↘
            aA < aC < aB

* 时间复杂度: O(nLogn)
* 空间复杂度: O(1)

* 过程:
    * 声明一个 数组(栈), 用来存放 aB 的候选值, 并且搭配倒叙遍历, 保证 A < B < C 的关系
    * 在遇到新的值的时候, 判断 新的值是否大于栈里的元素, 如果大于 aB 的候选, 那么说明这个候选不适合, 淘汰掉, 给 aC, 也就是 last, 这里也就隐含了 aB stack > aC
    * 所以, 接着只要 找到一个小于 aC 的元素即可,

* 祝你顺利 Pass
*/

func find132pattern(nums []int) bool {
	last := -1 << 31            // ACB 中的 B, 也就是 aC, 先给 last 取了极小值, 如果这个值已经有了值, 说明已经找到了B  也就是aB,
	stack := make([]int, 0, 64) // aB 的候选值

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < last { // 这里说明 last 已经不再是极小值,已经找到了 中间那个元素的值,
			return true
		}

		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			last = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return false

}
