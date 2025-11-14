package rpmver

import "strings"

func RpmVersionCompare(a, b string) int {
	if a == b {
		return 0
	}
	i, j := 0, 0
	la, lb := len(a), len(b)
	for i < la || j < lb {
		for i < la && !isAlnum(a[i]) && a[i] != '~' && a[i] != '^' {
			i++
		}
		for j < lb && !isAlnum(b[j]) && b[j] != '~' && b[j] != '^' {
			j++
		}
		if (i < la && a[i] == '~') || (j < lb && b[j] == '~') {
			if i >= la || a[i] != '~' {
				return 1
			}
			if j >= lb || b[j] != '~' {
				return -1
			}
			i++
			j++
			continue
		}
		if (i < la && a[i] == '^') || (j < lb && b[j] == '^') {
			if i >= la {
				return -1
			}
			if j >= lb {
				return 1
			}
			if a[i] != '^' {
				return 1
			}
			if b[j] != '^' {
				return -1
			}
			i++
			j++
			continue
		}
		if !(i < la && j < lb) {
			break
		}
		startI, startJ := i, j
		isNum := isDigit(a[i])
		if isNum {
			for i < la && isDigit(a[i]) {
				i++
			}
			for j < lb && isDigit(b[j]) {
				j++
			}
		} else {
			for i < la && isAlpha(a[i]) {
				i++
			}
			for j < lb && isAlpha(b[j]) {
				j++
			}
		}
		if startI == i {
			return -1
		}
		if startJ == j {
			if isNum {
				return 1
			}
			return -1
		}
		segA := a[startI:i]
		segB := b[startJ:j]
		if isNum {
			for len(segA) > 0 && segA[0] == '0' {
				segA = segA[1:]
			}
			for len(segB) > 0 && segB[0] == '0' {
				segB = segB[1:]
			}
			if len(segA) > len(segB) {
				return 1
			}
			if len(segB) > len(segA) {
				return -1
			}
		}
		cmp := strings.Compare(segA, segB)
		if cmp != 0 {
			if cmp < 0 {
				return -1
			}
			return 1
		}
	}
	if i >= la && j >= lb {
		return 0
	}
	if i >= la {
		return -1
	}
	return 1
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isAlnum(b byte) bool {
	return isDigit(b) || isAlpha(b)
}
