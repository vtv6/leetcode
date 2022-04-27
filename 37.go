package main

func solveSudoku(board [][]byte) {
	backTracking(board, 0, 0)
}

func backTracking(board [][]byte, i, j int) bool {
	if i > 8 {
		return true
	}
	if board[i][j] != '.' {
		if j == 8 {
			return backTracking(board, i+1, 0)
		} else {
			return backTracking(board, i, j+1)
		}
	}
	for x := 1; x <= 9; x++ {
		if checkValid(board, byte('0'+x), i, j) {
			board[i][j] = byte('0' + x)
			if j == 8 {
				if backTracking(board, i+1, 0) {
					return true
				} else {
					board[i][j] = '.'
					continue
				}
			} else {
				if backTracking(board, i, j+1) {
					return true
				} else {
					board[i][j] = '.'
					continue
				}
			}
		} else {
			continue
		}
	}
	return false
}

func checkValid(board [][]byte, num byte, i, j int) bool {
	//     check Row
	hashMap := make(map[byte]bool)
	for x := 0; x < 9; x++ {
		hashMap[board[i][x]] = true
	}
	if _, ok := hashMap[num]; ok {
		return false
	}

	//     check Col
	hashMap = make(map[byte]bool)
	for x := 0; x < 9; x++ {
		hashMap[board[x][j]] = true
	}
	if _, ok := hashMap[num]; ok {
		return false
	}

	//     check box
	h := i / 3
	k := j / 3
	hashMap = make(map[byte]bool)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			hashMap[board[h*3+x][k*3+y]] = true
		}
	}
	if _, ok := hashMap[num]; ok {
		return false
	}
	return true
}
