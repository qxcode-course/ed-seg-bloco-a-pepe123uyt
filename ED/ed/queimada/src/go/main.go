package main

import (
	"bufio"
	"fmt"
	"os"
)

func burnTrees(grid [][]rune, l, c int) {
	// se estiver fora da matriz, retorne
	// se o elemento atual não for uma arvore, retorne
	// queime a arvore colocando o caractere 'o' na posição atual
	// chame a recursão para todos os 4 vizinhos
	nl := len(grid)
	nc := len(grid[0])

	if l < 0 || c < 0 || l >= nl || c >= nc {
		return
	}
	if grid[l][c] != '#' {
		return
	}
	
	grid[l][c] = 'o'

	burnTrees(grid, l+1, c) 
	burnTrees(grid, l-1, c) 
	burnTrees(grid, l, c+1) 
	burnTrees(grid, l, c-1) 

	// if grid[l+1][c] == '#' {
		// 	grid[l+1][c] = 'o'
		// 	burnTrees(grid, l+1, c)
	// }
	// if grid[l-1][c] == '#' {
	// 	grid[l-1][c] = 'o'
	// 	burnTrees(grid, l-1, c)
	// } 
	// if grid[l][c+1] == '#' {
	// 	grid[l][c+1] = 'o'
	// 	burnTrees(grid, l, c+1)
	// } 
	// if grid[l][c-1] == '#' {
	// 	grid[l][c-1] = 'o'
	// 	burnTrees(grid, l, c-1)
	// }

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
