package main
import "fmt"

func matchingStrings(consultas, buscas []string) []int {
	cont := make(map[string]int)

	for _, v := range consultas {
		cont[v]++
	}

	var resu []int

	for _, q := range buscas {
		resu = append(resu, cont[q])
	}

	return resu
}

func main() {
	var n, m int
	var consultas, buscas []string

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		consultas = append(consultas, s)
	}

	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var q string
		fmt.Scan(&q)
		buscas = append(buscas, q)
	}

	resu := matchingStrings(consultas, buscas)

	for i := 0; i < len(resu); i++ {
        if i == 0 {
            fmt.Printf("%d", resu[0])
        } else {
            fmt.Printf(" %d", resu[i])
        }
	}

    fmt.Println()
}
