package main
import "fmt"

func matar(v []int, pos int) []int {
    novo := []int{}

    for i := 0; i < len(v); i++ {
        if i != pos {
            novo = append(novo, v[i])
        }
    }

    return novo
}

func main() {
    var num, espada int
    fmt.Scan(&num, &espada)

    fila := []int{}

    for i := 1; i <= num; i++ {
        fila = append(fila, i)
    }

    pos := espada + 1

    for len(fila) > 0 { 
        fmt.Print("[ ")
        for _, v:= range fila {
            fmt.Print(v, " ")
        }
        fmt.Println("]")

        fila = matar(fila, pos)

        if len(fila) > 0 {
            pos = pos % len(fila)
        }
    }

    
}
