package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, stack *Stack[Pos], l, c int) {
	if grid[l][c] == '.'{
		return
	}
	
	Posi := Pos{l,c}
	stack.Push(Posi)

	grid[l][c] = 'o'


	//fmt.Println(stack)

	burnTrees(grid, stack, l+1, c)
	burnTrees(grid, stack, l -1, c)
	burnTrees(grid, stack, l, c+1)
	burnTrees(grid, stack, l, c-1)
	
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha
	
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
	stack := NewStack[Pos]()
	burnTrees(grid, stack, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
