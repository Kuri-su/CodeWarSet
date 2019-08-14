package stack

/*
   利用 栈 处理此问题

   解题思想: 这题比较简单, 用栈的思路能很快求解,

   时间复杂度: O(n)
   空间复杂度: O(1)

   过程:
       文件名入栈, .. 出栈
*/

import "strings"

func simplifyPath(path string) string {
	folders := strings.Split(path, "/")
	stack := make([]string, 0, 100)
	for _, folderName := range folders {
		if folderName == ".." {
			if len(stack) == 0 {
				continue
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else if folderName == "" || folderName == "." {
			continue
		} else {
			stack = append(stack, folderName)
		}
	}
	resPath := ""
	if len(stack) == 0 {
		resPath += "/"
	} else {
		for _, value := range stack {
			resPath += "/" + value
		}
	}
	return resPath

}
