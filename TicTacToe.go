package main

import (
	"errors"
	"fmt"
)

func main() {
	//Instancio un nuevo board de tipo Board
	player1 := NewPlayer("X")
	player2 := NewPlayer("O")

	var n, m int

	fmt.Println("Hola, ingresa el tamaño del tablero que deseas (N x N)")
	_, err := fmt.Scanln(&n, &m)

	if err != nil {
		fmt.Println("Error: recuerda que solo se permiten numeros!")
	}

	board := initBoard(n, m)
	judge := NewJudge(board, player1, player2)
	turnManager := NewTurnManager(player1, player2)
	board.renderBoard()

	for !judge.gameOver() {

		fmt.Println("\nJugador", turnManager.actualTurn.markType, "ingresa tu movimiento (fila columna) filas 1 -", n, "columnas 1 -", m, ":")

		var row, col int
		_, err := fmt.Scanln(&row, &col)

		if err != nil {
			fmt.Println("Error: recuerda que solo se permiten numeros!")
		}

		turnManager.actualTurn.setMove(row-1, col-1)

		er := judge.isLegalMove(turnManager.actualTurn)

		if er != nil {
			fmt.Println("Error:", er)
		} else {
			turnManager.NextTurn()
		}

	}

}

// Board
type Board struct {
	Grid [][]string
	N    int
	M    int
	Size int
}

func initBoard(n int, m int) *Board {
	size := (m + n) / 2
	baseGrid := make([][]string, size) // Inicializar correctamente la matriz de tamaño `size`

	for i := 0; i < size; i++ {
		baseGrid[i] = make([]string, size) // Inicializar cada fila antes de usarla
		for j := 0; j < size; j++ {
			baseGrid[i][j] = "*"
		}
	}

	return &Board{
		Grid: baseGrid,
		Size: size,
		N:    n,
		M:    m,
	}
}

func (b *Board) renderBoard() {
	fmt.Println()

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Print(b.Grid[i][j] + "  ")
		}
		fmt.Println()
	}

	fmt.Println()
}

// Player
type Player struct {
	markType string
	move     [2]int
}

func NewPlayer(markType string) *Player {
	return &Player{markType: markType}
}

func (p *Player) setMove(row int, col int) {
	p.move[0] = row
	p.move[1] = col
}

//TurnManager

type TurnManager struct {
	actualTurn *Player
	nextTurn   *Player
}

func NewTurnManager(p1 *Player, p2 *Player) *TurnManager {
	return &TurnManager{
		actualTurn: p1,
		nextTurn:   p2,
	}
}

func (tm *TurnManager) NextTurn() {
	var tempPlayerForChange *Player = tm.actualTurn

	tm.actualTurn = tm.nextTurn
	tm.nextTurn = tempPlayerForChange
}

// Judge
type Judge struct {
	board   *Board
	player1 *Player
	player2 *Player
}

func NewJudge(b *Board, p1 *Player, p2 *Player) *Judge {
	return &Judge{board: b, player1: p1, player2: p2}
}

func (j *Judge) isLegalMove(player *Player) error {
	row := player.move[0]
	col := player.move[1]

	if row >= j.board.Size || col >= j.board.Size {
		return errors.New("recuerda que la matriz es 3x3")
	} else {
		if j.board.Grid[row][col] != "*" {
			return errors.New("casilla ocupada")
		} else {
			j.board.Grid[row][col] = player.markType
			j.board.renderBoard()
		}
	}

	return nil
}

func (j *Judge) gameOver() bool {
	//grid := j.board.Grid
	//size := j.board.Size
	//markTypePlayer1 := j.player1.markType
	//markTypePlayer2 := j.player2.markType

	/*
		TODO: Implementar funcion que me permita valdiar si gana player1 o player2 tener en cuenta que ahora la matriz es NxN y se puede ganar de manera HORIZONTAL, VERTICAL O DIAGONAL
	*/

	return false
}
