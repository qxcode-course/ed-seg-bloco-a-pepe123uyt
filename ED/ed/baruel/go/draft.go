package main

import (
	"fmt"
	"sort"
)
func main() {

    var total, tem int
    fmt.Scan(&total, &tem)

    figus := [] int{}
    
    for i := 0; i < tem; i++{
        var figu int
        fmt.Scan(&figu)

        figus = append(figus, figu)
    }

    sort.Ints(figus)

    repe := [] int{}
    
    for i := 1; i < len(figus); i++ {
        if figus[i] == figus[i-1] {
            repe = append(repe, figus[i])
		}
	}
    
    faltam := [] int{}

	for i := 1; i <= total; i++ {
        achei := false

        for j := 0; j < len(figus); j++ {
            if figus[j] == i {
                achei = true
                break
            }
        }

        if !achei {
            faltam = append(faltam, i)
        }
    }

    if (len(repe) == 0){
        fmt.Println("N")
    } else {
        for i := 0; i < len(repe); i++{
            if (i == 0){
                fmt.Print(repe[0])
            } else {
                fmt.Printf(" %d", repe[i])
            }
        } 
        fmt.Println()
    }
    
    if (len(faltam) == 0){
        fmt.Println("N")
    } else {
        for i := 0; i < len(faltam); i++{
            if (i == 0){
                fmt.Print(faltam[0])
            } else {
                fmt.Printf(" %d", faltam[i])
            }
        } 
        fmt.Println()
    }

}
