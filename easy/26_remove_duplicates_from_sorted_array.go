package easy

func removeDuplicates(nums []int) int {
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[cur] {
			cur = cur + 1
			nums[cur] = nums[i]
		}
	}
	return cur + 1
}
