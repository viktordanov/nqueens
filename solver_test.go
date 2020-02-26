package main

import (
	"testing"
)

const o = false
const W = true

var board = &Chessboard{[][]bool{
	{W, o, o, o, o, o, o, o},
	{o, o, o, o, o, o, o, o},
	{o, o, o, o, W, o, o, o},
	{o, o, o, o, o, o, o, o},
	{o, W, o, o, o, o, W, o},
	{o, o, o, o, o, o, o, o},
	{o, o, o, o, o, o, o, o},
	{o, o, o, W, o, o, o, o},
}}

var checkPositionTests = []struct {
	in  Pos
	out bool
}{
	{Pos{0, 1}, false},
	{Pos{1, 0}, false},
	{Pos{7, 2}, false},
	{Pos{7, 7}, false},
	{Pos{5, 5}, false},
	{Pos{4, 3}, false},
	{Pos{0, 7}, true},
}

func TestCheckPosition(t *testing.T) {
	for _, tt := range checkPositionTests {
		t.Run("CheckPosition", func(t *testing.T) {
			if board.CheckPosition(tt.in.x, tt.in.y) != tt.out {
				t.Errorf("Position %v -> %t expected %t", tt.in, !tt.out, tt.out)
			}
		})
	}
}
