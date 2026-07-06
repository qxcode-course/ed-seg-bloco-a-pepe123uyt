package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	positions := []Pos {
		{p.l + 1, p.c},
		{p.l, p.c + 1},
		{p.l - 1, p.c},
		{p.l, p.c - 1},
	}
	return positions
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	queue := NewQueue[Pos]()

	queue.Enqueue(startPos)

	visitados := make(map[Pos]bool)

	pos := Pos{startPos.l, startPos.c}

	visitados[pos] = true

	caminho := make(map[Pos]Pos)

	for !queue.IsEmpty() {
		n, _ := queue.Dequeue()

		if n == endPos {
			voltar(grid, caminho, startPos, endPos)
			return
		}

		for _, viz := range n.getNeig(){
			if !match(grid, viz, '#') && !visitados[viz]{
				visitados[viz] = true
				queue.Enqueue(viz)
				caminho[viz] = n
			}
			
		}
	}


}

func voltar(grid [][]rune, caminho map[Pos]Pos, startPos Pos, endPos Pos) {
	atual := endPos

	for atual != startPos {
		grid[atual.l][atual.c] = '.'
		atual = caminho[atual]
	}

	grid[startPos.l][startPos.c] = '.'

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
