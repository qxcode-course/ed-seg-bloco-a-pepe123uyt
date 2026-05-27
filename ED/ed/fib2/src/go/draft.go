package main
import "fmt"

func fibo(n int) int {
    if n == 1 || n == 2 {
		return 1
	}
    if n == 3 {
		return 2
	}
    return fibo(n - 2) + fibo(n - 3)
}
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(fibo(n))

}
