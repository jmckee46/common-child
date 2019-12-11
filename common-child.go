package main

func commonChild(s1 string, s2 string) int32 {
	s1Length := int32(len(s1))
	s2Length := int32(len(s2))

	return longestCommonSubstring(&s1, &s2, s1Length, s2Length)
}

func longestCommonSubstring(s1 *string, s2 *string, s1Length int32, s2Length int32) int32 {
	if s1Length == 0 || s2Length == 0 {
		return 0
	}
	if (*s1)[s1Length-1] == (*s2)[s2Length-1] {
		return 1 + longestCommonSubstring(s1, s2, s1Length-1, s2Length-1)
	}

	return max(longestCommonSubstring(s1, s2, s1Length, s2Length-1), longestCommonSubstring(s1, s2, s1Length-1, s2Length))
}

func max(x int32, y int32) int32 {
	if x > y {
		return x
	}

	return y
}

func main() {}
