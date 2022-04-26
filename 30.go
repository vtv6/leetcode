package main

func findSubstring(s string, words []string) []int {
	result := []int{}
	hashMap := make(map[string]int)

	//     hash map
	for _, word := range words {
		hashMap[word] += 1
	}

	wordLength := len(words[0])
	for i := 0; i < wordLength; i++ {
		result = append(result, slidingWindow(i, s, words, hashMap)...)
	}

	return result
}

func slidingWindow(left int, s string, words []string, hashMap map[string]int) []int {
	result := []int{}

	// wordsCount := len(words)
	wordLength := len(words[0])
	currentHashMap := make(map[string]int)

	right := left + wordLength
	for right <= len(s) {
		current := string(s[right-wordLength : right])
		if count, ok := hashMap[current]; ok {

			if count > currentHashMap[current] {
				currentHashMap[current] += 1

				valid := true
				for x, y := range hashMap {
					if currentHashMap[x] != y {
						valid = false
					}
				}

				if valid {
					result = append(result, left)
					currentHashMap[string(s[left:left+wordLength])] = currentHashMap[string(s[left:left+wordLength])] - 1
					left += wordLength
				}

			} else {
				checkIndex := left
				for ; string(s[checkIndex:checkIndex+wordLength]) != current && checkIndex < right-wordLength; checkIndex += wordLength {
					currentCheck := string(s[checkIndex : checkIndex+wordLength])
					currentHashMap[currentCheck] -= 1
				}
				left = checkIndex + wordLength
			}
		} else {
			left = right
			currentHashMap = make(map[string]int)
		}
		right += wordLength
	}

	return result
}
