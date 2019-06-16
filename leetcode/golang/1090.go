package golang

import "sort"

var vvs, lls xz

type xz []int

// Len()
func (s xz) Len() int {
	return len(s)
}

// Less():成绩将有低到高排序
func (s xz) Less(i, j int) bool {
	return s[i] > s[j]
}

// Swap()
func (s xz) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	lls[i], lls[j] = lls[j], lls[i]

}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	// 排序
	vvs = values
	lls = labels

	sort.Sort(vvs)

	var ns []int // 用来记录已经的取值
	var lablexs = make(map[int]int)

	for i := 0; i < len(values); i++ {

		if lablexs[labels[i]] >= use_limit {
			continue
		}
		ns = append(ns, values[i])
		lablexs[labels[i]] = lablexs[labels[i]] + 1

		if len(ns) == num_wanted {
			// fmt.Printf("%+v \n", ns)

			var sum int
			for _, value := range ns {
				sum += value
			}
			return sum

		}
	}

	var sum int
	// fmt.Printf("%+v \n", ns)
	for _, value := range ns {
		sum += value
	}
	return sum

}
