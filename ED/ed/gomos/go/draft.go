package main
import "fmt"

type Pontos struct{
    x, y int
}

func main() {
    var q int
    var d string

    fmt.Scan(&q, &d)

    cobra := make([]Pontos, q)

    for i := 0; i < q; i++{
        fmt.Scan(&cobra[i].x, &cobra[i].y)
    }

    aux := make([]Pontos, q)
    copy (aux, cobra)

    switch d{
    case "L" :
        cobra[0].x--
    case "R" :
        cobra[0].x++
    case "D" :
        cobra[0].y++
    case "U" :
        cobra[0].y--
    }

    for i := 1; i < q; i++ {
		cobra[i] = aux[i-1]
	}

	for i := 0; i < q; i++ {
		fmt.Println(cobra[i].x, cobra[i].y)
	}

}
