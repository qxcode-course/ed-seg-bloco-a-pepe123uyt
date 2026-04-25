package main
import "fmt"

func main() {
    var meses, pares int
    fmt.Scan(&meses, &pares)

    fmt.Println(fibo(meses, pares))
}

func fibo(n, k int) int {
    if n == 1 || n == 2{
		return 1
	}
	return fibo(n-1, k) + k * fibo(n-2, k)
}
