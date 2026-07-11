package main
import "fmt"

func backtrack(vet []int, n int) bool {
    if len(vet) == 0 || len(vet) == 1 && vet[0] != n {
        return false
    }

    
    return true
}

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    nums := make([]int, n)

    for i := range n {
        fmt.Scan(nums[i])
    }

    fmt.Println(backtrack(nums, k))

}