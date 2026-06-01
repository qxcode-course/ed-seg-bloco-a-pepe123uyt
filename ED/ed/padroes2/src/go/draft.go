package main
import "fmt"

func req(n int) int {
    // if n == 1 {
    //     return n + 2
    // }

    soma := (n*n) + (n*2)
    return soma
}

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(req(n))
}