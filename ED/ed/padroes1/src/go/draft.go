package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(bloquinhos(n))

}
func bloquinhos(num int) int {
    if num == 1 {
        return 20
    } 
    return 8 + bloquinhos(num-1)
}