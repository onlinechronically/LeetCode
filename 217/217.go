package main

func containsDuplicate(nums []int) bool {
	hMap := make(map[int]int)
	for _, value := range nums {
		if _, inMap := hMap[value]; !inMap {
			hMap[value] = 1
		} else {
			return true
		}
	}
	return false
}
