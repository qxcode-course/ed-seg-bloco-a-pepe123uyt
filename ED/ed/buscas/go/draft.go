package main
import "fmt"
func matchingStrings(strings [][]string, consultas )
func main() {
    var n int
    fmt.Scan(&n)

    consultas := []string{} 
    for i := 0; i < n; i++{
        var consu string
        fmt.Scan(&consu)
        consultas = append(consultas, consu)
    }

    fmt.Println(consultas)
}
