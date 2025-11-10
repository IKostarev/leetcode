package main

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	spl := strings.Split(strconv.Itoa(x), "")

	for i, j := 0, len(spl)-1; i < len(spl)/2; i-- {
		if spl[i] != spl[j] {
			return false
		}

		j--
	}

	return true
}
