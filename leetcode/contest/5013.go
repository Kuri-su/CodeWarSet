package contest

import "sort"

type IntSlice [][]int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	if c[i][0] == c[j][0] {
		return c[i][1] < c[j][1]
	} else {
		return c[i][0] < c[j][0]
	}
}

func indexPairs(text string, words []string) [][]int {
	var res IntSlice
	for key, value := range text {
		for _, word := range words {
			if value == int32(word[0]) {
				flag := true
				for charIndex, char := range word {
					if key+charIndex > len(text)-1 {
						flag = false
						break
					}
					if int32(text[key+charIndex]) != char {
						flag = false
					}
				}
				if flag {
					res = append(res, []int{key, key + len(word) - 1})
				}
			}
		}
	}
	sort.Sort(res)

	return res
}
