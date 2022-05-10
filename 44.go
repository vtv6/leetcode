package main

func isMatch(s string, p string) bool {
	cache := make([][]*bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		cache[i] = make([]*bool, len(p)+1)
	}
	return dp(0, 0, s, p, cache)
}

func dp(i, j int, s, p string, cache [][]*bool) bool {
	if cache[i][j] != nil {
		return *cache[i][j]
	}

	ans := false

	if j == len(p) {
		ans = i == len(s)
	} else {
		a := i < len(s) && (s[i] == p[j] || p[j] == '?')
		if p[j] == '*' {
			ans = dp(i, j+1, s, p, cache) || (i < len(s) && dp(i+1, j, s, p, cache))
		} else {
			ans = a && dp(i+1, j+1, s, p, cache)
		}
	}

	cache[i][j] = &ans
	return ans
}
