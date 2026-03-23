package main
import "fmt"

func toString(v []int) {
	for i := 0; i < len(v); i++ {
		fmt.Printf("%d ", v[i])
	}
}
func main() {
	var n, m int
	fmt.Scan(&n)

	fila := []int{}

	for v := 0; v < n; v++ {
		var num int
		fmt.Scan(&num)
		fila = append(fila, num)
	}

	// fmt.Println(fila)

	fmt.Scan(&m)

	tirar := make(map[int]bool)

	for v := 0; v < m; v++ {
		var num int
		fmt.Scan(&num)
		tirar[num] = true
	} 

	// fmt.Println(tirar)

	novaFila := []int{}

	for _, v := range fila {
		if !tirar[v] {
			novaFila = append(novaFila, v)
		}
	}

	toString(novaFila)
	fmt.Println()
}
