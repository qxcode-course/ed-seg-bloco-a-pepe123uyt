package main
import "fmt"

func main() {
    var num int
    fmt.Scan(&num)
    // primeira(num)
    // segunda(num, 2)
    quadrado(num)
}
func quadrado(num int) int {
    if num == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", num, num-1, num-1)

    prev := quadrado(num-1)

    atual := prev + 2*(num-1) + 1
    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", num, num-1, num-1, atual)

    return atual
}
// primeira versao ficou um for disfaçado >-<

// func primeira(num int) {
//     if num == 1 {
//         fmt.Println("1^2 = 1")
//         return
//     }
//     fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", num, num-1, num-1)
//     primeira(num-1)
// }
// func segunda(num, i int){
//     if i > num {
//         return
//     }
//     fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", i, i-1, i-1, pow(i, 2))
//     segunda(num, i+1)
// }
// func pow(base, exp int) int {
// 	if exp == 0 {
// 		return 1
// 	}
// 	return base * pow(base, exp-1)
// }



