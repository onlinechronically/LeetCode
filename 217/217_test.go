package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := map[string]bool{
		"1,2,3,1":             true,
		"1,2,3,4":             false,
		"1,1,1,3,3,4,3,2,4,2": true,
	}
	counter := 0
	for k, v := range tests {
		counter++
		arr := make([]int, len(strings.Split(k, ",")))
		for i, s := range strings.Split(k, ",") {
			val, _ := strconv.Atoi(s)
			arr[i] = val
		}
		ret := containsDuplicate(arr)
		if ret != v {
			t.Fatalf("%q returned %t but expected %t", k, ret, v)
		}
	}
}
