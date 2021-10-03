package main

import (
	"strings"
)

// Pos is a tuple representing cartesian coords
type Pos struct {
	x int
	y int
}

// Chessboard ...
type Chessboard struct {
	Board [][]bool
}

// NewEmptyChessboard returns a new empty chessboard of the given size
func NewEmptyChessboard(size int) *Chessboard {
	boardArray := make([][]bool, size)
	for i := 0; i < size; i++ {
		boardArray[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			boardArray[i][j] = false
		}
	}

	return &Chessboard{Board: boardArray}
}

func (c *Chessboard) String() string {
	var sb strings.Builder
	for i := 0; i < len(c.Board); i++ {
		for j := 0; j < len(c.Board); j++ {
			var rune rune
			if !c.Board[i][j] {
				rune = '.'
			} else {
				rune = 'Q'
			}
			sb.WriteRune(rune)
			if j != len(c.Board)-1 {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (c *Chessboard) CheckPosition(x, y int) bool {
	for i := 0; i < len(c.Board); i++ {
		if x == i {
			continue
		}
		if c.Board[x][i] {
			return false
		}
	}

	for i := 0; i < len(c.Board); i++ {
		if x == i {
			continue
		}
		if c.Board[i][y] {
			return false
		}
	}

	dirs := [4]Pos{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	i := 0

	for i < 4 {
		xx, yy := x, y
		xx += dirs[i].x
		yy += dirs[i].y
		for xx >= 0 && xx < len(c.Board) && yy >= 0 && yy < len(c.Board) {
			if c.Board[xx][yy] {
				return false
			}
			xx += dirs[i].x
			yy += dirs[i].y
		}
		i++
	}

	return true
}

func (c *Chessboard) Solve(depth int) bool {
	if depth == len(c.Board) {
		return true
	}
	for i := 0; i < len(c.Board); i++ {
		if c.CheckPosition(depth, i) {
			c.Board[depth][i] = true
			if c.Solve(depth + 1) {
				return true
			}
			c.Board[depth][i] = false
		}
	}
	return false
}
