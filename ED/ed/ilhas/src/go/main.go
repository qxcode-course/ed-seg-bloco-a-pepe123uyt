package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func dfs(grid [][]byte, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] == '0' {
		return
	}

	grid[l][c] = '0'

	dfs(grid, l + 1, c)
	dfs(grid, l - 1, c)
	dfs(grid, l, c + 1)
	dfs(grid, l, c - 1)

}
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	islands := 0

	for l := range grid{
		for c := 0; c < len(grid[0]); c++{
			if grid[l][c] == '1' {
				dfs(grid, l, c)
				islands++
			}
		}
	}
	
	return islands
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
