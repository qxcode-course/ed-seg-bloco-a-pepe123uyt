package main

import (
	"bufio"
	"fmt"
	"os"
)

func dps(board[][]byte, r, c int) {
	if r < 0 || c < 0 || r > len(board) - 1 || c > len(board[0]) - 1{
		return 
	}

	if board[r][c] != 'X' {
		return
	}

	board[r][c] = '.'

	dps(board, r+1, c)
	dps(board, r-1, c)
	dps(board, r, c+1)
	dps(board, r, c-1)

}
// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	count := 0

	for l := range board{
		for c := 0; c < len(board[0]); c++{
			if board[l][c] == 'X' {
				dps(board, l, c)
				count++
			}
		}
	}
	
	return count

}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
