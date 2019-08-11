package contest

func isMajorityElement(nums []int, target int) bool {
    sum := 0
    numsLenHalf := len(nums) / 2
    for _, value := range nums {
        if value == target {
            sum++
            if sum > numsLenHalf {
                return true
            }
        }
    }
    return false
}
