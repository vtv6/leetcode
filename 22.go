package main

var cache = map[int]([]string){
	0: []string{""},
}

func generateParenthesis(n int) []string {

	if _, ok := cache[n]; ok {
		return cache[n]
	}

	result := []string{}

	for i := 0; i < n; i++ {
		for _, s1 := range generateParenthesis(i) {
			for _, s2 := range generateParenthesis(n - i - 1) {
				result = append(result, "("+s1+")"+s2)
			}
		}
	}
	cache[n] = result
	return result
}
