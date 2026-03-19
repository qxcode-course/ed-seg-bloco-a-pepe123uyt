package main
import "fmt"
func main() {
    var heli, poli, fugi, dire int
    fmt.Scan(&heli, &poli, &fugi, &dire)

    if (dire == -1){
        for i := fugi;; i--{
            if (i == 0 && i != heli) {
                i = 15
            }
            if (i == poli){
                fmt.Println("N")
                break
            }
            if (i == heli){
                fmt.Println("S")
                break
            }
        }
    }

    if (dire == 1){
        for i := fugi;; i++{
            if (i == 15 && i != heli) {
                i = 0
            }
            if (i == poli){
                fmt.Println("N")
                break
            }
            if (i == heli){
                fmt.Println("S")
                break
            }
        }
    }
}
