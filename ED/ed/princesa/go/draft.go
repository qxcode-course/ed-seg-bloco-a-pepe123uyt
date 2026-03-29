package main

import "fmt"

func imprimir(fila []int, pos int) {
	fmt.Print("[ ")
	for i := 0; i < len(fila); i++ {
		if i == pos {
			fmt.Printf("%d> ", fila[i])
		} else {
			fmt.Printf("%d ", fila[i])
		}
	}
	fmt.Println("]")
}

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fila[i] = i + 1
	}

	pos := e - 1

	for len(fila) > 0 {
		imprimir(fila, pos)

		if len(fila) == 1 {
			break
		}

		morto := (pos + 1) % len(fila)
		fila = append(fila[:morto], fila[morto+1:]...)
		pos = morto % len(fila)
	}
}