package contest

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) > len(str2) {
		if str1[0:len(str2)] != str2 {
			return ""
		}
	} else {
		if str2[0:len(str1)] != str1 {
			return ""
		}
	}

	if str1 == str2 {
		return str1
	}
	var i int
	var long, short string

	if len(str1) > len(str2) {
		i = len(str2)
		long = str1
		short = str2
	} else {
		i = len(str1)
		long = str2
		short = str1
	}

	for ; i >= 1; i-- {
		j, k := 0, i
		flag := true
		res := ""
		for {
			if k+i == len(long) && long[j:k] == long[k:k+i] {
				res = long[j:k]
				break
			} else if k+i > len(long) {
				flag = false
				break
			} else if k+i < len(long) && long[j:k] == long[k:k+i] {
				j = k
				k = k + i
			} else {
				flag = false
				break
			}
		}

		if flag {
			j, k := 0, len(res)
			for {
				if (j+1)*k == len(short) && res == short[j*k:(j+1)*k] {
					break
				} else if j*k > len(short) {
					flag = false
					break
				} else if (j+1)*k < len(short) && res == short[j*k:(j+1)*k] {
					j++
				} else {
					flag = false
					break
				}
			}
		}

		if flag {
			return long[0:i]
		}
	}
	return ""
}
