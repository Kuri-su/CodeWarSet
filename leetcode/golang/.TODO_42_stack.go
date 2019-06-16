package golang

// 使用栈来处理

func trap1(height []int) int {
	stack := make([]int, 0)
	for key, value := range height {
		if value != 0 {
			if len(stack) == 0 {
				stackPush(stack, key)
			}
			stackBack := stackBack(stack)
			if stackBack > value {
				stackPush(stack, key)
			} else {

			}

		}
	}

}

func stackPush(stack []int, a int) []int {
	return append(stack, a)
}
func stackPop(stack []int) int {
	a := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	return a
}
func stackBack(stack []int) int {
	return stack[len(stack)-1]
}
