package game

import (
	"math/rand"
	"time"
)

var moves = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

//Game struct type
type Game struct {
	board     [][]bool
	boardSize int
}

//NewGame generates new game instance
func NewGame(sz int) *Game {

	matrix := [][]bool{}
	for i := 0; i < sz; i++ {
		matrix = append(matrix, []bool{})
		for j := 0; j < sz; j++ {
			matrix[i] = append(matrix[i], false)
		}
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	matrix[r1.Intn(5)][r1.Intn(5)] = true
	matrix[r1.Intn(5)][r1.Intn(5)] = true
	matrix[r1.Intn(5)][r1.Intn(5)] = true
	matrix[r1.Intn(5)][r1.Intn(5)] = true

	return &Game{
		board:     matrix,
		boardSize: sz,
	}
}

func (g *Game) isCellExists(i, j int) bool {
	return i >= 0 && i < g.boardSize && j >= 0 && j < g.boardSize
}

func (g *Game) toggle(i, j int) {
	g.board[i][j] = !g.board[i][j]
}

//Move method
func (g *Game) Move(i, j int) {
	if !g.isCellExists(i, j) {
		return
	}

	for x := 0; x < len(moves); x++ {
		if g.isCellExists(moves[x][0], moves[x][1]) {
			g.toggle(moves[x][0], moves[x][1])
		}
	}
}

func (g *Game) randomToggle() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	g.toggle(r1.Intn(g.boardSize), r1.Intn(g.boardSize))
}

func (g *Game) isWin() bool {
	win := true
	for i := 0; i < g.boardSize; i++ {
		for j := 0; j < g.boardSize; j++ {
			win = win && g.board[i][j]
		}
	}
	return win
}