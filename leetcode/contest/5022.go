package contest

func numKLenSubstrNoRepeats(S string, K int) int {
	total := 0

	for i := 0; i < len(S); i++ {
		m := map[string]int{}

		for j := i; j < len(S); j++ {
			if m[string(S[j])] == 0 {
				m[string(S[j])] = 1
			} else {
				break
			}
			if len(m) == K {
				total++
				break
			}
		}
	}
	return total
}
