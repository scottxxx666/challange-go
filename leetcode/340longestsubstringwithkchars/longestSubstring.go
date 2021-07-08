package _40longestsubstringwithkchars

func longestSubstring(str string, target int) int {
	m := make(map[byte]int)
	res := 0
	for i, j := 0, 0; j < len(str); j++ {
		m[str[j]]++

		for len(m) > target {
			if j-i > res {
				res = j - i
			}

			m[str[i]]--
			if m[str[i]] == 0 {
				delete(m, str[i])
			}
			i++
		}

		if j == len(str)-1 {
			if j-i+1 > res {
				res = j - i + 1
			}
		}
	}
	return res
}
