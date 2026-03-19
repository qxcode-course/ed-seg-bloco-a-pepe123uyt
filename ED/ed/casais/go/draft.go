package main
import "fmt"
func main() {

    var num int
    fmt.Scan(&num)
    
    desca := []int{}
    pares := 0

    for i := 0; i < num; i++ {
		var ani int
		fmt.Scan(&ani)

		encontrou := false

		for j := 0; j < len(desca); j++ {
			if desca[j] == ani*(-1) {
                desca[j] = 0
				pares++
				encontrou = true
				break
			}
		}

		if !encontrou {
			desca = append(desca, ani)
		}
	}
    fmt.Println(pares)
}
