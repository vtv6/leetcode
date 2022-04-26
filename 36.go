package main

func isValidSudoku(board [][]byte) bool {
	//     row,col
	hashMapRow := make(map[byte]bool)
	hashMapCol := make(map[byte]bool)
	for i := range board {
		hashMapRow = make(map[byte]bool)
		hashMapCol = make(map[byte]bool)
		for j := range board[0] {
			if board[i][j] != '.' && board[i][j] >= '0' && board[i][j] <= '9' {
				if _, ok := hashMapRow[board[i][j]]; ok {
					return false
				} else {
					hashMapRow[board[i][j]] = true
				}
			}

			if board[j][i] != '.' && board[j][i] >= '0' && board[j][i] <= '9' {
				if _, ok := hashMapCol[board[j][i]]; ok {
					return false
				} else {
					hashMapCol[board[j][i]] = true
				}
			}

		}
	}

	//    sub-boxes
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			hashMapRow = make(map[byte]bool)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[3*i+x][3*j+y] != '.' && board[3*i+x][3*j+y] >= '0' && board[3*i+x][3*j+y] <= '9' {
						if _, ok := hashMapRow[board[3*i+x][3*j+y]]; ok {
							return false
						} else {
							hashMapRow[board[3*i+x][3*j+y]] = true
						}
					}
				}
			}
		}
	}
	return true
}
