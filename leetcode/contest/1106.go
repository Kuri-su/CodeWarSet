package contest

type tmp struct {
	status  int // 1 | , 2 & ,3 !
	sum     bool
	sumFlag bool
}

func parseBoolExpr(expression string) bool {
	stack := make([]tmp, 0)

	now := tmp{}

	for i := 0; i < len(expression); i++ {
		switch {
		case expression[i] == '|':
			stack = append(stack, now)
			now = tmp{}
			now.status = 1

		case expression[i] == '&':
			stack = append(stack, now)
			now = tmp{}
			now.status = 2
		case expression[i] == '!':
			stack = append(stack, now)
			now = tmp{}
			now.status = 3
		case expression[i] == 't' || expression[i] == 'f':
			c := false
			if expression[i] == 't' {
				c = true
			} else {
				c = false
			}
			if now.sumFlag == false {
				if now.status == 3 {
					now.sum = !c
					now.sumFlag = true
				} else {
					now.sum = c
					now.sumFlag = true
				}
			} else {
				switch now.status {
				case 1:
					now.sum = now.sum || c
				case 2:
					now.sum = now.sum && c
				}
			}
		case expression[i] == ')':
			// 出栈
			tmpNow := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			// 计算
			if tmpNow.sumFlag == false {
				if tmpNow.status == 3 {
					tmpNow.sum = !now.sum
					tmpNow.sumFlag = true
				} else {
					tmpNow.sum = now.sum
					tmpNow.sumFlag = true
				}

			} else {
				switch tmpNow.status {
				case 1:
					tmpNow.sum = now.sum || tmpNow.sum
				case 2:
					tmpNow.sum = now.sum && tmpNow.sum

				}
			}

			now = tmpNow
		}
	}

	return now.sum
}
