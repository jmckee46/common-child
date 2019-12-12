package main

func commonChild(s1 string, s2 string) int32 {
	s1Length := int32(len(s1))
	s2Length := int32(len(s2))

	return longestCommonSubstring(s1, s2, s1Length, s2Length)
}

func longestCommonSubstring(s1 string, s2 string, s1Length int32, s2Length int32) int32 {

	lcsLength := make([][]int32, s1Length+1)
	for i := range lcsLength {
		lcsLength[i] = make([]int32, s2Length+1)
	}

	for i := int32(0); i <= s1Length; i++ {
		for j := int32(0); j <= s2Length; j++ {
			if i == 0 || j == 0 {
				lcsLength[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				lcsLength[i][j] = lcsLength[i-1][j-1] + 1
			} else {
				lcsLength[i][j] = max(lcsLength[i-1][j], lcsLength[i][j-1])
			}
		}
	}

	return lcsLength[s1Length][s2Length]
}

func max(x int32, y int32) int32 {
	if x > y {
		return x
	}

	return y
}

func main() {}
