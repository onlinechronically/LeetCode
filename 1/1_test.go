package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string][]int{
		"2,7,11,15 9": {0, 1},
		"3,2,4 6":     {1, 2},
		"3,3 6":       {0, 1},
	}
	for k, v := range tests {
		splitOne := strings.Split(k, " ")
		target, _ := strconv.Atoi(splitOne[1])
		nums := strings.Split(splitOne[0], ",")
		arr := make([]int, len(nums))
		for i, s := range nums {
			val, _ := strconv.Atoi(s)
			arr[i] = val
		}
		ret := twoSum(arr, target)
		if ret[0] != v[0] || ret[1] != v[1] {
			t.Fatalf("%q returned %q but expected %q", k, ret, v)
		}
	}
}
