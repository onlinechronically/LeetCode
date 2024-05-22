package main

import "strings"

func isAnagram(s string, t string) bool {
	s = strings.ToLower(s)
	t = strings.ToLower(t)
	letters := make(map[rune]int)
	for _, letter := range s {
		_, inMap := letters[letter]
		if inMap {
			letters[letter]++
		} else {
			letters[letter] = 1
		}
	}
	for _, letter := range t {
		_, inMap := letters[letter]
		if !inMap {
			return false
		} else {
			letters[letter]--
		}
	}
	for _, val := range letters {
		if val != 0 {
			return false
		}
	}
	return true
}
