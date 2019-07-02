package contest

import "sort"

type ax [][]int

// Len()
func (s ax) Len() int {
	return len(s)
}

// Less():成绩将有低到高排序
func (s ax) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

// Swap()
func (s ax) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func highFive(items [][]int) [][]int {
	a := make(map[int][]int)
	for _, value := range items {
		id := value[0]
		score := value[1]

		a[id] = append(a[id], score)
	}

	result := make([][]int, 0)

	for id, scores := range a {
		sort.Sort(sort.Reverse(sort.IntSlice(scores)))
		avg := (scores[0] + scores[1] + scores[2] + scores[3] + scores[4]) / 5
		result = append(result, []int{id, avg})
	}

	var axx ax
	for _, value := range result {
		axx = append(axx, value)
	}
	sort.Sort(axx)

	return axx
}
