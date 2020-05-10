package contest

const push = "Push"
const pop = "Pop"

func buildArray(target []int, n int) []string {
	var (
		res          []string
		successCount int
	)

	for i := 1; i <= n; i++ {

		for _, v := range target {
			if v == i {
				res = append(res, push)
				successCount++
				break
			}
			if v > i {
				res = append(res, push, pop)
				break
			}
		}
		if successCount >= len(target) {
			return res
		}

	}
	return []string{"-1"}
}
