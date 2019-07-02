package contest

import (
	"sort"
	"strings"
)

func permute(S string) []string {
	if !strings.Contains(S, "{") {
		return []string{S}
	}

	stringxs := strings.Split(S, "}")

	arr := make([][]string, 0)

	for key, value := range stringxs {
		if strings.Contains(value, "{") {
			if key != len(stringxs)-1 {
				tmp := strings.Split(value, "{")
				for key, value := range tmp {
					if value == "" {
						tmp = append(tmp[0:key], tmp[key+1:]...)
					}
				}
				stringxs = append(stringxs[0:key], append(tmp, stringxs[key+1:]...)...)
			} else {
				tmp := strings.Split(value, "{")
				stringxs = append(stringxs[0:key], tmp...)
			}
		}
	}

	for _, value := range stringxs {
		arr = append(arr, strings.Split(value, ","))
	}

	deep := 0

	result := make([]string, 0)
	for _, value := range arr[deep] {
		result = append(result, dfs(""+string(value), deep+1, arr)...)
	}

	// fmt.Printf("%+v", result)

	sort.Strings(result)

	return result
}

func dfs(str string, deep int, arr [][]string) []string {
	if deep == len(arr) {
		return []string{str}
	}
	tmp := make([]string, 0)
	for _, value := range arr[deep] {
		tmp = append(tmp, dfs(str+string(value), deep+1, arr)...)
	}
	return tmp
}
