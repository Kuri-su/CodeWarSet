package contest

func removeDuplicates(S string) string {
	for {
		var tmp int32
		var flag = false
		for key, value := range S {
			if key == 0 {
				tmp = value
				continue
			}
			if value == tmp {
				S = string(append([]byte(S[:key-1]), []byte(S[key+1:])...))
				flag = true
				break
			}
			tmp = value
		}
		if !flag {
			break
		}
	}
	return S
}
