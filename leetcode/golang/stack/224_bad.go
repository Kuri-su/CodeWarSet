package stack

import (
	"log"
	"strconv"
	"unicode"
)

func calculate1(s string) int {
	stack := make([]string, 0)

	for i := 0; i <= len(s)-1; i++ {
		switch string(s[i]) {
		case ")":
			// 开始出栈
			arr := make([]string, 0)

			// 寻找边界
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] != "(" {
					arr = append(arr, stack[i])
					// 出栈
					stack = stack[0 : len(stack)-1]
				} else {
					stack = stack[0 : len(stack)-1]
					break
				}
			}

			ccc := make([]string, 0)

			for i := len(arr) - 1; i >= 0; i-- {
				ccc = append(ccc, arr[i])
			}

			// 计算
			stack = append(stack, cal(ccc))
		case " ":
			continue

		default:
			// 入栈
			if unicode.IsNumber(rune(s[i])) {
				tmp := ""
				for {
					if unicode.IsNumber(rune(s[i])) {
						tmp += string(s[i])
						if i == len(s)-1 {
							break
						} else {
							i++
						}
					} else {
						i--
						break
					}
				}
				stack = append(stack, tmp)
			} else {
				stack = append(stack, string(s[i]))
			}

		}
	}

	strsum := cal(stack)
	i, e := strconv.Atoi(strsum)
	if e != nil {
		log.Fatalln(e)
	}
	return i
}

func cal(arr []string) string {
	sum := 0
	t := 0
	for i := 0; i < len(arr); i++ {
		if isSymbol(arr[i]) {
			c, err := strconv.Atoi(arr[i+1])
			if err != nil {
				log.Fatalln(err)
			}
			if arr[i] == "+" {
				sum += t + c
			} else {
				sum += t - c
			}
			t = 0
			i++
		} else {
			tcc, err := strconv.Atoi(arr[i])
			t = tcc
			if err != nil {
				log.Fatalln(err)
			}

		}
	}

	if sum == 0 && t != 0 {
		sum += t
	}

	st := strconv.Itoa(sum)
	return st
}

func isSymbol(s string) bool {
	switch s {
	case "+":
		return true
	case "-":
		return true
	default:
		return false
	}
}
