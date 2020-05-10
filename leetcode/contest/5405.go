package contest

func countTriplets(arr []int) int {
	var res int
	lenarr := len(arr)

	if len(arr) < 2 {
		return 0
	}

	for i := 0; i < lenarr-1; i++ {
		for j := i + 1; j < lenarr; j++ {
			//g:
			for k := j; k < lenarr; k++ {
				aTmp := arr[i]
				for l := i + 1; l < j; l++ {
					aTmp ^= arr[l]
				}

				bTmp := arr[j]
				for l := j + 1; l <= k; l++ {
					bTmp ^= arr[l]
					//if bTmp == aTmp {
					//	res++
					//	break g
					//
					//}
				}
				if aTmp == bTmp {
					//fmt.Println(i, j, k)
					res++
				}

			}
		}
	}
	return res
}
