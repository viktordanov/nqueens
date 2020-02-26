package main

import "fmt"

func main() {
	chessboard := NewEmptyChessboard(8)

	chessboard.Solve(0)
	fmt.Println(chessboard)
}
