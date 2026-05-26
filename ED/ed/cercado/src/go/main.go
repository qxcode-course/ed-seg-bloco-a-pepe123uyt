package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	r, c int
}

func dps(board [][]byte, visi map[Pos]bool, r, c int) {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] == 'X' {
		return
	}

	pos := Pos{r, c}
	if visi[pos] {
		return
	}
	visi[pos] = true
	
	dps(board, visi, r+1, c)
	dps(board, visi, r-1, c)
	dps(board, visi, r, c+1)
	dps(board, visi, r, c-1)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	visi := make(map[Pos]bool)

	for c := range board[0] {
		if board[0][c] == 'O' {
			dps(board, visi, 0, c)
		}
		if board[len(board)-1][c] == 'O' {
			dps(board, visi, len(board)-1, c)
		}
	}

	for r := range board {
		if board[r][0] == 'O' {
			dps(board, visi, 0, r)
		}
		if board[r][len(board[0])-1] == 'O' {
			dps(board,visi, r, len(board[0])-1)
		}
	}

	for r := range board {
		for c := range board[0]{
			pos := Pos{r,c}

			if board[r][c] == 'O' && !visi[pos] {
				board[r][c] = 'X'
			}
		} 
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
