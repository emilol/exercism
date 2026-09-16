package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	boardFile := cb[file]
	occupiedCount := 0

	for _, square := range boardFile {
		if square {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	occupiedCount := 0
	for _, file := range cb {
		square := file[rank-1]
		if square {
			occupiedCount++
		}
	}
	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedCount := 0
	for file := range cb {
		occupiedCount += CountInFile(cb, file)
	}
	return occupiedCount
}
