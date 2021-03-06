package _longestsubstringwithoutrepeatchar

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	max := 0
	for i, j := 0, 0; j < len(s); j++ {
		if m[s[j]] {
			for s[i] != s[j] {
				m[s[i]] = false
				i++
			}
			i++
		} else {
			m[s[j]] = true
			if j-i+1 > max {
				max = j - i + 1
			}
		}
	}
	return max
}
