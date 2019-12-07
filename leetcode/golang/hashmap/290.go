package hashmap

/*
   解题思想:
	双向哈希表
​		使用双向哈希表保证唯一的对应关系而又避免了为了保证反向的对应关系而需要的搜索操作

   时间复杂度: O(n)
   空间复杂度: O(lgn)

   过程:
*/

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	var m = make(map[string]string)
	var m2 = make(map[string]string)
	pSlice := strings.Split(pattern, "")
	slice := strings.Split(str, " ")
	if len(pSlice) != len(slice) {
		return false
	}

	for k, v := range pSlice {
		if m[v] == "" && m2[slice[k]] == "" {
			m[v] = slice[k]
			m2[slice[k]] = v
		} else if m[v] != slice[k] {
			return false
		}
	}
	return true
}
