package main

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}

	for _, char := range s {
		if _, in := pairs[char]; in {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}
		lastChar := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if a := pairs[lastChar]; a != char {
			return false
		}
	}

	return len(stack) == 0
}
