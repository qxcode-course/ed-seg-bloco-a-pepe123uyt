package main
import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    fmt.Println(fibo(a, b))
}

func fibo(n, k int) int {
    if n == 1 || n == 2{
		return 1
	}
	return fibo(n-1, k) + k * fibo(n-2, k)
}
