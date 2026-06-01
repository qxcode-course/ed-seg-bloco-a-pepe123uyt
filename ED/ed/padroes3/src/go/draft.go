package main
import "fmt"
func main() {
    var n, m int
    fmt.Scan(&n, &m)
    fmt.Println(points(n,m))
}

func points(n, m int) int {
    if m == 1 {
        return 1
    }

    return points(n , m - 1) + (n-2)*(m-1) + 1
}