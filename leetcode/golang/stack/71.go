package stack

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
