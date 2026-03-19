package main
import "fmt"
func main() {
    var comp int
    var ganhador int
    fmt.Scan(&comp)
    

    for i := 0; i < comp; i++{
        var a, b int
        fmt.Scan(&a, &b)

        if (a < 10 || b < 10) {
            continue
        }

        distA := a - b
        distB := distA
        

        if (distA >= distB){
            ganhador = i
        }
        

    }

    if (ganhador == 0){
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(ganhador)
    }
}
