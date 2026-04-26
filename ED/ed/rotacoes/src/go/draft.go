package main
import "fmt"
func main() {
    var t, r int
    fmt.Scan(&t, &r)

    vet := make([]int, t)

    for i := 0; i < len(vet); i++ {
        fmt.Scan(&vet[i])
    }
    
    vetRot := rotacao(vet, r)

    for i := 0; i < len(vetRot); i++{
        if i == 0 {
            fmt.Printf("[ %d",vetRot[i])
            continue
        }

        fmt.Printf(" %d", vetRot[i])

        if i == len(vetRot)-1 {
            fmt.Println(" ]")
        }
    }
}
func rotacao(vet []int, num int)[]int {
    if num == 0 {
        return vet
    }

    ultimo := vet[len(vet)-1]

    for i := len(vet)-1; i > 0; i-- {
		vet[i] = vet[i-1]
	}

    vet[0] = ultimo

    return rotacao(vet, num-1)

}
