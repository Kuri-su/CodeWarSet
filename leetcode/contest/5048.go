package contest

func countCharacters(words []string, chars string) int {

	var m = make(map[int32]int)
	for _, value := range chars {
		m[value] += 1
	}

	var num int

	for _, word := range words {
		flag := true
		var mm = make(map[int32]int)

		for _, char := range word {
			if _, ok := m[char]; !ok {
				flag = false
				break
			}
			if mm[char] >= m[char] {
				flag = false
				break
			}
			mm[char] += 1
		}

		if flag {
			num += len(word)
		}
	}
	return num

}
