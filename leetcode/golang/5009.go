package golang

var distanceArr [][]int

func assignBikes(workers [][]int, bikes [][]int) int {
	distanceArr = [][]int{}
	for _, value := range workers {
		distances := make([]int, 0)
		for _, value2 := range bikes {
			a := calDistance(value[0], value[1], value2[0], value2[1])
			distances = append(distances, a)
		}
		distanceArr = append(distanceArr, distances)
	}

	min := -1
	for i := 0; i < len(distanceArr[0]); i++ {
		if min == -1 {
			min = findMin(i, 0, 0, []int{})
		} else {
			now := findMin(i, 0, 0, []int{})
			if min > now {
				min = now
			}
		}
	}

	return min
}

func findMin(index, sum, deep int, usedArr []int) int {
	usedArr = append(usedArr, index)
	sum += distanceArr[deep][index]

	if deep == len(distanceArr)-1 {
		return sum
	}

	min := -1
	for i := 0; i < len(distanceArr[deep]); i++ {
		flag := true
		for _, j := range usedArr {
			if j == i {
				flag = false
				break
			}
		}
		if flag {
			if min == -1 {
				min = findMin(i, sum, deep+1, usedArr)
			} else {
				now := findMin(i, sum, deep+1, usedArr)
				if min > now {
					min = now
				}
			}

		}
	}
	return min
}

func calDistance(x1, y1, x2, y2 int) int {
	xC := x1 - x2
	yC := y1 - y2
	if xC < 0 {
		xC *= -1
	}
	if yC < 0 {
		yC *= -1
	}
	return xC + yC
}
