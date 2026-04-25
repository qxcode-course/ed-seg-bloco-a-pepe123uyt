package main
import "fmt"
func resto(num int){
    if num <= 0 {
        return
    }
    
    aux := num
    num /= 2

    resto(num)
    fmt.Println(num, aux%2)
}
func main() {
    var num int

    fmt.Scan(&num)

    resto(num)
}
