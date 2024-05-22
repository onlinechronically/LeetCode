package main

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val, inMap := targetMap[nums[i]]
		if inMap {
			return []int{val, i}
		} else {
			targetMap[target-nums[i]] = i
		}
	}
	return []int{-1, -1}
}
