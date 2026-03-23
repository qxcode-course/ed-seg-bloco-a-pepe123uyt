package main
import "fmt"

func procurar_vivo(elementos []int) int {
    for i := 0; i < len(elementos); i++ {
        if elementos[i] != 0 {
            return elementos[i]
        }
    }
    return 0
}

func main(){
    var n, e int
    fmt.Scan(&n, &e)
    
    fila := []int{}
    
    for i := 1; i <= n; i++{
        fila = append(fila, i)
    }
    
    //abordagem 1
    
    for i := 0;; i++{
        fmt.Print("[ ")
        if i == e {
            
        }
        for j := 0; j < len(fila); j++{
            fmt.Printf("%d ", procurar_vivo(fila))
            if j == len(fila)-1 {
                fmt.Println("]")
            }
            
        }
        e++
    } 
}