package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	n_linhas := len(grid)
	n_colunas := len(grid[0])

	for l := range n_linhas {
		for c := range n_colunas {
			if backtrack(grid, word, n_linhas, n_colunas, l, c) {
				return true
			}
		}
	}
	return false
}

func backtrack(grid [][]byte, suffix string, n_linhas, n_colunas, linha, coluna int) bool {
	if len(suffix) == 0 {
		return true
	}

	if linha < 0 || linha == n_linhas || coluna < 0 || coluna == n_colunas {
		return false
	}

	if grid[linha][coluna] != suffix[0] {
		return false
	}

	grid[linha][coluna] = '#'

	directions := [][]int {
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, dir := range directions {
		linhaOffset := dir[0]
		colunaOffset := dir[1]

		res := backtrack(
			grid,
			suffix[1:],
			n_linhas,
			n_colunas,
			linha + linhaOffset,
			coluna + colunaOffset,
		)

		if res {
			return true
		}
	}

	grid[linha][coluna] = suffix[0]
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
