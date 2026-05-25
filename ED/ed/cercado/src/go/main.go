package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	r, c int
}

func siege(board[][]byte, visitacao map[Pos]bool, r, c int) {
	if r == 0 && c == 0 {
		for 
	}

	if board[r][c] == 'O' {
		if board[r+1][c] == 'X' && board[r-1][c] == 'X' && board[r][c+1] == 'X' && board[r][c-1] == 'X' {
			board[r][c] = 'X'
		}
	}
	siege(board, r+1, c)
	
}
// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	visitacao := make(map[Pos]bool)
	siege(board, visitacao, 0, 0)
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
