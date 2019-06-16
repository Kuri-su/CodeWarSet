package golang

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			var tmp = make([]int, len(arr))
			copy(tmp, arr)
			for j, p := i+1, i; j < len(arr); j++ {
				arr[j] = tmp[p]
				p++
			}
			i++
		}
	}
	// fmt.Printf("%+v", arr)
}
