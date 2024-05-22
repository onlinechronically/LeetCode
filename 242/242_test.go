package main

import (
	"strings"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := map[string]bool{
		"anagram,nagaram": true,
		"rat,car":         false,
	}
	counter := 0
	for k, v := range tests {
		counter++
		arr := strings.Split(k, ",")
		ret := isAnagram(arr[0], arr[1])
		if ret != v {
			t.Fatalf("%q returned %t but expected %t", k, ret, v)
		}
	}
}
